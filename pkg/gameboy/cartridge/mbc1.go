package Cartridge

// MBC1 Used for MBC1 and RomOnly
type MBC1 struct {
	rom, ram []uint8

	saveFilePath string

	romBank, ramBank, maxRomBankCount, maxRamBankCount uint16

	ramEnabled, saveEnabled, bankingMode, bankingModeAvailable bool
}

func CreateMBC1(rom []uint8, romPath string, ramCart, batteryCart bool) MBC1 {
	var romBankCount, ramBankCount uint16 = 2, 0
	//check if not ROMONLY
	if rom[0x0147] != 0 {
		romBankCount = getROMSize(rom[0x0148])
		ramBankCount = getRAMSize(rom[0x0149])
	}
	if !ramCart {
		ramBankCount = 0
	}
	savePath := ""
	var ram []uint8
	ramSize := int(ramBankCount) * 0x2000
	if batteryCart {
		savePath, ram = createOrLoadSaveFile(romPath, ramSize)
	} else {
		ram = make([]uint8, ramSize)
	}

	return MBC1{
		rom:                  rom,
		ram:                  ram,
		romBank:              1,
		maxRomBankCount:      romBankCount,
		maxRamBankCount:      ramBankCount,
		ramEnabled:           false,
		saveEnabled:          batteryCart,
		bankingMode:          false,
		saveFilePath:         savePath,
		bankingModeAvailable: romBankCount > 32 && ramBankCount > 1}
}

func (m MBC1) Write(val uint8, adr uint16) {

	switch {
	case adr < 0x2000: // RAM Enable
		newRamEnableVal := val&0b1111 == 0xA
		if m.maxRamBankCount > 0 && !m.ramEnabled && newRamEnableVal && m.saveEnabled {
			writeToSaveFile(m.saveFilePath, m.ram)
		}
		m.ramEnabled = newRamEnableVal

	case adr < 0x4000: // Lower 5 bits of ROM Bank Select
		val = val & 0b1_1111
		if val&0b1111 == 0 { // the lower 4 bytes cannot be 0x0, they will always be at least 0x1
			val |= 0b1
		}
		if m.maxRomBankCount <= 0b1111 { // the MBC will not allow selecting a bank that does not exist
			val &= 0b1111
		}
		m.romBank = (m.romBank & 0b1110_0000) | uint16(val)

	case adr < 0x6000: // Ram Bank Select and 2MSB from Rom Bank select
		bits := uint16(val & 0b11)
		m.ramBank = bits
		m.romBank = m.romBank&0b1_1111 + (bits << 5)

	case adr < 0x8000:
		m.bankingMode = val&0b1 == 1 // TODO implement bankingMode

	case m.ramEnabled && adr >= 0xA000 && adr < 0xC000: //RAM Write
		m.ram[adr-0xA000+m.ramBank*0x2000] = val

	}
}

func (m MBC1) Load(adr uint16) uint8 {
	switch {
	case adr < 0x4000:
		return m.rom[adr]
	case adr < 0x8000:
		return m.rom[adr+(0x4000*(m.romBank-1))]
	case m.ramEnabled && adr >= 0xA000 && adr < 0xC000:
		return m.ram[adr-0xA000+m.ramBank*0x2000]
	default:
		return 0x00
	}
}
