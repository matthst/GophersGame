package Input

type Input struct {
	aBtn, bBtn, upBtn, downBtn, leftBtn, rightBtn, selectBtn, startBtn bool
}

func (i *Input) Load(adr uint16) uint8 {
	return 0x0F
}

func (i *Input) Write(val uint8, adr uint16) {

}
