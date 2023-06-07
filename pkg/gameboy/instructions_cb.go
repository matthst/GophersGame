package gameboy

func (gb *Gameboy) execCBInstr() int {

	opcode := gb.getImmediate() //fetch opcode

	var reg *uint8
	reg = new(uint8)

	switch opcode & 0xF {
	case 0x0, 0x8:
		reg = &gb.BC[1]
	case 0x1, 0x9:
		reg = &gb.BC[0]
	case 0x2, 0xA:
		reg = &gb.DE[1]
	case 0x3, 0xB:
		reg = &gb.DE[0]
	case 0x4, 0xC:
		reg = &gb.HL[1]
	case 0x5, 0xD:
		reg = &gb.HL[0]
	case 0x6, 0xE:
		*reg = gb.load(gb.getHL())
	case 0x7, 0xF:
		reg = &gb.AF[1]
	}

	switch i := opcode / 8; {
	case i == 0:
		gb.shiftCB(reg, gb.rotateLeftCircularInternal)
	case i == 1:
		gb.shiftCB(reg, gb.rotateRightCircularInternal)
	case i == 2:
		gb.shiftCB(reg, gb.rotateLeftInternal)
	case i == 3:
		gb.shiftCB(reg, gb.rotateRightInternal)
	case i == 4:
		gb.shiftCB(reg, gb.shiftLeftInternal)
	case i == 5:
		gb.shiftCB(reg, gb.shiftRightInternal)
	case i == 6:
		gb.shiftCB(reg, gb.swapInternal)
	case i == 7:
		gb.shiftCB(reg, gb.shiftRightMSBResetInternal)
	case i < 16:
		gb.setZFlagCB(reg, i-8)
	case i < 24:
		*reg = setBitCB(*reg, i-16)
	default:
		*reg = unsetBitCB(*reg, i-24)
	}

	if i := opcode & 0xF; i == 0x6 || i == 0xE {
		gb.write(*reg, gb.getHL())
		if opcode >= 0x40 && opcode < 0x7F {
			return 3
		}
		return 4
	}
	return 2
}

func (gb *Gameboy) shiftCB(reg *uint8, funcDef shiftInternalFuncDef) {
	*reg = funcDef(*reg)
	gb.setZFlag(*reg == 0)
}

func setBitCB(val, bit uint8) uint8 {
	return val | (0b0000_0001 << bit)
}

func unsetBitCB(val, bit uint8) uint8 {
	return val & (^(0b0000_0001 << bit))
}

func (gb *Gameboy) setZFlagCB(reg *uint8, bit uint8) {
	gb.setZFlag((*reg>>bit)&0x01 == 1)
	gb.setNFlag(false)
	gb.setHFlag(true)
}
