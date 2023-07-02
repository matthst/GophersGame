package Cartridge

// MBC1 Used for MBC1 and RomOnly
type MBC1 struct {
	rom, ram                                 []uint8
	romBank1, romBank2, ramBank, maxRomBanks uint32
	saveFilePath                             string
	uint8
	ramEnabled, ramCart, saveEnabled bool
}

func CreateMBC1(rom []uint8, romPath string, ramCart, batteryCart bool) *MBC1 {
	romBankCount := uint32(len(rom) / 0x4000)
	ramBankCount := getRAMSize(rom[0x0149])

	savePath := ""
	var ram []uint8
	ramSize := int(ramBankCount) * 0x2000
	if batteryCart {
		savePath, ram = createOrLoadSaveFile(romPath, ramSize)
	} else {
		ram = make([]uint8, ramSize)
	}

	return &MBC1{
		rom:          rom,
		ram:          ram,
		romBank1:     1 << 14,
		maxRomBanks:  romBankCount,
		ramCart:      ramCart,
		ramEnabled:   false,
		saveEnabled:  batteryCart,
		saveFilePath: savePath}
}

func (m *MBC1) Write(val uint8, adr uint16) {
	switch {
	case adr < 0x2000: // RAM Enable
		newRamEnableVal := val&0b1111 == 0xA
		if !m.ramEnabled && newRamEnableVal && m.saveEnabled {
			writeToSaveFile(m.saveFilePath, m.ram)
		}
		m.ramEnabled = newRamEnableVal && m.ramCart

	case adr < 0x4000: // Lower 5 bits of ROM Bank Select
		val &= 0b1_1111
		if val == 0 { // the value cannot be 0x0, they will always be at least 0x1
			val = 1
		}
		m.romBank1 = (uint32(val) % m.maxRomBanks) << 14

	case adr < 0x6000: // Ram Bank Select and 2MSB from Rom Bank select
		val %= 0b11
		m.ramBank = uint32(val) << 13
		m.romBank2 = ((uint32(val) << 5) % m.maxRomBanks) << 14

	case m.ramEnabled && adr >= 0xA000 && adr < 0xC000: //RAM Write
		m.ram[m.ramBank|(uint32(adr)&0x1FFF)] = val
	}
}

func (m *MBC1) Load(adr uint16) uint8 {
	switch {
	case adr < 0x4000:
		return m.rom[adr]
	case adr < 0x8000:
		return m.rom[m.romBank2|m.romBank1|(uint32(adr&0x3FFF))]
	case m.ramEnabled && adr >= 0xA000 && adr < 0xC000:
		return m.ram[m.ramBank|(uint32(adr&0x1FFF))]
	default:
		return 0x00 //0b 0100 0000 0000 0000
	}
}

func (m *MBC1) GetCartType() string {
	return "MBC1"
}
