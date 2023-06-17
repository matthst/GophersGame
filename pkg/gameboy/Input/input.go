package Input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	keyPresses []ebiten.Key

	FF00 uint8

	actionBtns, dirBtns, interrupt                                     bool
	aBtn, bBtn, upBtn, downBtn, leftBtn, rightBtn, selectBtn, startBtn bool
)

func Load() uint8 {
	result := uint8(0b1111)
	if dirBtns {
		if aBtn {
			result &= 0b1110
		}
		if bBtn {
			result &= 0b1101
		}
		if selectBtn {
			result &= 0b1011
		}
		if startBtn {
			result &= 0b0111
		}
	}

	if actionBtns {
		if rightBtn {
			result &= 0b1110
		}
		if leftBtn {
			result &= 0b1101
		}
		if upBtn {
			result &= 0b1011
		}
		if downBtn {
			result &= 0b0111
		}
	}

	return FF00 + result
}

func Write(val uint8) {
	dirBtns = val&0b1_0000 == 0
	actionBtns = val&0b10_0000 == 0
	FF00 = val & 0b0011_0000
	setInterruptValue()
}

func Cycle() uint8 {
	if interrupt {
		interrupt = false
		return 0b1_0000
	}
	return 0
}

// SetInputState is called once during Tick update
func RunTick() {
	keyPresses = inpututil.AppendPressedKeys(keyPresses[:0])
	aBtn = containsKey(ebiten.KeyA)
	bBtn = containsKey(ebiten.KeyB)
	upBtn = containsKey(ebiten.KeyUp)
	downBtn = containsKey(ebiten.KeyArrowDown)
	leftBtn = containsKey(ebiten.KeyArrowLeft)
	rightBtn = containsKey(ebiten.KeyArrowRight)
	startBtn = containsKey(ebiten.KeyEnter)
	selectBtn = containsKey(ebiten.KeyBackspace)
	setInterruptValue()
}

func setInterruptValue() {
	if (actionBtns && (aBtn || bBtn || startBtn || selectBtn)) || (dirBtns && (upBtn || downBtn || leftBtn || rightBtn)) {
		interrupt = true
	}
}

func containsKey(key ebiten.Key) bool {
	for _, keyPress := range keyPresses {
		if keyPress == key {
			return true
		}
	}
	return false
}
