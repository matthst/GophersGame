package Cartridge

type RomOnly struct {
	rom []uint8
}

func CreateRomOnly(rom []uint8) RomOnly {
	return RomOnly{rom: rom}
}

func (r RomOnly) Write(val uint8, adr uint16) {
}

func (r RomOnly) Load(adr uint16) uint8 {
	if adr < 0x8000 {
		return r.rom[adr]
	}
	return 0x00
}

func (m RomOnly) GetCartType() string {
	return "Rom Only"
}
