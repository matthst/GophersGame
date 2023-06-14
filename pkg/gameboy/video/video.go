package video

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"sort"
)

type Video struct {

	// render section

	RenderImage *ebiten.Image

	// TODO [GBC] Hi-Color Mode
	// TODO OAM DMA TRANSFER

	// data section
	vram [0x2000]uint8
	oam  [0x1000]uint8

	bgPalette, obPalette0, obPalette1 [4]color.Color
	bgp, obp0, obp1                   uint8
}

var (
	monochromePalette = [4]color.Color{color.Gray{0xFF}, color.Gray{0xF0}, color.Gray{0x0F}, color.Gray{0x00}}

	scanLineCycleCount int

	scx, scy, wx, wy, ly, lyc, lcdc, stat, statMode uint8

	statBlock, lycEqualsLY bool

	// LCDC bool stuff
	ppuEnable, winTileMapArea, winEnable, bgWinAddressingMode, bgTileMapArea, objSize, objEnable, bgWinEnable bool
)

func GetDmgVideo() Video {
	vid := Video{
		bgPalette:   monochromePalette,
		obPalette0:  monochromePalette,
		obPalette1:  monochromePalette,
		RenderImage: ebiten.NewImage(160, 144),
	}
	vid.WriteToIORegisters(0x91, 0xFF40)
	vid.WriteToIORegisters(0xFC, 0xFF47)
	return vid
}

func (v *Video) LoadFromVRAM(adr uint16) uint8 {
	if statMode == 3 {
		return 0xFF
	}
	return v.vram[adr-0x8000]
}

func (v *Video) LoadFromOAM(adr uint16) uint8 {
	if statMode > 1 {
		return 0xFF
	}
	return v.oam[adr-0xFE00]
}

func (v *Video) LoadFromIORegisters(adr uint16) uint8 {
	switch adr { //IO Registers
	case 0xFF40:
		return lcdc
	case 0xFF41:
		lycEqualsLYVal := uint8(0)
		if lycEqualsLY {
			lycEqualsLYVal = 4
		}
		return (stat & 0b0111_1100) + lycEqualsLYVal + statMode
	case 0xFF42:
		return scy
	case 0xFF43:
		return scx
	case 0xFF44:
		return ly
	case 0xFF45:
		return lyc
	case 0xFF46:
		// TODO OAM COPY handling
	case 0xFF47:
		return v.bgp //TODO [CGB] palette handling (don't forget no access during Mode 3
	case 0xFF48:
		return v.obp0
	case 0xFF49:
		return v.obp1
	case 0xFF4A:
		return wy
	case 0xFF4B:
		return wx
	case 0xFF4D:
		return 0x00 // TODO [CGB] speed switching
	case 0xFF4F:
		return 0x00 // TODO [CGB] vram bank switching
	case 0xFF55:
		return 0x00 // TODO [CGB] more stuff
	}
	return 0xFF
}

func (v *Video) WriteToVRAM(val uint8, adr uint16) {
	if statMode != 3 {
		v.vram[adr-0x8000] = val
	}
}

func (v *Video) WriteToOAM(val uint8, adr uint16) {
	if statMode < 2 {
		v.vram[adr-0xFE00] = val
	}
}

func (v *Video) WriteToIORegisters(val uint8, adr uint16) {
	switch adr {
	case 0xFF40:
		lcdc = val
		updateLCDCFlags()
	case 0xFF41:
		stat = (val & 0b0111_1000) | statMode
	case 0xFF42:
		scy = val
	case 0xFF43:
		scx = val
	case 0xFF45:
		lyc = val
	case 0xFF46:
		// TODO OAM COPY handling
		panic("OAM Copy not implemented")
	case 0xFF47:
		v.bgp = val //TODO [CGB] palette handling
		updatePalette(val, &v.bgPalette)
	case 0xFF48:
		v.obp0 = val
		updatePalette(val, &v.obPalette0)
	case 0xFF49:
		v.obp1 = val
		updatePalette(val, &v.obPalette1)
	case 0xFF4A:
		wy = val
	case 0xFF4B:
		wx = val
		// TODO [CGB] speed switching
		// TODO [CGB] vram bank switching
		// TODO [CGB]
	}
}

// MCycle executes an mCycle for the video controller and returns interrupts if any occur
func (v *Video) MCycle() uint8 {
	statResult := uint8(0)
	scanLineCycleCount++
	if scanLineCycleCount == 114 {
		scanLineCycleCount = 0 // MODE 2, Search
		statMode = 2
		ly++

		if ly == 144 { // MODE 1, VBLANK
			statMode = 1
			statResult = 1
		} else if ly == 154 {
			ly = 0
		}
	}

	if ly < 144 {
		switch {
		case scanLineCycleCount == 20: // MODE 3, transfer to LCD controller
			statMode = 3
			v.drawScanLine()
		case scanLineCycleCount == 63: // MODE 0, HBLANK
			statMode = 0
		}
	}
	// interrupt handling
	statInterruptSource := false
	lycEqualsLY = lyc == ly

	switch {
	case statMode == 0 && stat&0b1000 != 0:
		statInterruptSource = true
	case statMode == 1 && stat&0b0001_0000 != 0:
		statInterruptSource = true
	case statMode == 2 && stat&0b0010_0000 != 0:
		statInterruptSource = true
	case lycEqualsLY && stat&0b0100_0000 != 0:
		statInterruptSource = true
	}

	if !statInterruptSource {
		statBlock = false
	} else if statInterruptSource && !statBlock {
		statResult |= 0b10
		statBlock = true
	}

	return statResult
}

func updateLCDCFlags() {
	ppuEnable = lcdc&0b1000_0000 != 0
	winTileMapArea = lcdc&0b0100_0000 != 0
	winEnable = lcdc&0b0010_0000 != 0
	bgWinAddressingMode = lcdc&0b0001_0000 != 0
	bgTileMapArea = lcdc&0b0000_1000 != 0
	objSize = lcdc&0b0000_0100 != 0
	objEnable = lcdc&0b0000_0010 != 0
	bgWinEnable = lcdc&0b0000_0001 != 0
}

func updatePalette(val uint8, palette *[4]color.Color) { //TODO [CGB] Color palette handling
	for i := 0; i < 4; i++ {
		palette[i] = monochromePalette[((val >> (i * 2)) & 0b11)]
	}
}

func (v *Video) drawScanLine() {

	//setup for BG and Win
	bgMapBaseAddress := 0x1800 + (uint16(ly-scy)/8)*32 //background map base address with right row selected
	if bgTileMapArea {

		bgMapBaseAddress += 0x0400 //switch bg map to second map if needed
	}

	// Setup for objects
	spriteHeight := uint8(8)
	objIds := make([]int, 0)
	if false && objEnable { // TODO enable objects later
		if objSize {
			spriteHeight += 8
		}

		// grab objects for drawing, up to 10
		for oamIndex, selectIndex := 0, 0; oamIndex < 160 && selectIndex < 10; oamIndex += 4 {
			if v.oam[oamIndex]+16 <= ly && v.oam[oamIndex]+16+spriteHeight >= ly {
				objIds = append(objIds, oamIndex)
				selectIndex++
			}
		}

		//sort the chosen objects by their x coordinate to make iterating over them faster
		sort.SliceIsSorted(objIds, func(a, b int) bool {
			return v.oam[a+1] < v.oam[b+1]
		})
	}

DrawLoop:
	for x := uint8(0); x < 160; x++ {
		if false && objEnable { //OBJ // TODO enable objects later
			for _, objId := range objIds {
				if v.oam[objId+1]+16 < x { // obj is before x
					continue
				}
				if v.oam[objId+1]+8 > x { // obj is after x
					break
				}
				// TODO GET Pixel
				pixel := 0      //placeholder
				if pixel != 0 { // TODO add condition that checks if obj is behind background
					// DRAW
					continue DrawLoop // go to next pixel
				}
			}
		}

		// BG and Window enable TODO [CGB] bg and window priority
		if !bgWinEnable {
			v.RenderImage.Set(int(x), int(ly), color.White)
			continue
		}

		//Window
		if winEnable && wy <= ly && wx+7 <= x {
			// TODO get pixel from window
		}

		//BG
		// TODO get pixel from BG

		tileIndexAddress := bgMapBaseAddress + uint16(x-scx)/8
		baseTileIndex := uint16(v.vram[tileIndexAddress])
		tileAddress := baseTileIndex*16 + (uint16(ly)%8)*2
		if !bgWinAddressingMode && baseTileIndex < 128 {
			tileAddress += 0x1000
		}
		xOffset := 6 - x%8
		colorIndex := (v.vram[tileAddress]>>xOffset+1)&0b1 + (v.vram[tileAddress]>>xOffset)&0b10
		v.RenderImage.Set(int(x), int(ly), v.bgPalette[colorIndex])
	}

}
