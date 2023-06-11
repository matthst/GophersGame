package video

import "github.com/hajimehoshi/ebiten/v2"

type Video struct {

	// render section

	renderImage *ebiten.Image

	// TODO [GBC] Hi-Color Mode
	// TODO OAM DMA TRANSFER
	// TODO LCD STAT Interrupt handling
	// TODO LYC LY FLAG

	// data section
	tileData []uint8

	tileMaps []uint8

	scanLineCycleCount, StatMode int

	SCX, SCY, WX, WY, LY, LYC, LCDC uint8

	LYCLY, LYCLYStatI, Mode2OAMStatI, Mode1VBlankStatI, Mode0HBlankStatI, statBlock bool
}

func (v *Video) Load(adr uint16) uint8 {

	switch {
	case adr == 0xFF44:
		return v.LY
	}

	// TODO
	return 1
}

func (v *Video) Write(val uint8, adr uint16) {

	if v.StatMode != 3 { // TODO access control because of vblank periods
		switch {
		}
	}
}

// Tick executes an mCycle for the video controller and returns interrupts if any occur
func (v *Video) Tick() uint8 {
	statResult := uint8(0)
	v.scanLineCycleCount++
	if v.scanLineCycleCount == 114 {
		v.scanLineCycleCount = 0 // MODE 2, search
		v.StatMode = 2
		v.LY++

		if v.LY == 144 { // MODE 1, VBLANK
			v.StatMode = 1
			statResult = 1
		} else if v.LY == 154 {
			v.LY = 0
		}
	}

	if v.LY < 144 {
		switch {
		case v.scanLineCycleCount == 20: // MODE 3, transfer to LCD controller
			v.StatMode = 3
			// TODO draw scanLine
		case v.scanLineCycleCount == 63: // MODE 0, HBLANK
			v.StatMode = 0
		}
	}

	// interrupt handling
	statInterruptSource := false
	v.LYCLY = v.LYC == v.LY

	switch {
	case v.StatMode == 0 && v.Mode0HBlankStatI:
		statInterruptSource = true
	case v.StatMode == 1 && v.Mode1VBlankStatI:
		statInterruptSource = true
	case v.StatMode == 2 && v.Mode2OAMStatI:
		statInterruptSource = true
	case v.LYCLY && v.LYCLYStatI:
		statInterruptSource = true

	}

	if !statInterruptSource {
		v.statBlock = false
	} else if statInterruptSource && !v.statBlock {
		statResult |= 0b10
		v.statBlock = true
	}

	return statResult
}
