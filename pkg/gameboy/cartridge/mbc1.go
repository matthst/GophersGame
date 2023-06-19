package Cartridge

// MBC1 Used for MBC1 and RomOnly
type MBC1 struct {
	rom, ram                    []uint8
	romBank1, romBank2, ramBank uint32
	saveFilePath                string

	ramEnabled, ramCart, saveEnabled, bankingMode, bankingModeAvailable bool
}

func CreateMBC1(rom []uint8, romPath string, ramCart, batteryCart bool) MBC1 {
	romBankCount, ramBankCount := getROMSize(rom[0x0148]), getRAMSize(rom[0x0149])

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
		romBank1:             1 << 14,
		ramCart:              ramBankCount > 0,
		ramEnabled:           false,
		saveEnabled:          batteryCart,
		bankingMode:          false,
		saveFilePath:         savePath,
		bankingModeAvailable: romBankCount > 32 && ramBankCount > 1}
}

func (m MBC1) Write(val uint8, adr uint16) {

	switch {
	case m.ramCart && adr < 0x2000: // RAM Enable
		newRamEnableVal := val&0b1111 == 0xA
		if !m.ramEnabled && newRamEnableVal && m.saveEnabled {
			writeToSaveFile(m.saveFilePath, m.ram)
		}
		m.ramEnabled = newRamEnableVal

	case adr < 0x4000: // Lower 5 bits of ROM Bank Select
		val &= 0b1_1111
		if val == 0 { // the value cannot be 0x0, they will always be at least 0x1
			val = 1
		}
		m.romBank1 = uint32(val) << 14

	case adr < 0x6000: // Ram Bank Select and 2MSB from Rom Bank select
		bits := uint16(val & 0b11)
		m.ramBank = uint32(bits) << 13
		m.romBank2 = uint32(bits) << 19

	case adr < 0x8000:
		m.bankingMode = val&0b1 == 1 && m.bankingModeAvailable

	case m.ramCart && m.ramEnabled && adr >= 0xA000 && adr < 0xC000: //RAM Write
		if m.bankingMode {
			m.ram[m.ramBank|(uint32(adr)&0x1FFF)] = val
		}
		m.ram[adr&0x1FFF] = val
	}
}

func (m MBC1) Load(adr uint16) uint8 {
	switch {
	case adr < 0x4000:
		if m.bankingMode {
			return m.rom[m.romBank2|uint32(adr&0x3FFF)]
		}
		return m.rom[adr]
	case adr < 0x8000:
		return m.rom[m.romBank2|m.romBank1|(uint32(adr)&0x3FFF)]
	case m.ramCart && m.ramEnabled && adr >= 0xA000 && adr < 0xC000:
		if m.bankingMode {
			return m.ram[m.ramBank|(uint32(adr)&0x1FFF)]
		}
		return m.ram[adr&0x1FFF]
	default:
		return 0x00
	}
}

func (m MBC1) GetCartType() string {
	return "MBC1"
}
