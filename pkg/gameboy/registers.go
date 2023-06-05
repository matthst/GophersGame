package gameboy

// for the size-2 arrays, the first val is lo, the second is hi
type registers struct {
	PC                 uint16
	AF, BC, DE, HL, SP [2]uint8
}

func (regs *registers) getHL() uint16 {
	return getWord(&regs.HL)
}

func (regs *registers) setFlags(Z, N, H, C bool) {
	regs.setZ(Z)
	regs.setN(N)
	regs.setH(H)
	regs.setC(C)
}

func (regs *registers) getZ() bool {
	return regs.AF[0]&0b1000_0000 != 0
}

func (regs *registers) setZ(val bool) {
	if val {
		regs.AF[0] |= 0b1000_0000
	} else {
		regs.AF[0] &= 0b0111_1111
	}
}

func (regs *registers) getN() bool {
	return regs.AF[0]&0b0100_0000 != 0
}

func (regs *registers) setN(val bool) {
	if val {
		regs.AF[0] |= 0b0100_0000
	} else {
		regs.AF[0] &= 0b1011_1111
	}
}

func (regs *registers) getH() bool {
	return regs.AF[0]&0b0010_0000 != 0
}

func (regs *registers) setH(val bool) {
	if val {
		regs.AF[0] |= 0b0010_0000
	} else {
		regs.AF[0] &= 0b1101_1111
	}
}

func (regs *registers) getC() bool {
	return regs.AF[0]&0b0001_0000 != 0
}

func (regs *registers) setC(val bool) {
	if val {
		regs.AF[0] |= 0b0001_0000
	} else {
		regs.AF[0] &= 0b1110_1111
	}
}

func (regs *registers) getCarryValue() uint8 {
	if regs.getC() {
		return 1
	}
	return 0
}
