package components

type HRAM struct {
	ram [0x80]uint8
}

func (h *HRAM) Load(adr uint16) uint8 {
	return h.ram[adr-0xFF80]
}

func (h *HRAM) Write(val uint8, adr uint16) {
	h.ram[adr-0xFF80] = val
}
