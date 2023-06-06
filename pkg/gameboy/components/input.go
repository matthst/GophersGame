package components

type Input struct {
}

func (i *Input) Load(adr uint16) uint8 {
	return 0x0000
}

func (i *Input) Write(val uint8, adr uint16) {

}
