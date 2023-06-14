package components

type WRAM struct {
	ram [0x2000]uint8
}

func (w *WRAM) Load(adr uint16) uint8 {
	return w.ram[adr - 0xC000]
}

func (w *WRAM) Write(val uint8, adr uint16) {
	w.ram[adr - 0xC000] = val
}
