package video

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Video struct {

	// data section

	// render section
	tileImages [384]ebiten.Image
}

func (v *Video) Load(adr uint16) uint8 {
	// TODO
	return 1
}

func (v *Video) Write(val uint8, adr uint16) {

}
