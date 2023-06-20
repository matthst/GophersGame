package video

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"sort"
)

type Video struct {

	// render section

	// TODO [GBC] Hi-Color Mode

	// data section
	tileData [0x1800]uint8
	tileMaps [0x0800]uint8
	oam      [160]uint8

	bgPalette, obPalette0, obPalette1 [4]color.Color
	bgp, obp0, obp1                   uint8
}

var (
	monochromePalette = [4]color.Color{
		color.RGBA{R: 0xC7, G: 0xC7, B: 0xC7},
		color.RGBA{R: 0xA0, G: 0xA0, B: 0x88},
		color.RGBA{R: 0x66, G: 0x66, B: 0x4d},
		color.RGBA{R: 0x27, G: 0x27, B: 0x15}}

	CurrentFrame *ebiten.Image
	lastFrame    *ebiten.Image

	scanLineCycle int

	windowLineCounter                               uint16
	scx, scy, wx, wy, ly, lyc, lcdc, stat, statMode uint8

	statBlock, lycEqualsLY, startedWindowDrawing bool

	// LCDC bool stuff
	lcdEnable, winTileMapArea, winEnable, bgWinAddressingMode, bgTileMapArea, objSize, objEnable, bgWinEnable bool
)

func GetDmgVideo() Video {
	vid := Video{
		bgPalette:  monochromePalette,
		obPalette0: monochromePalette,
		obPalette1: monochromePalette,
	}
	ly = 91
	lastFrame = ebiten.NewImage(160, 144)
	CurrentFrame = ebiten.NewImage(160, 144)
	lastFrame.Fill(monochromePalette[0])
	CurrentFrame.Fill(monochromePalette[0])
	vid.WriteToIORegisters(0x91, 0xFF40)
	vid.WriteToIORegisters(0x81, 0xFF41)
	vid.WriteToIORegisters(0xFC, 0xFF47)
	return vid
}

func (v *Video) LoadFromVRAM(adr uint16) uint8 {
	if statMode == 3 {
		return 0xFF
	}
	if adr < 0x9800 {
		return v.tileData[adr-0x8000]
	} else {
		return v.tileMaps[adr-0x9800]
	}
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
	if statMode < 3 {
		if adr < 0x9800 {
			v.tileData[adr-0x8000] = val
		} else {
			v.tileMaps[adr-0x9800] = val
		}
	}
}

func (v *Video) WriteToOAM(val uint8, adr uint16) {
	if statMode < 2 {
		v.oam[adr-0xFE00] = val
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
	if !lcdEnable {
		return 0
	}

	statResult := uint8(0)
	scanLineCycle++
	if scanLineCycle == 114 {
		scanLineCycle = 0
		ly++

		if ly == 144 { // MODE 1, VBLANK
			swapFrames()
			statMode = 1
			statResult = 1
			windowLineCounter = 0
			startedWindowDrawing = false
		} else if ly == 154 {
			ly = 0
		}
	}

	if ly < 144 {
		switch {
		case scanLineCycle == 0: // MODE 2, Search
			statMode = 2
		case scanLineCycle == 20: // MODE 3, transfer to LCD controller
			statMode = 3
		case scanLineCycle == 40:
			v.drawScanLine()
		case scanLineCycle == 63: // MODE 0, HBLANK
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

func (v *Video) drawScanLine() {
	if !lcdEnable {
		return
	}

	//setup for BG and Win
	windowVisible := winEnable && wy <= ly && wx < 166
	if windowVisible {
		windowLineCounter++
		if !startedWindowDrawing {
			startedWindowDrawing = true
			windowLineCounter = uint16(ly - wy)
		}
	}

	ySCYOffset := uint16(ly + scy)
	bgMapBaseAddress := (ySCYOffset / 8) * 32         //background map base address with right row selected
	winMapBaseAddress := (windowLineCounter / 8) * 32 //window map base address with right row selected
	if bgTileMapArea {
		bgMapBaseAddress += 0x0400 //switch bg map to second map if needed
	}
	if winTileMapArea {
		winMapBaseAddress += 0x0400 //switch window map to second map if needed
	}

	// Setup for objects
	spriteHeight := uint8(8)
	objIds := make([]int, 0)
	if objEnable {
		if objSize {
			spriteHeight += 8
		}

		// grab objects for drawing, up to 10
		for oamIndex, selectIndex := 0, 0; oamIndex < 160 && selectIndex < 10; oamIndex += 4 {
			if v.oam[oamIndex] <= ly+16 && v.oam[oamIndex]+spriteHeight > ly+16 {
				objIds = append(objIds, oamIndex)
				selectIndex++
			}
		}

		//sort the chosen objects by their x coordinate to make iterating over them faster
		sort.SliceStable(objIds, func(a, b int) bool {
			return v.oam[objIds[a]+1] < v.oam[objIds[b]+1] || (v.oam[objIds[a]+1] == v.oam[objIds[b]+1] && objIds[a] < objIds[b])
		})
	}

DrawLoop:
	for x := uint8(0); x < 160; x++ {
		bgOverObj := false

		if objEnable { //OBJ
			for _, objId := range objIds {
				bgOverObj = false
				if x >= v.oam[objId+1] { // obj is before x
					continue
				} else if x+8 < v.oam[objId+1] { // obj is after x
					break
				}

				var yFlip, xFlip, palette bool
				bgOverObj, yFlip, xFlip, palette = getObjFlags(v.oam[objId+3])
				yOffset := uint16((ly + 16 - v.oam[objId]) % spriteHeight)
				xOffset := 7 - ((x + 8 - v.oam[objId+1]) % 8)
				if yFlip {
					yOffset = (uint16(spriteHeight) - 1) - yOffset
				}
				if xFlip {
					xOffset = 7 - xOffset
				}

				tileIndex := uint16(v.oam[objId+2])
				if objSize {
					tileIndex &= 0xFE
				}

				tileByteRowAddr := tileIndex*16 + yOffset*2
				colorIndex := ((v.tileData[tileByteRowAddr] >> xOffset) & 1) + (((v.tileData[tileByteRowAddr+1] >> xOffset) << 1) & 2)
				if colorIndex != 0 {
					var paletteRef *[4]color.Color
					if palette {
						paletteRef = &v.obPalette1
					} else {
						paletteRef = &v.obPalette0
					}
					lastFrame.Set(int(x), int(ly), paletteRef[colorIndex])
					if bgOverObj {
						break
					}
					continue DrawLoop // go to next pixel
				} else {
					bgOverObj = false
				}
			}
		}

		// BG and Window enable TODO [CGB] bg and window priority
		if !bgWinEnable {
			lastFrame.Set(int(x), int(ly), monochromePalette[0b00])
			continue
		}

		//Window
		if windowVisible && wx < x+8 {
			xWXOffset := uint16(x + 7 - wx)
			colorIndex := v.getPixelFromMap(winMapBaseAddress, xWXOffset, windowLineCounter)
			if bgOverObj && colorIndex == 0 {
				continue
			}
			lastFrame.Set(int(x), int(ly), v.bgPalette[colorIndex])
			continue
		}

		//BG
		xSCXOffset := uint16(x + scx)
		colorIndex := v.getPixelFromMap(bgMapBaseAddress, xSCXOffset, ySCYOffset)
		if bgOverObj && colorIndex == 0 {
			continue
		}
		lastFrame.Set(int(x), int(ly), v.bgPalette[colorIndex])
	}
}

func (v *Video) getPixelFromMap(bgMapBaseAddress uint16, xMapOffset uint16, yMapOffset uint16) uint8 {
	tileIndexAddress := bgMapBaseAddress + xMapOffset/8
	baseTileIndex := uint16(v.tileMaps[tileIndexAddress])
	tileByteRowAddr := baseTileIndex*16 + (yMapOffset%8)*2
	if !bgWinAddressingMode && baseTileIndex < 128 {
		tileByteRowAddr += 0x1000
	}
	xOffset := 7 - xMapOffset%8
	colorIndex := ((v.tileData[tileByteRowAddr] >> xOffset) & 1) + (((v.tileData[tileByteRowAddr+1] >> xOffset) << 1) & 2)
	return colorIndex
}

func getObjFlags(val uint8) (bool, bool, bool, bool) {
	bgOverObj := val&0b1000_0000 != 0
	yFlip := val&0b0100_0000 != 0
	xFlip := val&0b0010_0000 != 0
	palette := val&0b0001_0000 != 0
	return bgOverObj, yFlip, xFlip, palette
}

func updateLCDCFlags() {
	lcdEnable = lcdc&0b1000_0000 != 0
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

func swapFrames() {
	swap := CurrentFrame
	CurrentFrame = lastFrame
	lastFrame = swap
}
