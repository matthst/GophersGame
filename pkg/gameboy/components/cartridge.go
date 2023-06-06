package components

type Cart interface {
	Write(val uint8, adr uint16)
	Read(adr uint16) uint8
}

type RomOnly struct {
	Rom []uint8
}

func (r RomOnly) Write(val uint8, adr uint16) {}

func (r RomOnly) Read(adr uint16) uint8 {
	return r.Rom[adr]
}

// getROMSize returns the number of ROM banks based on the header byte 0x0148
func getROMSize(val uint8) int {
	return 32 * (1 << val)
}

// getRAMSize returns the number of RAM bank based on the header byte 0x0147
func getRAMSize(val uint8) int {
	switch val {
	case 0x02:
		return 8
	case 0x03:
		return 32
	case 0x04:
		return 128
	case 0x05:
		return 64
	default:
		return 0
	}
}
