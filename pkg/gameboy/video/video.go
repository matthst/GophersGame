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

	scanLineCycleCount int

	SCX, SCY, WX, WY, LY, LYC, LCDC, STAT uint8
}

func (v *Video) Load(adr uint16) uint8 {
	// TODO
	return 1
}

func (v *Video) Write(val uint8, adr uint16) {

	if v.STAT&0b0011 != 3 { // TODO access control because of vblank periods
		switch {
		}
	}
}

func (v *Video) Tick() {
	v.scanLineCycleCount++
	if v.scanLineCycleCount == 114 {
		v.scanLineCycleCount = 0 // MODE 2, search
		v.STAT = (v.STAT & 0b1111_1100) + 0b10
		v.LY++

		if v.LY == 144 { // MODE 1, VBLANK
			v.STAT = (v.STAT & 0b1111_1100) + 0b01
			// TODO Interrupt handling
		} else if v.LY == 154 {
			v.LY = 0
		}
	}

	if v.LY < 144 {
		switch {
		case v.scanLineCycleCount >= 20: // MODE 3, transfer to LCD controller
			v.STAT = (v.STAT & 0b1111_1100) + 0b11
			// TODO draw scanLine
		case v.scanLineCycleCount >= 63: // MODE 0, HBLANK
			v.STAT = (v.STAT & 0b1111_1100) + 0b00
			//TODO VBLANK STAT Interrupt handling
		}
	}
}
