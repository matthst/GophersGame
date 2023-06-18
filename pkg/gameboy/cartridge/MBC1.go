package Cartridge

// MBC1 Used for MBC1 and RomOnly
type MBC1 struct {
	rom, ram []uint8

	saveFilePath string

	romBank, ramBank, maxRomBankCount, maxRamBankCount uint16

	ramEnabled, bankingMode, bankingModeAvailable bool
}

func CreateMBC1(rom []uint8, romPath string) MBC1 {
	var romBankCount, ramBankCount uint16 = 2, 0
	//check if not ROMONLY
	if rom[0x0147] != 0 {
		romBankCount = getROMSize(rom[0x0148])
		ramBankCount = getRAMSize(rom[0x0149])
	}
	savePath := ""
	var ram []uint8
	if ramBankCount > 0 {
		savePath, ram = createOrLoadSaveFile(romPath, int(ramBankCount)*0x2000)
	}

	return MBC1{
		rom:                  rom,
		ram:                  ram,
		maxRomBankCount:      romBankCount,
		maxRamBankCount:      ramBankCount,
		ramEnabled:           false,
		bankingMode:          false,
		saveFilePath:         savePath,
		bankingModeAvailable: romBankCount > 32 && ramBankCount > 1}

}

func (c MBC1) Write(val uint8, adr uint16) {
	if adr < 0x2000 { // RAM Enable
		newRamEnableVal := val&0b1111 == 0xA
		if c.maxRamBankCount > 0 && !c.ramEnabled && newRamEnableVal {
			writeToSaveFile(c.saveFilePath, c.ram)
		}
		c.ramEnabled = newRamEnableVal
	} else if adr < 0x4000 {
		val = val & 0b1_1111
		if val&0b1111 == 0 { // the lower 4 bytes cannot be 0x0, they will always be at least 0x1
			val |= 0b1
		}
		if c.maxRomBankCount <= 0b1111 { // the MBC will not allow selecting a bank that does not exist
			val &= 0b1111
		}
		c.romBank = (c.romBank & 0b1110_0000) | uint16(val)
	} else if adr < 0x6000 {
		bits := uint16(val & 0b11)
		c.ramBank = bits
		c.romBank = c.romBank&0b1_1111 + (bits << 5)
	} else if adr < 0x8000 {
		c.bankingMode = val&0b1 == 1 // TODO implement bankingMode
	}
}

func (c MBC1) Load(adr uint16) uint8 {
	if adr < 0x4000 {
		return c.rom[adr]
	} else if adr < 0x8000 {
		return c.rom[adr+(0x4000*c.romBank)]
	} else if c.ramEnabled && adr >= 0xA000 && adr < 0xC000 {
		return c.ram[adr-0xA000+c.ramBank*0x2000]
	}
	return 0x00
}
func (c MBC1) GetCartridgeType() string {
	return "MBC1"
}
