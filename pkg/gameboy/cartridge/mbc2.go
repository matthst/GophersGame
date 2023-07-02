package Cartridge

type MBC2 struct {
	rom, ram                 []uint8
	saveFilePath             string
	romBank, maxRomBankCount uint16
	ramEnabled, saveEnabled  bool
}

func CreateMBC2(rom []uint8, romPath string, batteryCart bool) *MBC2 {
	romBankCount := getROMSize(rom[0x0148])

	savePath := ""
	var ram []uint8
	if batteryCart {
		savePath, ram = createOrLoadSaveFile(romPath, 512)
	} else {
		ram = make([]uint8, 512)
	}

	return &MBC2{
		rom:             rom,
		ram:             ram,
		romBank:         1,
		maxRomBankCount: romBankCount,
		saveFilePath:    savePath,
		ramEnabled:      false,
		saveEnabled:     batteryCart}
}

func (m *MBC2) Write(val uint8, adr uint16) {
	if adr < 0x4000 {
		if adr&0b1_0000_0000 != 0 { // LSB of upper byte of adr word is set
			m.romBank = uint16(val & 0b1111)
			if m.romBank == 0 || m.romBank > m.maxRomBankCount {
				m.romBank = 1
			}
		} else {
			newRamEnabledVal := val == 0x0A
			if m.saveEnabled && m.ramEnabled && !newRamEnabledVal {
				writeToSaveFile(m.saveFilePath, m.ram)
			}
			m.ramEnabled = newRamEnabledVal
		}

	} else if m.ramEnabled && adr >= 0xA000 && adr < 0xC000 {
		m.ram[(adr-0xA000)&0b1_1111_1111] = val & 0b1111
	}
}

func (m *MBC2) Load(adr uint16) uint8 {
	switch {
	case adr < 0x4000:
		return m.rom[adr]
	case adr < 0x8000:
		return m.rom[adr+(0x4000*(m.romBank-1))]
	case adr >= 0xA000 && adr < 0xC000:
		return m.ram[(adr-0xA000)&0b1_1111_1111] & 0b1111
	default:
		return 0x00
	}
}

func (m *MBC2) GetCartType() string {
	return "MBC2"
}
