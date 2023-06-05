package gameboy

func (gb *gameboy) execNextInstr() int {

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */

	opcode := gb.getImmediate() //fetch opcode

	switch opcode {
	case 0x00:
		return nop()
	case 0x10:
		return stop()
	case 0x20:
		return gb.JumpRelativeI8(!gb.regs.getZ())
	case 0x30:
		return gb.JumpRelativeI8(!gb.regs.getN())
	case 0x01:
		return gb.loadI16(&gb.regs.BC)
	case 0x11:
		return gb.loadI16(&gb.regs.DE)
	case 0x21:
		return gb.loadI16(&gb.regs.HL)
	case 0x31:
		return gb.loadI16SP()
	case 0x02:
		return gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.BC))
	case 0x12:
		return gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.DE))
	case 0x22:
		return gb.storeR8(gb.regs.AF[1], getWordInc(&gb.regs.HL))
	case 0x32:
		return gb.storeR8(gb.regs.AF[1], getWordDec(&gb.regs.HL))
	case 0x03:
		return incR16(&gb.regs.BC)
	case 0x13:
		return incR16(&gb.regs.DE)
	case 0x23:
		return incR16(&gb.regs.HL)
	case 0x33:
		return gb.regs.incSP()
	case 0x04:
		return gb.regs.incR8(&gb.regs.BC[1])
	case 0x14:
		return gb.regs.incR8(&gb.regs.DE[1])
	case 0x24:
		return gb.regs.incR8(&gb.regs.HL[1])
	case 0x34:
		return gb.incM8(getWord(&gb.regs.HL))
	case 0x05:
		return gb.regs.decR8(&gb.regs.BC[1])
	case 0x15:
		return gb.regs.decR8(&gb.regs.DE[1])
	case 0x25:
		return gb.regs.decR8(&gb.regs.HL[1])
	case 0x35:
		return gb.decM8(getWord(&gb.regs.HL))
	case 0x06:
		return gb.loadI8(&gb.regs.BC[1])
	case 0x16:
		return gb.loadI8(&gb.regs.DE[1])
	case 0x26:
		return gb.loadI8(&gb.regs.HL[1])
	case 0x36:
		return gb.storeI8(&gb.regs.HL)
	case 0x07:
		return gb.regs.rotateLeftCircularA()
	case 0x17:
		return gb.regs.rotateLeftA()
	case 0x27:
		return gb.regs.decimalAdjustA()
	case 0x37:
		return gb.regs.setCarryFlag(true)
	case 0x08:
		return gb.storeSPI16()
	case 0x18:
		return gb.JumpRelativeI8(true)
	case 0x28:
		return gb.JumpRelativeI8(gb.regs.getZ())
	case 0x38:
		return gb.JumpRelativeI8(gb.regs.getC())
	case 0x09:
		return gb.regs.addR16R16(&gb.regs.HL, getWord(&gb.regs.BC))
	case 0x19:
		return gb.regs.addR16R16(&gb.regs.HL, getWord(&gb.regs.DE))
	case 0x29:
		return gb.regs.addR16R16(&gb.regs.HL, getWord(&gb.regs.HL))
	case 0x39:
		return gb.regs.addR16R16(&gb.regs.HL, gb.regs.SP)
	case 0x0A:
		return gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.BC))
	case 0x1A:
		return gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.DE))
	case 0x2A:
		return gb.loadR8(&gb.regs.AF[1], getWordInc(&gb.regs.HL))
	case 0x3A:
		return gb.loadR8(&gb.regs.AF[1], getWordDec(&gb.regs.HL))
	case 0x0B:
		return decR16(&gb.regs.BC)
	case 0x1B:
		return decR16(&gb.regs.DE)
	case 0x2B:
		return decR16(&gb.regs.HL)
	case 0x3B:
		return gb.regs.decSP()
	case 0x0C:
		return gb.regs.incR8(&gb.regs.BC[0])
	case 0x1C:
		return gb.regs.incR8(&gb.regs.DE[0])
	case 0x2C:
		return gb.regs.incR8(&gb.regs.HL[0])
	case 0x3C:
		return gb.regs.incR8(&gb.regs.AF[1])
	case 0x0D:
		return gb.regs.decR8(&gb.regs.BC[0])
	case 0x1D:
		return gb.regs.decR8(&gb.regs.DE[0])
	case 0x2D:
		return gb.regs.decR8(&gb.regs.HL[0])
	case 0x3D:
		return gb.regs.decR8(&gb.regs.AF[1])
	case 0x0E:
		return gb.loadI8(&gb.regs.BC[0])
	case 0x1E:
		return gb.loadI8(&gb.regs.DE[0])
	case 0x2E:
		return gb.loadI8(&gb.regs.HL[0])
	case 0x3E:
		return gb.loadI8(&gb.regs.AF[1])
	case 0x0F:
		return gb.regs.rotateRightCircularA()
	case 0x1F:
		return gb.regs.rotateRightA()
	case 0x2F:
		return complementR8(&gb.regs.AF[1])
	case 0x3F:
		return gb.regs.setCarryFlag(!gb.regs.getC())
	case 0x40:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.BC[1])
	case 0x50:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.BC[1])
	case 0x60:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.BC[1])
	case 0x70:
		return gb.storeR8(gb.regs.BC[1], getWord(&gb.regs.HL))
	case 0x41:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.BC[0])
	case 0x51:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.BC[0])
	case 0x61:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.BC[0])
	case 0x71:
		return gb.storeR8(gb.regs.BC[0], getWord(&gb.regs.HL))
	case 0x42:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.DE[1])
	case 0x52:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.DE[1])
	case 0x62:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.DE[1])
	case 0x72:
		return gb.storeR8(gb.regs.DE[1], getWord(&gb.regs.HL))
	case 0x43:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.DE[0])
	case 0x53:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.DE[0])
	case 0x63:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.DE[0])
	case 0x73:
		return gb.storeR8(gb.regs.DE[0], getWord(&gb.regs.HL))
	case 0x44:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.HL[1])
	case 0x54:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.HL[1])
	case 0x64:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.HL[1])
	case 0x74:
		return gb.storeR8(gb.regs.HL[1], getWord(&gb.regs.HL))
	case 0x45:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.HL[0])
	case 0x55:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.HL[0])
	case 0x65:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.HL[0])
	case 0x75:
		return gb.storeR8(gb.regs.HL[0], getWord(&gb.regs.HL))
	case 0x46:
		return gb.loadR8(&gb.regs.BC[1], getWord(&gb.regs.HL))
	case 0x56:
		return gb.loadR8(&gb.regs.DE[1], getWord(&gb.regs.HL))
	case 0x66:
		return gb.loadR8(&gb.regs.HL[1], getWord(&gb.regs.HL))
	case 0x76:
		return halt()
	case 0x47:
		return loadR8R8(&gb.regs.BC[1], &gb.regs.AF[1])
	case 0x57:
		return loadR8R8(&gb.regs.DE[1], &gb.regs.AF[1])
	case 0x67:
		return loadR8R8(&gb.regs.HL[1], &gb.regs.AF[1])
	case 0x77:
		return gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.HL))
	case 0x48:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.BC[1])
	case 0x58:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.BC[1])
	case 0x68:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.BC[1])
	case 0x78:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.BC[1])
	case 0x49:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.BC[0])
	case 0x59:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.BC[0])
	case 0x69:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.BC[0])
	case 0x79:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.BC[0])
	case 0x4A:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.DE[1])
	case 0x5A:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.DE[1])
	case 0x6A:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.DE[1])
	case 0x7A:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.DE[1])
	case 0x4B:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.DE[0])
	case 0x5B:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.DE[0])
	case 0x6B:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.DE[0])
	case 0x7B:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.DE[0])
	case 0x4C:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.HL[1])
	case 0x5C:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.HL[1])
	case 0x6C:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.HL[1])
	case 0x7C:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.HL[1])
	case 0x4D:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.HL[0])
	case 0x5D:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.HL[0])
	case 0x6D:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.HL[0])
	case 0x7D:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.HL[0])
	case 0x4E:
		return gb.loadR8(&gb.regs.BC[0], getWord(&gb.regs.HL))
	case 0x5E:
		return gb.loadR8(&gb.regs.DE[0], getWord(&gb.regs.HL))
	case 0x6E:
		return gb.loadR8(&gb.regs.HL[0], getWord(&gb.regs.HL))
	case 0x7E:
		return gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.HL))
	case 0x4F:
		return loadR8R8(&gb.regs.BC[0], &gb.regs.AF[1])
	case 0x5F:
		return loadR8R8(&gb.regs.DE[0], &gb.regs.AF[1])
	case 0x6F:
		return loadR8R8(&gb.regs.HL[0], &gb.regs.AF[1])
	case 0x7F:
		return loadR8R8(&gb.regs.AF[1], &gb.regs.AF[1])
	case 0x80:
		return gb.regs.addR8(gb.regs.BC[1])
	case 0x90:
		return gb.regs.subR8(gb.regs.BC[1])
	case 0xA0:
		return gb.regs.andR8(gb.regs.BC[1])
	case 0xB0:
		return gb.regs.orR8(gb.regs.BC[1])
	case 0x81:
		return gb.regs.addR8(gb.regs.BC[0])
	case 0x91:
		return gb.regs.subR8(gb.regs.BC[0])
	case 0xA1:
		return gb.regs.andR8(gb.regs.BC[0])
	case 0xB1:
		return gb.regs.orR8(gb.regs.BC[0])
	case 0x82:
		return gb.regs.addR8(gb.regs.DE[1])
	case 0x92:
		return gb.regs.subR8(gb.regs.DE[1])
	case 0xA2:
		return gb.regs.andR8(gb.regs.DE[1])
	case 0xB2:
		return gb.regs.orR8(gb.regs.DE[1])
	case 0x83:
		return gb.regs.addR8(gb.regs.DE[0])
	case 0x93:
		return gb.regs.subR8(gb.regs.DE[0])
	case 0xA3:
		return gb.regs.andR8(gb.regs.DE[0])
	case 0xB3:
		return gb.regs.orR8(gb.regs.DE[0])
	case 0x84:
		return gb.regs.addR8(gb.regs.HL[1])
	case 0x94:
		return gb.regs.subR8(gb.regs.HL[1])
	case 0xA4:
		return gb.regs.andR8(gb.regs.HL[1])
	case 0xB4:
		return gb.regs.orR8(gb.regs.HL[1])
	case 0x85:
		return gb.regs.addR8(gb.regs.HL[0])
	case 0x95:
		return gb.regs.subR8(gb.regs.HL[0])
	case 0xA5:
		return gb.regs.andR8(gb.regs.HL[0])
	case 0xB5:
		return gb.regs.orR8(gb.regs.HL[0])
	case 0x86:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.addR8)
	case 0x96:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.subR8)
	case 0xA6:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.andR8)
	case 0xB6:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.orR8)
	case 0x87:
		return gb.regs.addR8(gb.regs.AF[1])
	case 0x97:
		return gb.regs.subR8(gb.regs.AF[1])
	case 0xA7:
		return gb.regs.andR8(gb.regs.AF[1])
	case 0xB7:
		return gb.regs.orR8(gb.regs.AF[1])
	case 0x88:
		return gb.regs.adcR8(gb.regs.BC[1])
	case 0x98:
		return gb.regs.sbcR8(gb.regs.BC[1])
	case 0xA8:
		return gb.regs.xorR8(gb.regs.BC[1])
	case 0xB8:
		return gb.regs.cpR8(gb.regs.BC[1])
	case 0x89:
		return gb.regs.adcR8(gb.regs.BC[0])
	case 0x99:
		return gb.regs.sbcR8(gb.regs.BC[0])
	case 0xA9:
		return gb.regs.xorR8(gb.regs.BC[0])
	case 0xB9:
		return gb.regs.cpR8(gb.regs.BC[0])
	case 0x8A:
		return gb.regs.adcR8(gb.regs.DE[1])
	case 0x9A:
		return gb.regs.sbcR8(gb.regs.DE[1])
	case 0xAA:
		return gb.regs.xorR8(gb.regs.DE[1])
	case 0xBA:
		return gb.regs.cpR8(gb.regs.DE[1])
	case 0x8B:
		return gb.regs.adcR8(gb.regs.DE[0])
	case 0x9B:
		return gb.regs.sbcR8(gb.regs.DE[0])
	case 0xAB:
		return gb.regs.xorR8(gb.regs.DE[0])
	case 0xBB:
		return gb.regs.cpR8(gb.regs.DE[0])
	case 0x8C:
		return gb.regs.adcR8(gb.regs.HL[1])
	case 0x9C:
		return gb.regs.sbcR8(gb.regs.HL[1])
	case 0xAC:
		return gb.regs.xorR8(gb.regs.HL[1])
	case 0xBC:
		return gb.regs.cpR8(gb.regs.HL[1])
	case 0x8D:
		return gb.regs.adcR8(gb.regs.HL[0])
	case 0x9D:
		return gb.regs.sbcR8(gb.regs.HL[0])
	case 0xAD:
		return gb.regs.xorR8(gb.regs.HL[0])
	case 0xBD:
		return gb.regs.cpR8(gb.regs.HL[0])
	case 0x8E:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.adcR8)
	case 0x9E:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.sbcR8)
	case 0xAE:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.xorR8)
	case 0xBE:
		return gb.aluM8(getWord(&gb.regs.HL), gb.regs.cpR8)
	case 0x8F:
		return gb.regs.adcR8(gb.regs.AF[1])
	case 0x9F:
		return gb.regs.sbcR8(gb.regs.AF[1])
	case 0xAF:
		return gb.regs.xorR8(gb.regs.AF[1])
	case 0xBF:
		return gb.regs.cpR8(gb.regs.AF[1])
	case 0xC0:
		return gb.ret(!gb.regs.getZ())
	case 0xD0:
		return gb.ret(!gb.regs.getC())
	case 0xE0:
		return gb.storeAI8()
	case 0xF0:
		return gb.loadAI8()
	case 0xC1:
		return gb.pop(&gb.regs.BC)
	case 0xD1:
		return gb.pop(&gb.regs.DE)
	case 0xE1:
		return gb.pop(&gb.regs.HL)
	case 0xF1:
		return gb.pop(&gb.regs.AF)
	case 0xC2:
		return gb.JumpI16(!gb.regs.getZ())
	case 0xD2:
		return gb.JumpI16(!gb.regs.getC())
	case 0xE2:
		return gb.storeAC()
	case 0xF2:
		return gb.loadAC()
	case 0xC3:
		return gb.JumpI16(true)
	case 0xF3:
		return disableInterrupts()
	case 0xC4:
		return gb.call(!gb.regs.getZ())
	case 0xD4:
		return gb.call(!gb.regs.getC())
	case 0xCB:
		return gb.execCBInstr()
	}

	// TODO throw exception
	return -1
}

func (gb *gameboy) getImmediate() uint8 {
	return 0
}

func (gb *gameboy) clockCycle(mCycles int) {
}

func nop() int {
	// TODO

	return 1
}

func stop() int {
	// TODO

	return 1
}

func halt() int {
	// TODO

	return 1
}

func disableInterrupts() int {
	// TODO

	return 1
}

// JumpI16 conditional jump
func (gb *gameboy) JumpI16(flag bool) int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if flag {
		gb.regs.PC = getWordFromBytes(hi, lo)
		return 4
	}
	return 3
}

// JumpRelativeI8 relative conditional jump
func (gb *gameboy) JumpRelativeI8(flag bool) int {
	im8 := gb.getImmediate()
	if flag {
		gb.regs.PC += uint16(im8)
		return 3
	}
	return 2
}

func (gb *gameboy) ret(cond bool) int {
	if cond {
		P := gb.mem.load(gb.regs.SP)
		gb.regs.SP++
		S := gb.mem.load(gb.regs.SP)
		gb.regs.SP++
		gb.regs.PC = getWordFromBytes(S, P)
		return 5
	}
	return 2
}

func (gb *gameboy) call(cond bool) int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if cond {
		p, c := getBytesFromWord(gb.regs.PC)
		gb.mem.store(p, gb.regs.SP)
		gb.regs.SP--
		gb.mem.store(c, gb.regs.SP)
		gb.regs.SP--
		gb.regs.PC = getWordFromBytes(hi, lo)
		return 6
	}
	return 3
}

// loadI8 load an 8-bit immediate into a register
func (gb *gameboy) loadI8(reg *uint8) int {
	*reg = gb.getImmediate()
	return 2
}

// loadR8R8 copy r2 into r1
func loadR8R8(r1, r2 *uint8) int {
	*r1 = *r2
	return 1
}

// loadI16 load a 16-bit immediate into a register
func (gb *gameboy) loadI16(reg *[2]uint8) int {
	reg[0] = gb.getImmediate()
	reg[1] = gb.getImmediate()
	return 3
}

func (gb *gameboy) loadI16SP() int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.regs.SP = getWordFromBytes(hi, lo)
	return 3
}

// loadR8 load an 8-bit val from memory into the given register
func (gb *gameboy) loadR8(reg *uint8, adr uint16) int {
	*reg = gb.mem.load(adr)
	return 2
}

// loadAI8 load the content of an address in the block 0xFF00 - 0xFFFF given by an i8 into register A
func (gb *gameboy) loadAI8() int {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.regs.AF[1] = gb.mem.load(adr)
	return 3
}

// loadAC load the content of an address in the block 0xFF00 - 0xFFFF given by register C into register A
func (gb *gameboy) loadAC() int {
	adr := uint16(0xFF00) | uint16(gb.regs.BC[0])
	gb.regs.AF[1] = gb.mem.load(adr)
	return 2
}

// storeR8 store the content of register at regVal in the address specified by RegAdr.
func (gb *gameboy) storeR8(val uint8, adr uint16) int {
	gb.mem.store(val, adr)
	return 2
}

// storeSPI16 store the stack pointer in the memory address provided by the 16-bit immediate
func (gb *gameboy) storeSPI16() int {
	adr := uint16(gb.getImmediate())
	adr += uint16(gb.getImmediate()) << 8
	S, P := getBytesFromWord(gb.regs.SP)
	gb.mem.store(P, adr)
	gb.mem.store(S, adr+1)
	return 5
}

// storeI8 store the immediate 8-bit value into the memory address specified by the 16-bit register
func (gb *gameboy) storeI8(reg *[2]uint8) int {
	gb.mem.store(gb.getImmediate(), getWord(reg))
	return 3
}

// storeAI8 store the content of register A in an address in the block 0xFF00 - 0xFFFF given by an i8
func (gb *gameboy) storeAI8() int {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.mem.store(gb.regs.AF[1], adr)
	return 3
}

// storeAC store the content of register A in an address in the block 0xFF00 - 0xFFFF given by C
func (gb *gameboy) storeAC() int {
	adr := uint16(0xFF00) | uint16(gb.regs.BC[0])
	gb.mem.store(gb.regs.AF[1], adr)
	return 2
}

func (gb *gameboy) push(reg *[2]uint8) int {
	gb.mem.store(reg[1], gb.regs.SP)
	gb.regs.SP++
	gb.mem.store(reg[0], gb.regs.SP)
	gb.regs.SP++
	return 4
}

// pop load a 16bit value from memory and increment the stack pointer during the load (twice in total)
func (gb *gameboy) pop(reg *[2]uint8) int {
	reg[0] = gb.mem.load(gb.regs.SP)
	gb.regs.SP++
	reg[1] = gb.mem.load(gb.regs.SP)
	gb.regs.SP++
	return 3
}

// incR16 increments a combine 16-bit register.
func incR16(reg *[2]uint8) int {
	setWord(reg, getWord(reg)+1)
	return 2
}

// incSP increments a combine 16-bit register.
func (regs *registers) incSP() int {
	regs.SP++
	return 2
}

// incR8 increment the given 8-bit register
func (regs *registers) incR8(reg *uint8) int {
	regs.setN(false)
	regs.setH(halfCarryAddCheck8Bit(*reg, 1))
	*reg++
	regs.setZ(*reg == 0)
	return 1
}

// incM8 increment the 8 bit value at the specified memory address
func (gb *gameboy) incM8(adr uint16) int {
	val := gb.mem.load(adr)
	gb.regs.setN(false)
	gb.regs.setH(halfCarryAddCheck8Bit(val, 1))
	val++
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	return 3
}

// decR16 increments a combine 16-bit register.
func decR16(reg *[2]uint8) int {
	setWord(reg, getWord(reg)-1)
	return 2
}

// decSP increments a combine 16-bit register.
func (regs *registers) decSP() int {
	regs.SP--
	return 2
}

// decR8 decrement the given 8-bit register
func (regs *registers) decR8(reg *uint8) int {
	regs.setN(true)
	regs.setH(halfCarrySubCheck8Bit(*reg, 1))
	*reg--
	regs.setZ(*reg == 0)
	return 1
}

// decM8 decrement the 8 bit value at the specified memory address
func (gb *gameboy) decM8(adr uint16) int {
	val := gb.mem.load(adr)
	gb.regs.setN(true)
	gb.regs.setH(halfCarrySubCheck8Bit(val, 1))
	val--
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	return 3
}

// addR16R16 add the contents of one 16-bit register pair to another
func (regs *registers) addR16R16(reg *[2]uint8, val uint16) int {
	a := getWord(reg)

	regs.setH(halfCarryAddCheck16Bit(a, val))
	regs.setC(a+val < a)
	regs.setN(false)
	setWord(reg, a+val)
	return 2
}

// addR8 add the 8-bit value of a register to A
func (regs *registers) addR8(val uint8) int {
	a := regs.AF[1]
	regs.setFlags(a+val == 0, false, halfCarryAddCheck8Bit(a, val), a+val < a)
	regs.AF[1] += val
	return 1
}

// adcR8 add the 8-bit value of a register to A
func (regs *registers) adcR8(val uint8) int {
	if regs.getC() {
		return regs.addR8(val + 1)
	}
	return regs.addR8(val)
}

// subR8 subtract the 8-bit value of a register from A
func (regs *registers) subR8(val uint8) int {
	a := regs.AF[1]
	regs.setFlags(a+val == 0, false, halfCarrySubCheck8Bit(a, val), a-val > a)
	regs.AF[1] -= val
	return 1
}

// sbcR8 subtract the 8-bit value of a register from A
func (regs *registers) sbcR8(val uint8) int {
	if regs.getC() {
		return regs.subR8(val + 1)
	}
	return regs.subR8(val)
}

// andR8 logical AND the 8-bit value of a register with A
func (regs *registers) andR8(val uint8) int {
	regs.AF[1] &= val
	regs.setFlags(regs.AF[1] == 0, false, true, false)
	return 1
}

// orR8 logical OR the 8-bit value of a register with A
func (regs *registers) orR8(val uint8) int {
	regs.AF[1] |= val
	regs.setFlags(regs.AF[1] == 0, false, false, false)
	return 1
}

// xorR8 logical XOR the 8-bit value of a register with A
func (regs *registers) xorR8(val uint8) int {
	regs.AF[1] ^= val
	regs.setFlags(regs.AF[1] == 0, false, false, false)
	return 1
}

// cpR8 compare the 8-bit value of a register with A
func (regs *registers) cpR8(val uint8) int {
	a := regs.AF[1]
	regs.setFlags(a+val == 0, false, halfCarrySubCheck8Bit(a, val), a-val > a)
	return 1
}

// aluR8Def function definition of an 8-bit alu function
type aluR8Def func(uint8) int

// aluM8 executes an 8-bit alu function with the value from the given memory address
func (gb *gameboy) aluM8(adr uint16, aluFunc aluR8Def) int {
	aluFunc(gb.mem.load(adr))
	return 2
}

// complementR8 bit-swap the register
func complementR8(r *uint8) int {
	*r = ^*r
	return 1
}

// rotateLeftCircularA circular rotate register A left
func (regs *registers) rotateLeftCircularA() int {
	regs.AF[1] = regs.rotateLeftCircularInternal(regs.AF[1])
	return 1
}

// rotateLeftA rotate register A left
func (regs *registers) rotateLeftA() int {
	regs.AF[1] = regs.rotateLeftInternal(regs.AF[1])
	return 1
}

// rotateRightCircularA circular rotate register A left
func (regs *registers) rotateRightCircularA() int {
	regs.AF[1] = regs.rotateRightCircularInternal(regs.AF[1])
	return 1
}

// rotateRightA rotate register A left
func (regs *registers) rotateRightA() int {
	regs.AF[1] = regs.rotateRightInternal(regs.AF[1])
	return 1
}

// setCarryFlag sets the carry flag and unsets N and C
func (regs *registers) setCarryFlag(val bool) int {
	regs.setN(false)
	regs.setH(false)
	regs.setC(val)
	return 1
}

/*
decimalAdjustA decimal-adjusts the number

this is nuts
*/
func (regs *registers) decimalAdjustA() int {
	if !regs.getN() {
		if regs.getC() || regs.AF[1] > 0x99 {
			regs.AF[1] += 0x000_0060
			regs.setC(true)
		}
		if regs.getH() || (regs.AF[1]&0x0f) > 0x09 {
			regs.AF[1] += 0x000_0006
		}
	} else {
		if regs.getC() {
			regs.AF[1] -= 0x000_0060
		}
		if regs.getH() {
			regs.AF[1] += 0x000_0006
		}
	}

	regs.setZ(regs.AF[1] == 0)
	regs.setH(false)
	return 1
}
