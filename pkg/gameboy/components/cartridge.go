package components

type Cartridge interface {
	Write(val uint8, adr uint16)
	Load(adr uint16) uint8
	GetCartridgeType() string
}

//ROM ONLY

type RomOnly struct {
	rom []uint8
}

func CreateRomOnly(rom []uint8) RomOnly {
	return RomOnly{rom: rom}
}

func (c RomOnly) Write(val uint8, adr uint16) {
}

func (c RomOnly) Load(adr uint16) uint8 {
	if adr < 0x8000 {
		return c.rom[adr]
	}
	return 0x00
}

func (c RomOnly) GetCartridgeType() string {
	return "No MBC"
}

//MBC 1

type MBC1 struct {
	rom, ram []uint8

	romBank, ramBank, maxRomBank, maxRamBank int

	ramEnabled bool
}

func createMBC1(rom []uint8) MBC1 {
	romSize := getROMSize(rom[0x0148])
	ramSize := getRAMSize(rom[0x0149])

	return MBC1{
		rom:        rom,
		maxRomBank: romSize,
		maxRamBank: ramSize,
		ramEnabled: false,
	}

}

func (c MBC1) Write(val uint8, adr uint16) {

}

func (c MBC1) Load(adr uint16) uint8 {
	if adr < 0x4000 {
		return c.rom[adr]
	}

	return 0x00
}
func (c MBC1) GetCartridgeType() string {
	return "MBC1"
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
