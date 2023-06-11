package video

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Video struct {

	// render section

	renderImage *ebiten.Image

	// TODO [GBC] Hi-Color Mode
	// TODO OAM DMA TRANSFER

	// data section
	vram []uint8
	oam  []uint8

	bgPalette, obPalette0, obPalette1 [4]color.Color
	bgp, obp0, obp1                   uint8
}

var (
	monochromePalette = [4]color.Color{color.Gray{0x00}, color.Gray{0x0F}, color.Gray{0xF0}, color.Gray{0xFF}}

	scanLineCycleCount int

	scx, scy, wx, wy, ly, lyc, lcdc, stat, statMode uint8

	statBlock bool
)

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
		return (stat & 0b0111_1100) + statMode
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
			// TODO draw scanLine
		case scanLineCycleCount == 63: // MODE 0, HBLANK
			statMode = 0
		}
	}
	// interrupt handling
	statInterruptSource := false
	if lyc == ly {
		statMode |= 0b0100
	} else {
		statMode &= 0b1011
	}

	switch {
	case statMode == 0 && stat&0b1000 != 0:
		statInterruptSource = true
	case statMode == 1 && stat&0b0001_0000 != 0:
		statInterruptSource = true
	case statMode == 2 && stat&0b0010_0000 != 0:
		statInterruptSource = true
	case stat&0b_0100 != 0 && stat&0b0100_0000 != 0:
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

func updatePalette(val uint8, palette *[4]color.Color) { //TODO [CGB] Color palette handling
	for i := 0; i < 4; i++ {
		palette[i] = monochromePalette[((val >> (i * 2)) & 0b11)]
	}
}

func drawScanLine() {
	for x := 0; x < 160; x++ {
		//OBJ

		//Window

		//BG

	}
}
