package gameboy

// get the combined register
func getWord(r *[2]uint8) uint16 {
	val := uint16(r[0]) | (uint16(r[1]) << 8)
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

func (regs *registers) rotateLeftInternal(reg *uint8) {
	val := (*reg << 1) + regs.getCarryValue()
	regs.AF[0] = 0b0000_0000 + ((*reg >> 7) << 4)
	*reg = val
}

func (regs *registers) rotateLeftCircularInternal(reg *uint8) {
	regs.AF[0] = 0b0000_0000 + ((*reg >> 7) << 4)
	*reg = (*reg << 1) + (*reg >> 7)
}

func (regs *registers) rotateRightInternal(reg *uint8) {
	val := (*reg >> 1) + (regs.getCarryValue() << 7)
	regs.AF[0] = 0b0000_0000 + ((*reg & 0x1) << 4)
	*reg = val
}

func (regs *registers) rotateRightCircularInternal(reg *uint8) {
	regs.AF[0] = 0b0000_0000 + ((*reg & 0x1) << 4)
	*reg = (*reg >> 1) + (*reg << 7)
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
