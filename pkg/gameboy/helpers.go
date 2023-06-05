package gameboy

// get the combined register
func getWord(r *[2]uint8) uint16 {
	val := uint16(r[0]) | (uint16(r[1]) << 8)
	return val
}

// get the combined register
func getWordFromBytes(hi, lo uint8) uint16 {
	val := uint16(lo) | (uint16(hi) << 8)
	return val
}

func setWord(r *[2]uint8, val uint16) {
	r[0] = uint8(val)
	r[1] = uint8(val >> 8)
}

func getWordInc(r *[2]uint8) uint16 {
	val := getWord(r)
	setWord(r, val+1)
	return val
}

func getWordDec(r *[2]uint8) uint16 {
	val := getWord(r)
	setWord(r, val-1)
	return val
}

// aluR8Def function definition for the internal helpers used for the 0xCB shift functions
type shiftInternalFuncDef func(uint8) uint8

func (regs *registers) rotateLeftInternal(val uint8) uint8 {
	result := (val << 1) + regs.getCarryValue()
	regs.AF[0] = 0b0000_0000 + ((val >> 7) << 4)
	return result
}

func (regs *registers) rotateLeftCircularInternal(val uint8) uint8 {
	regs.AF[0] = 0b0000_0000 + ((val >> 7) << 4)
	return (val << 1) + (val >> 7)
}

func (regs *registers) rotateRightInternal(val uint8) uint8 {
	result := (val >> 1) + (regs.getCarryValue() << 7)
	regs.AF[0] = 0b0000_0000 + ((val & 0x1) << 4)
	return result
}

func (regs *registers) rotateRightCircularInternal(val uint8) uint8 {
	regs.AF[0] = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) + (val << 7)
}

func (regs *registers) shiftLeftInternal(val uint8) uint8 {
	regs.AF[0] = 0b0000_0000 + ((val >> 7) << 4)
	return val << 1
}

func (regs *registers) shiftRightInternal(val uint8) uint8 {
	regs.AF[0] = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) | (val & 0b1000_0000)
}

func (regs *registers) shiftRightMSBResetInternal(val uint8) uint8 {
	regs.AF[0] = 0b0000_0000 + ((val & 0x1) << 4)
	return val >> 1
}

func (regs *registers) swapInternal(val uint8) uint8 {
	regs.AF[0] = 0
	regs.setZ(val == 0)
	return (val >> 4) + (val << 4)
}

func halfCarryAddCheck8Bit(a, b uint8) bool {
	return (((a & 0xf) + (b & 0xf)) & 0x10) == 0x10
}

func halfCarrySubCheck8Bit(a, b uint8) bool {
	return (((a & 0xf) - (b & 0xf)) & 0x10) == 0x10
}

func halfCarryAddCheck16Bit(a, b uint16) bool {
	return (((a & 0xf) + (b & 0xf)) & 0x10) == 0x10
}
