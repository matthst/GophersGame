package gameboy

func (gb *Gameboy) execCBInstr() int {

	opcode := gb.getImmediate() //fetch opcode

	var reg *uint8
	reg = new(uint8)

	switch opcode & 0xF {
	case 0x0, 0x8:
		reg = &gb.Regs.BC[1]
	case 0x1, 0x9:
		reg = &gb.Regs.BC[0]
	case 0x2, 0xA:
		reg = &gb.Regs.DE[1]
	case 0x3, 0xB:
		reg = &gb.Regs.DE[0]
	case 0x4, 0xC:
		reg = &gb.Regs.HL[1]
	case 0x5, 0xD:
		reg = &gb.Regs.HL[0]
	case 0x6, 0xE:
		*reg = gb.load(gb.Regs.getHL())
	case 0x7, 0xF:
		reg = &gb.Regs.AF[1]
	}

	switch i := opcode / 8; {
	case i == 0:
		gb.Regs.shiftCB(reg, gb.Regs.rotateLeftCircularInternal)
	case i == 1:
		gb.Regs.shiftCB(reg, gb.Regs.rotateRightCircularInternal)
	case i == 2:
		gb.Regs.shiftCB(reg, gb.Regs.rotateLeftInternal)
	case i == 3:
		gb.Regs.shiftCB(reg, gb.Regs.rotateRightInternal)
	case i == 4:
		gb.Regs.shiftCB(reg, gb.Regs.shiftLeftInternal)
	case i == 5:
		gb.Regs.shiftCB(reg, gb.Regs.shiftRightInternal)
	case i == 6:
		gb.Regs.shiftCB(reg, gb.Regs.swapInternal)
	case i == 7:
		gb.Regs.shiftCB(reg, gb.Regs.shiftRightMSBResetInternal)
	case i < 16:
		gb.Regs.setZFlagCB(reg, i-8)
	case i < 24:
		*reg = setBitCB(*reg, i-16)
	default:
		*reg = unsetBitCB(*reg, i-24)
	}

	if i := opcode & 0xF; i == 0x6 || i == 0xE {
		gb.write(*reg, gb.Regs.getHL())
		if opcode >= 0x40 && opcode < 0x7F {
			return 3
		}
		return 4
	}
	return 2
}

func (regs *registers) shiftCB(reg *uint8, funcDef shiftInternalFuncDef) {
	*reg = funcDef(*reg)
	regs.setZ(*reg == 0)
}

func setBitCB(val, bit uint8) uint8 {
	return val | (0b0000_0001 << bit)
}

func unsetBitCB(val, bit uint8) uint8 {
	return val & (^(0b0000_0001 << bit))
}

func (regs *registers) setZFlagCB(reg *uint8, bit uint8) {
	regs.setZ((*reg>>bit)&0x01 == 1)
	regs.setN(false)
	regs.setH(true)
}
