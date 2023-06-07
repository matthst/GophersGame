package gameboy

import "fmt"

func (gb *Gameboy) execNextInstr() int {

	opcode := gb.getImmediate() //fetch opcode

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */
	switch opcode {
	case 0x00:
		return nop()
	case 0x10:
		return stop()
	case 0x20:
		return gb.JumpRelativeI8(!gb.getZFlag())
	case 0x30:
		return gb.JumpRelativeI8(!gb.getNFlag())
	case 0x01:
		return gb.loadI16(&gb.B, &gb.C)
	case 0x11:
		return gb.loadI16(&gb.D, &gb.E)
	case 0x21:
		return gb.loadI16(&gb.H, &gb.L)
	case 0x31:
		return gb.loadI16SP()
	case 0x02:
		return gb.storeR8(gb.A, gb.getBC())
	case 0x12:
		return gb.storeR8(gb.A, gb.getDE())
	case 0x22:
		return gb.storeR8(gb.A, getWordInc(&gb.H, &gb.L))
	case 0x32:
		return gb.storeR8(gb.A, getWordDec(&gb.H, &gb.L))
	case 0x03:
		return incR16(&gb.B, &gb.C)
	case 0x13:
		return incR16(&gb.D, &gb.E)
	case 0x23:
		return incR16(&gb.H, &gb.L)
	case 0x33:
		return gb.incSP()
	case 0x04:
		return gb.incR8(&gb.B)
	case 0x14:
		return gb.incR8(&gb.D)
	case 0x24:
		return gb.incR8(&gb.H)
	case 0x34:
		return gb.incM8(gb.getHL())
	case 0x05:
		return gb.decR8(&gb.B)
	case 0x15:
		return gb.decR8(&gb.D)
	case 0x25:
		return gb.decR8(&gb.H)
	case 0x35:
		return gb.decM8(gb.getHL())
	case 0x06:
		return gb.loadI8(&gb.B)
	case 0x16:
		return gb.loadI8(&gb.D)
	case 0x26:
		return gb.loadI8(&gb.H)
	case 0x36:
		return gb.storeI8()
	case 0x07:
		return gb.rotateLeftCircularA()
	case 0x17:
		return gb.rotateLeftA()
	case 0x27:
		return gb.decimalAdjustA()
	case 0x37:
		return gb.setCarryFlag(true)
	case 0x08:
		return gb.storeSPI16()
	case 0x18:
		return gb.JumpRelativeI8(true)
	case 0x28:
		return gb.JumpRelativeI8(gb.getZFlag())
	case 0x38:
		return gb.JumpRelativeI8(gb.getCFlag())
	case 0x09:
		return gb.addR16R16(gb.getBC())
	case 0x19:
		return gb.addR16R16(gb.getDE())
	case 0x29:
		return gb.addR16R16(gb.getHL())
	case 0x39:
		return gb.addR16R16(gb.SP)
	case 0x0A:
		return gb.loadR8(&gb.A, gb.getBC())
	case 0x1A:
		return gb.loadR8(&gb.A, gb.getDE())
	case 0x2A:
		return gb.loadR8(&gb.A, getWordInc(&gb.H, &gb.L))
	case 0x3A:
		return gb.loadR8(&gb.A, getWordDec(&gb.H, &gb.L))
	case 0x0B:
		return decR16(&gb.B, &gb.C)
	case 0x1B:
		return decR16(&gb.D, &gb.E)
	case 0x2B:
		return decR16(&gb.H, &gb.L)
	case 0x3B:
		return gb.decSP()
	case 0x0C:
		return gb.incR8(&gb.C)
	case 0x1C:
		return gb.incR8(&gb.E)
	case 0x2C:
		return gb.incR8(&gb.L)
	case 0x3C:
		return gb.incR8(&gb.A)
	case 0x0D:
		return gb.decR8(&gb.C)
	case 0x1D:
		return gb.decR8(&gb.E)
	case 0x2D:
		return gb.decR8(&gb.L)
	case 0x3D:
		return gb.decR8(&gb.A)
	case 0x0E:
		return gb.loadI8(&gb.C)
	case 0x1E:
		return gb.loadI8(&gb.E)
	case 0x2E:
		return gb.loadI8(&gb.L)
	case 0x3E:
		return gb.loadI8(&gb.A)
	case 0x0F:
		return gb.rotateRightCircularA()
	case 0x1F:
		return gb.rotateRightA()
	case 0x2F:
		return complementR8(&gb.A)
	case 0x3F:
		return gb.setCarryFlag(!gb.getCFlag())
	case 0x40:
		return loadR8R8(&gb.B, &gb.B)
	case 0x50:
		return loadR8R8(&gb.D, &gb.B)
	case 0x60:
		return loadR8R8(&gb.H, &gb.B)
	case 0x70:
		return gb.storeR8(gb.B, gb.getHL())
	case 0x41:
		return loadR8R8(&gb.B, &gb.C)
	case 0x51:
		return loadR8R8(&gb.D, &gb.C)
	case 0x61:
		return loadR8R8(&gb.H, &gb.C)
	case 0x71:
		return gb.storeR8(gb.C, gb.getHL())
	case 0x42:
		return loadR8R8(&gb.B, &gb.D)
	case 0x52:
		return loadR8R8(&gb.D, &gb.D)
	case 0x62:
		return loadR8R8(&gb.H, &gb.D)
	case 0x72:
		return gb.storeR8(gb.D, gb.getHL())
	case 0x43:
		return loadR8R8(&gb.B, &gb.E)
	case 0x53:
		return loadR8R8(&gb.D, &gb.E)
	case 0x63:
		return loadR8R8(&gb.H, &gb.E)
	case 0x73:
		return gb.storeR8(gb.E, gb.getHL())
	case 0x44:
		return loadR8R8(&gb.B, &gb.H)
	case 0x54:
		return loadR8R8(&gb.D, &gb.H)
	case 0x64:
		return loadR8R8(&gb.H, &gb.H)
	case 0x74:
		return gb.storeR8(gb.H, gb.getHL())
	case 0x45:
		return loadR8R8(&gb.B, &gb.L)
	case 0x55:
		return loadR8R8(&gb.D, &gb.L)
	case 0x65:
		return loadR8R8(&gb.H, &gb.L)
	case 0x75:
		return gb.storeR8(gb.L, gb.getHL())
	case 0x46:
		return gb.loadR8(&gb.B, gb.getHL())
	case 0x56:
		return gb.loadR8(&gb.D, gb.getHL())
	case 0x66:
		return gb.loadR8(&gb.H, gb.getHL())
	case 0x76:
		return gb.halt()
	case 0x47:
		return loadR8R8(&gb.B, &gb.A)
	case 0x57:
		return loadR8R8(&gb.D, &gb.A)
	case 0x67:
		return loadR8R8(&gb.H, &gb.A)
	case 0x77:
		return gb.storeR8(gb.A, gb.getHL())
	case 0x48:
		return loadR8R8(&gb.C, &gb.B)
	case 0x58:
		return loadR8R8(&gb.E, &gb.B)
	case 0x68:
		return loadR8R8(&gb.L, &gb.B)
	case 0x78:
		return loadR8R8(&gb.A, &gb.B)
	case 0x49:
		return loadR8R8(&gb.C, &gb.C)
	case 0x59:
		return loadR8R8(&gb.E, &gb.C)
	case 0x69:
		return loadR8R8(&gb.L, &gb.C)
	case 0x79:
		return loadR8R8(&gb.A, &gb.C)
	case 0x4A:
		return loadR8R8(&gb.C, &gb.D)
	case 0x5A:
		return loadR8R8(&gb.E, &gb.D)
	case 0x6A:
		return loadR8R8(&gb.L, &gb.D)
	case 0x7A:
		return loadR8R8(&gb.A, &gb.D)
	case 0x4B:
		return loadR8R8(&gb.C, &gb.E)
	case 0x5B:
		return loadR8R8(&gb.E, &gb.E)
	case 0x6B:
		return loadR8R8(&gb.L, &gb.E)
	case 0x7B:
		return loadR8R8(&gb.A, &gb.E)
	case 0x4C:
		return loadR8R8(&gb.C, &gb.H)
	case 0x5C:
		return loadR8R8(&gb.E, &gb.H)
	case 0x6C:
		return loadR8R8(&gb.L, &gb.H)
	case 0x7C:
		return loadR8R8(&gb.A, &gb.H)
	case 0x4D:
		return loadR8R8(&gb.C, &gb.L)
	case 0x5D:
		return loadR8R8(&gb.E, &gb.L)
	case 0x6D:
		return loadR8R8(&gb.L, &gb.L)
	case 0x7D:
		return loadR8R8(&gb.A, &gb.L)
	case 0x4E:
		return gb.loadR8(&gb.C, gb.getHL())
	case 0x5E:
		return gb.loadR8(&gb.E, gb.getHL())
	case 0x6E:
		return gb.loadR8(&gb.L, gb.getHL())
	case 0x7E:
		return gb.loadR8(&gb.A, gb.getHL())
	case 0x4F:
		return loadR8R8(&gb.C, &gb.A)
	case 0x5F:
		return loadR8R8(&gb.E, &gb.A)
	case 0x6F:
		return loadR8R8(&gb.L, &gb.A)
	case 0x7F:
		return loadR8R8(&gb.A, &gb.A)
	case 0x80:
		return gb.addR8(gb.B)
	case 0x90:
		return gb.subR8(gb.B)
	case 0xA0:
		return gb.andR8(gb.B)
	case 0xB0:
		return gb.orR8(gb.B)
	case 0x81:
		return gb.addR8(gb.C)
	case 0x91:
		return gb.subR8(gb.C)
	case 0xA1:
		return gb.andR8(gb.C)
	case 0xB1:
		return gb.orR8(gb.C)
	case 0x82:
		return gb.addR8(gb.D)
	case 0x92:
		return gb.subR8(gb.D)
	case 0xA2:
		return gb.andR8(gb.D)
	case 0xB2:
		return gb.orR8(gb.D)
	case 0x83:
		return gb.addR8(gb.E)
	case 0x93:
		return gb.subR8(gb.E)
	case 0xA3:
		return gb.andR8(gb.E)
	case 0xB3:
		return gb.orR8(gb.E)
	case 0x84:
		return gb.addR8(gb.H)
	case 0x94:
		return gb.subR8(gb.H)
	case 0xA4:
		return gb.andR8(gb.H)
	case 0xB4:
		return gb.orR8(gb.H)
	case 0x85:
		return gb.addR8(gb.L)
	case 0x95:
		return gb.subR8(gb.L)
	case 0xA5:
		return gb.andR8(gb.L)
	case 0xB5:
		return gb.orR8(gb.L)
	case 0x86:
		return gb.aluM8(gb.addR8)
	case 0x96:
		return gb.aluM8(gb.subR8)
	case 0xA6:
		return gb.aluM8(gb.andR8)
	case 0xB6:
		return gb.aluM8(gb.orR8)
	case 0x87:
		return gb.addR8(gb.A)
	case 0x97:
		return gb.subR8(gb.A)
	case 0xA7:
		return gb.andR8(gb.A)
	case 0xB7:
		return gb.orR8(gb.A)
	case 0x88:
		return gb.adcR8(gb.B)
	case 0x98:
		return gb.sbcR8(gb.B)
	case 0xA8:
		return gb.xorR8(gb.B)
	case 0xB8:
		return gb.cpR8(gb.B)
	case 0x89:
		return gb.adcR8(gb.C)
	case 0x99:
		return gb.sbcR8(gb.C)
	case 0xA9:
		return gb.xorR8(gb.C)
	case 0xB9:
		return gb.cpR8(gb.C)
	case 0x8A:
		return gb.adcR8(gb.D)
	case 0x9A:
		return gb.sbcR8(gb.D)
	case 0xAA:
		return gb.xorR8(gb.D)
	case 0xBA:
		return gb.cpR8(gb.D)
	case 0x8B:
		return gb.adcR8(gb.E)
	case 0x9B:
		return gb.sbcR8(gb.E)
	case 0xAB:
		return gb.xorR8(gb.E)
	case 0xBB:
		return gb.cpR8(gb.E)
	case 0x8C:
		return gb.adcR8(gb.H)
	case 0x9C:
		return gb.sbcR8(gb.H)
	case 0xAC:
		return gb.xorR8(gb.H)
	case 0xBC:
		return gb.cpR8(gb.H)
	case 0x8D:
		return gb.adcR8(gb.L)
	case 0x9D:
		return gb.sbcR8(gb.L)
	case 0xAD:
		return gb.xorR8(gb.L)
	case 0xBD:
		return gb.cpR8(gb.L)
	case 0x8E:
		return gb.aluM8(gb.adcR8)
	case 0x9E:
		return gb.aluM8(gb.sbcR8)
	case 0xAE:
		return gb.aluM8(gb.xorR8)
	case 0xBE:
		return gb.aluM8(gb.cpR8)
	case 0x8F:
		return gb.adcR8(gb.A)
	case 0x9F:
		return gb.sbcR8(gb.A)
	case 0xAF:
		return gb.xorR8(gb.A)
	case 0xBF:
		return gb.cpR8(gb.A)
	case 0xC0:
		return gb.retCond(!gb.getZFlag())
	case 0xD0:
		return gb.retCond(!gb.getCFlag())
	case 0xE0:
		return gb.storeAI8()
	case 0xF0:
		return gb.loadAI8()
	case 0xC1:
		return gb.pop(&gb.B, &gb.C)
	case 0xD1:
		return gb.pop(&gb.D, &gb.E)
	case 0xE1:
		return gb.pop(&gb.H, &gb.L)
	case 0xF1:
		return gb.pop(&gb.A, &gb.F)
	case 0xC2:
		return gb.JumpI16(!gb.getZFlag())
	case 0xD2:
		return gb.JumpI16(!gb.getCFlag())
	case 0xE2:
		return gb.storeAC()
	case 0xF2:
		return gb.loadAC()
	case 0xC3:
		return gb.JumpI16(true)
	case 0xF3:
		return gb.disableInterrupts()
	case 0xC4:
		return gb.call(!gb.getZFlag())
	case 0xD4:
		return gb.call(!gb.getCFlag())
	case 0xC6:
		return gb.aluI8(gb.addR8)
	case 0xD6:
		return gb.aluI8(gb.subR8)
	case 0xE6:
		return gb.aluI8(gb.andR8)
	case 0xF6:
		return gb.aluI8(gb.orR8)
	case 0xC5:
		return gb.push(gb.B, gb.C)
	case 0xD5:
		return gb.push(gb.D, gb.E)
	case 0xE5:
		return gb.push(gb.H, gb.L)
	case 0xF5:
		return gb.push(gb.A, gb.F)
	case 0xC7:
		return gb.rst(0x00)
	case 0xD7:
		return gb.rst(0x10)
	case 0xE7:
		return gb.rst(0x20)
	case 0xF7:
		return gb.rst(0x30)
	case 0xC8:
		return gb.retCond(gb.getZFlag())
	case 0xD8:
		return gb.retCond(gb.getCFlag())
	case 0xE8:
		return gb.addSPS8SP()
	case 0xF8:
		return gb.addSPS8HL()
	case 0xC9:
		return gb.ret()
	case 0xD9:
		return gb.retInterrupt()
	case 0xE9:
		return gb.jumpHL()
	case 0xF9:
		return gb.loadHLSP()
	case 0xCA:
		return gb.JumpI16(gb.getZFlag())
	case 0xDA:
		return gb.JumpI16(gb.getCFlag())
	case 0xEA:
		return gb.storeAMI16()
	case 0xFA:
		return gb.loadAMI16()
	case 0xCB:
		return gb.execCBInstr()
	case 0xFB:
		return gb.enableInterrupts()
	case 0xCC:
		return gb.call(gb.getZFlag())
	case 0xDC:
		return gb.call(gb.getCFlag())
	case 0xCD:
		return gb.call(true)
	case 0xCE:
		return gb.aluM8(gb.adcR8)
	case 0xDE:
		return gb.aluM8(gb.sbcR8)
	case 0xEE:
		return gb.aluM8(gb.xorR8)
	case 0xFE:
		return gb.aluM8(gb.cpR8)
	case 0xCF:
		return gb.rst(0x08)
	case 0xDF:
		return gb.rst(0x18)
	case 0xEF:
		return gb.rst(0x28)
	case 0xFF:
		return gb.rst(0x38)
	}
	panic(fmt.Sprintf("Opcode '%X' is not a valid opcode", opcode))
}

func nop() int {
	return 1
}

func stop() int {
	// TODO implement the stop instruction
	panic("Stop instruction (0x10) not implemented")
	return 1
}

func (gb *Gameboy) halt() int {
	if gb.IE&gb.IF&0x1F != 0 {
		gb.haltBug = true
	} else {
		gb.haltMode = true
	}
	return 1
}

func (gb *Gameboy) disableInterrupts() int {
	gb.IME = false
	return 1
}

func (gb *Gameboy) enableInterrupts() int {
	if gb.EICounter == 0 {
		gb.EICounter = 2
	}
	return 1
}

// JumpI16 conditional jump
func (gb *Gameboy) JumpI16(flag bool) int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if flag {
		gb.PC = getWord(hi, lo)
		return 4
	}
	return 3
}

// JumpRelativeI8 relative conditional jump
func (gb *Gameboy) JumpRelativeI8(flag bool) int {
	im8 := gb.getImmediate()
	if flag {
		gb.PC += uint16(im8)
		return 3
	}
	return 2
}

func (gb *Gameboy) jumpHL() int {
	gb.SP = gb.getHL()
	return 1
}

func (gb *Gameboy) retCond(cond bool) int {
	if cond {
		P := gb.load(gb.SP)
		gb.SP++
		S := gb.load(gb.SP)
		gb.SP++
		gb.PC = getWord(S, P)
		return 5
	}
	return 2
}

func (gb *Gameboy) ret() int {
	P := gb.load(gb.SP)
	gb.SP++
	S := gb.load(gb.SP)
	gb.SP++
	gb.PC = getWord(S, P)
	return 4
}

func (gb *Gameboy) retInterrupt() int {
	gb.IME = true
	return gb.ret()
}

func (gb *Gameboy) call(cond bool) int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if cond {
		gb.rst(getWord(hi, lo))
		return 6
	}
	return 3
}

func (gb *Gameboy) rst(adr uint16) int {
	P, C := getBytes(gb.PC)
	gb.write(P, gb.SP)
	gb.SP--
	gb.write(C, gb.SP)
	gb.SP--
	gb.PC = adr
	return 4
}

// loadI8 load an 8-bit immediate into a register
func (gb *Gameboy) loadI8(reg *uint8) int {
	*reg = gb.getImmediate()
	return 2
}

// loadR8R8 copy r2 into r1
func loadR8R8(r1, r2 *uint8) int {
	*r1 = *r2
	return 1
}

// loadI16 load a 16-bit immediate into a register
func (gb *Gameboy) loadI16(hi, lo *uint8) int {
	*lo = gb.getImmediate()
	*hi = gb.getImmediate()
	return 3
}

// loadHLSP load the value of SP into HL
func (gb *Gameboy) loadHLSP() int {
	setBytes(&gb.H, &gb.L, gb.SP)
	return 2
}

func (gb *Gameboy) loadI16SP() int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.SP = getWord(hi, lo)
	return 3
}

// loadR8 load an 8-bit val from memory into the given register
func (gb *Gameboy) loadR8(reg *uint8, adr uint16) int {
	*reg = gb.load(adr)
	return 2
}

// loadMAI16 load an 8-bit val into A from the memory address specified by the 16-bit immediate
func (gb *Gameboy) loadAMI16() int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.A = gb.load(getWord(hi, lo))
	return 4
}

// loadAI8 load the content of an address in the block 0xFF00 - 0xFFFF given by an i8 into register A
func (gb *Gameboy) loadAI8() int {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.A = gb.load(adr)
	return 3
}

// loadAC load the content of an address in the block 0xFF00 - 0xFFFF given by register C into register A
func (gb *Gameboy) loadAC() int {
	adr := uint16(0xFF00) | uint16(gb.C)
	gb.A = gb.load(adr)
	return 2
}

// storeR8 store the content of register at regVal in the address specified by RegAdr.
func (gb *Gameboy) storeR8(val uint8, adr uint16) int {
	gb.write(val, adr)
	return 2
}

// storeSPI16 store the stack pointer in the memory address provided by the 16-bit immediate
func (gb *Gameboy) storeSPI16() int {
	adr := uint16(gb.getImmediate())
	adr += uint16(gb.getImmediate()) << 8
	S, P := getBytes(gb.SP)
	gb.write(P, adr)
	gb.write(S, adr+1)
	return 5
}

// storeI8 store the immediate 8-bit value into the memory address specified by HL
func (gb *Gameboy) storeI8() int {
	gb.write(gb.getImmediate(), gb.getHL())
	return 3
}

// storeAI8 store the content of register A in an address in the block 0xFF00 - 0xFFFF given by an i8
func (gb *Gameboy) storeAI8() int {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.write(gb.A, adr)
	return 3
}

// storeAMI16 store an 8-bit val from A into the memory address specified by the 16-bit immediate
func (gb *Gameboy) storeAMI16() int {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.write(gb.A, getWord(hi, lo))
	return 4
}

// storeAC store the content of register A in an address in the block 0xFF00 - 0xFFFF given by C
func (gb *Gameboy) storeAC() int {
	gb.write(gb.A, uint16(0xFF00)|uint16(gb.C))
	return 2
}

func (gb *Gameboy) push(hi, lo uint8) int {
	gb.write(hi, gb.SP)
	gb.SP--
	gb.write(lo, gb.SP)
	gb.SP--
	return 4
}

// pop load a 16bit value from memory and increment the stack pointer during the load (twice in total)
func (gb *Gameboy) pop(hi, lo *uint8) int {
	*lo = gb.load(gb.SP)
	gb.SP++
	*hi = gb.load(gb.SP)
	gb.SP++
	return 3
}

// incR16 increments a combine 16-bit register.
func incR16(hi, lo *uint8) int {
	setBytes(hi, lo, getWord(*hi, *lo)+1)
	return 2
}

// incSP increments a combine 16-bit register.
func (gb *Gameboy) incSP() int {
	gb.SP++
	return 2
}

// incR8 increment the given 8-bit register
func (gb *Gameboy) incR8(reg *uint8) int {
	gb.setNFlag(false)
	gb.setHFlag(halfCarryAddCheck8Bit(*reg, 1))
	*reg++
	gb.setZFlag(*reg == 0)
	return 1
}

// incM8 increment the 8 bit value at the specified memory address
func (gb *Gameboy) incM8(adr uint16) int {
	val := gb.load(adr)
	gb.setNFlag(false)
	gb.setHFlag(halfCarryAddCheck8Bit(val, 1))
	val++
	gb.setZFlag(val == 0)
	gb.write(val, adr)
	return 3
}

// decR16 increments a combine 16-bit register.
func decR16(hi, lo *uint8) int {
	setBytes(hi, lo, getWord(*hi, *lo)-1)
	return 2
}

// decSP increments a combine 16-bit register.
func (gb *Gameboy) decSP() int {
	gb.SP--
	return 2
}

// decR8 decrement the given 8-bit register
func (gb *Gameboy) decR8(reg *uint8) int {
	gb.setNFlag(true)
	gb.setHFlag(halfCarrySubCheck8Bit(*reg, 1))
	*reg--
	gb.setZFlag(*reg == 0)
	return 1
}

// decM8 decrement the 8 bit value at the specified memory address
func (gb *Gameboy) decM8(adr uint16) int {
	val := gb.load(adr)
	gb.setNFlag(true)
	gb.setHFlag(halfCarrySubCheck8Bit(val, 1))
	val--
	gb.setZFlag(val == 0)
	gb.write(val, adr)
	return 3
}

// addR16R16 add the contents of one 16-bit register pair to the register HL
func (gb *Gameboy) addR16R16(val uint16) int {
	HL := gb.getHL()
	gb.setHFlag(halfCarryAddCheck16Bit(HL, val))
	gb.setCFlag(HL+val < HL)
	gb.setNFlag(false)
	setBytes(&gb.H, &gb.L, HL+val)
	return 2
}

// addSP8SP add the signed 2's complement immediate to the stack pointer and write it to HL
func (gb *Gameboy) addSPS8SP() int {
	gb.SP = gb.addSPS8Internal()
	return 4
}

// addSPS8HL add the signed 2's complement immediate to the stack pointer and write it to HL
func (gb *Gameboy) addSPS8HL() int {
	setBytes(&gb.H, &gb.L, gb.addSPS8Internal())
	return 3
}

// addSPS8 add the signed 2's complement immediate to the stack pointer and return the value
func (gb *Gameboy) addSPS8Internal() uint16 {
	val := gb.getImmediate()
	P := uint8(gb.SP)
	gb.setZFlag(false)
	gb.setNFlag(false)
	if val < 128 { // positive 2's complement value :=
		gb.setHFlag(halfCarryAddCheck8Bit(P, val))
		gb.setCFlag(P+val < P)
		return gb.SP + uint16(val)
	}
	// negative 2's complement value
	val = ^val + 1 //get positive value from 2's complement signed number
	gb.setHFlag(halfCarrySubCheck8Bit(P, val))
	gb.setCFlag(P-val > P)
	return gb.SP - uint16(val)
}

// addR8 add the 8-bit value of a register to A
func (gb *Gameboy) addR8(val uint8) int {
	a := gb.A
	gb.setFlags(a+val == 0, false, halfCarryAddCheck8Bit(a, val), a+val < a)
	gb.A += val
	return 1
}

// adcR8 add the 8-bit value of a register to A
func (gb *Gameboy) adcR8(val uint8) int {
	if gb.getCFlag() {
		return gb.addR8(val + 1)
	}
	return gb.addR8(val)
}

// subR8 subtract the 8-bit value of a register from A
func (gb *Gameboy) subR8(val uint8) int {
	a := gb.A
	gb.setFlags(a+val == 0, false, halfCarrySubCheck8Bit(a, val), a-val > a)
	gb.A -= val
	return 1
}

// sbcR8 subtract the 8-bit value of a register from A
func (gb *Gameboy) sbcR8(val uint8) int {
	if gb.getCFlag() {
		return gb.subR8(val + 1)
	}
	return gb.subR8(val)
}

// andR8 logical AND the 8-bit value of a register with A
func (gb *Gameboy) andR8(val uint8) int {
	gb.A &= val
	gb.setFlags(gb.A == 0, false, true, false)
	return 1
}

// orR8 logical OR the 8-bit value of a register with A
func (gb *Gameboy) orR8(val uint8) int {
	gb.A |= val
	gb.setFlags(gb.A == 0, false, false, false)
	return 1
}

// xorR8 logical XOR the 8-bit value of a register with A
func (gb *Gameboy) xorR8(val uint8) int {
	gb.A ^= val
	gb.setFlags(gb.A == 0, false, false, false)
	return 1
}

// cpR8 compare the 8-bit value of a register with A
func (gb *Gameboy) cpR8(val uint8) int {
	gb.setFlags(gb.A+val == 0, false, halfCarrySubCheck8Bit(gb.A, val), gb.A-val > gb.A)
	return 1
}

// aluR8Def function definition of an 8-bit alu function
type aluR8Def func(uint8) int

// aluI8 executes an 8-bit alu function with the given immediate
func (gb *Gameboy) aluI8(aluFunc aluR8Def) int {
	aluFunc(gb.getImmediate())
	return 2
}

// aluM8 executes an 8-bit alu function with the value from the given memory address
func (gb *Gameboy) aluM8(aluFunc aluR8Def) int {
	aluFunc(gb.load(gb.getHL()))
	return 2
}

// complementR8 bit-swap the register
func complementR8(r *uint8) int {
	*r = ^*r
	return 1
}

// rotateLeftCircularA circular rotate register A left
func (gb *Gameboy) rotateLeftCircularA() int {
	gb.A = gb.rotateLeftCircularInternal(gb.A)
	return 1
}

// rotateLeftA rotate register A left
func (gb *Gameboy) rotateLeftA() int {
	gb.A = gb.rotateLeftInternal(gb.A)
	return 1
}

// rotateRightCircularA circular rotate register A left
func (gb *Gameboy) rotateRightCircularA() int {
	gb.A = gb.rotateRightCircularInternal(gb.A)
	return 1
}

// rotateRightA rotate register A left
func (gb *Gameboy) rotateRightA() int {
	gb.A = gb.rotateRightInternal(gb.A)
	return 1
}

// setCarryFlag sets the carry flag and unsets N and C
func (gb *Gameboy) setCarryFlag(val bool) int {
	gb.setNFlag(false)
	gb.setHFlag(false)
	gb.setCFlag(val)
	return 1
}

/*
decimalAdjustA decimal-adjusts the number

this is nuts
*/
func (gb *Gameboy) decimalAdjustA() int {
	if !gb.getNFlag() {
		if gb.getCFlag() || gb.A > 0x99 {
			gb.A += 0x000_0060
			gb.setCFlag(true)
		}
		if gb.getHFlag() || (gb.A&0x0f) > 0x09 {
			gb.A += 0x000_0006
		}
	} else {
		if gb.getCFlag() {
			gb.A -= 0x000_0060
		}
		if gb.getHFlag() {
			gb.A += 0x000_0006
		}
	}

	gb.setZFlag(gb.A == 0)
	gb.setHFlag(false)
	return 1
}
