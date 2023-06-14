package components

type Cartridge interface {
	Write(val uint8, adr uint16)
	Load(adr uint16) uint8
	GetCartridgeType() string
}

//ROM ONLY

//MBC 1

type MBC1 struct {
	rom, ram []uint8

	romBank, ramBank, maxRomBankCount, maxRamBankCount uint16

	ramEnabled, bankingMode, bankingModeAvailable bool
}

func CreateMBC1(rom []uint8) MBC1 {
	var romBankCount, ramBankCount uint16 = 2, 0
	//check if not ROMONLY
	if rom[0x0147] != 0 {
		romBankCount = getROMSize(rom[0x0148])
		ramBankCount = getRAMSize(rom[0x0149])
	}

	return MBC1{
		rom:                  rom,
		maxRomBankCount:      romBankCount,
		maxRamBankCount:      ramBankCount,
		ramEnabled:           false,
		bankingMode:          false,
		bankingModeAvailable: romBankCount > 32 && ramBankCount > 1}

}

func (c MBC1) Write(val uint8, adr uint16) {
	switch {
	case adr < 0x2000:
		c.ramEnabled = val&0b1111 == 0xA
	case adr < 0x4000:
		val = val & 0b1_1111
		if val&0b1111 == 0 {
			val |= 0b1
		}
		if c.maxRomBankCount <= 0b1111 {
			val &= 0b1111
		}
		c.romBank = (c.romBank & 0b1110_0000) | uint16(val)
	case adr < 0x6000:

	}

	if c.romBank == 0 { //TODO finish / properly implement MBC1
		c.romBank = 1
	}
}

func (c MBC1) Load(adr uint16) uint8 {
	if adr < 0x4000 {
		return c.rom[adr]
	} else if adr < 0x8000 {
		return c.rom[adr+(0x4000*c.ramBank)]
	}

	return 0x00
}
func (c MBC1) GetCartridgeType() string {
	return "MBC1"
}

// getROMSize returns the number of ROM banks based on the header byte 0x0148
func getROMSize(val uint8) uint16 {
	return 32 * (1 << val)
}

// getRAMSize returns the number of RAM bank based on the header byte 0x0147
func getRAMSize(val uint8) uint16 {
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
