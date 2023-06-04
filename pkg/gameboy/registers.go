package gameboy

// for the size-2 arrays, the first val is lo, the second is hi
type registers struct {
	PC                 uint16
	AF, BC, DE, HL, SP [2]uint8
}

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

func (regs *registers) getZ() bool {
	return regs.AF[0]&0b1000_0000 != 0
}

func (regs *registers) setZ(val bool) {
	if val {
		regs.AF[0] |= 0b1000_0000
	} else {
		regs.AF[0] |= 0b0111_1111
	}
}

func (regs *registers) getN() bool {
	return regs.AF[0]&0b0100_0000 != 0
}

func (regs *registers) setN(val bool) {
	if val {
		regs.AF[0] |= 0b0100_0000
	} else {
		regs.AF[0] |= 0b1011_1111
	}
}

func (regs *registers) getH() bool {
	return regs.AF[0]&0b0010_0000 != 0
}

func (regs *registers) setH(val bool) {
	if val {
		regs.AF[0] |= 0b0010_0000
	} else {
		regs.AF[0] |= 0b1101_1111
	}
}

func (regs *registers) getC() bool {
	return regs.AF[0]&0b0001_0000 != 0
}

func (regs *registers) setC(val bool) {
	if val {
		regs.AF[0] |= 0b0001_0000
	} else {
		regs.AF[0] |= 0b1110_1111
	}
}

func (regs *registers) getCarryValue() uint8 {
	if regs.getC() {
		return 1
	}
	return 0
}
