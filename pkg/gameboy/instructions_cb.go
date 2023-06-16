package gameboy

func execCBInstr() {

	opcode := getImmediate() //fetch opcode

	var reg *uint8
	reg = new(uint8)

	switch opcode & 0xF {
	case 0x0, 0x8:
		reg = &bReg
	case 0x1, 0x9:
		reg = &cReg
	case 0x2, 0xA:
		reg = &dReg
	case 0x3, 0xB:
		reg = &eReg
	case 0x4, 0xC:
		reg = &hReg
	case 0x5, 0xD:
		reg = &lReg
	case 0x6, 0xE:
		*reg = memConLoad(getHL())
	case 0x7, 0xF:
		reg = &aReg
	}

	switch i := opcode / 8; {
	case i == 0:
		shiftCB(reg, rotateLeftCircularInternal)
	case i == 1:
		shiftCB(reg, rotateRightCircularInternal)
	case i == 2:
		shiftCB(reg, rotateLeftInternal)
	case i == 3:
		shiftCB(reg, rotateRightInternal)
	case i == 4:
		shiftCB(reg, shiftLeftInternal)
	case i == 5:
		shiftCB(reg, shiftRightInternal)
	case i == 6:
		shiftCB(reg, swapInternal)
	case i == 7:
		shiftCB(reg, shiftRightMSBResetInternal)
	case i < 16:
		setZFlagCB(reg, i-8)
	case i < 24:
		*reg = unsetBitCB(*reg, i-16)
	default:
		*reg = setBitCB(*reg, i-24)
	}

	if i := opcode & 0xF; (opcode < 0x40 || opcode > 0x7F) && (i == 0x6 || i == 0xE) {
		memConWrite(*reg, getHL())
	}
}

func shiftCB(reg *uint8, funcDef shiftInternalFuncDef) {
	*reg = funcDef(*reg)
	setZFlag(*reg == 0)
}

func setBitCB(val, bit uint8) uint8 {
	return val | (0b0000_0001 << bit)
}

func unsetBitCB(val, bit uint8) uint8 {
	return val & (^(0b0000_0001 << bit))
}

func setZFlagCB(reg *uint8, bit uint8) {
	setZFlag((*reg>>bit)&0x01 == 0)
	setNFlag(false)
	setHFlag(true)
}
