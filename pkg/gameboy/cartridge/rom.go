package Cartridge

type RomOnly struct {
	rom []uint8
}

func CreateRomOnly(rom []uint8) *RomOnly {
	return &RomOnly{rom: rom}
}

func (c *RomOnly) Write(val uint8, adr uint16) {
}

func (c *RomOnly) Load(adr uint16) uint8 {
	if adr < 0x8000 {
		return c.rom[adr]
	}
	return 0x00
}

func (c *RomOnly) GetCartType() string {
	return "Rom Only"
}
