package cartridge

type Cart interface {
	write(val uint8, adr uint16)
	read(adr uint16) uint8
}

type RomOnly struct {
	Rom []uint8
}

func (r RomOnly) write(val uint8, adr uint16) {}

func (r RomOnly) read(adr uint16) uint8 {
	return r.Rom[adr]
}
