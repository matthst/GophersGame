package gameboy

import "fmt"

func (gb *Gameboy) execNextInstr() {

	opcode := gb.getImmediate() //fetch opcode

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */
	switch opcode {
	case 0x00:
		gb.nop()
	case 0x10:
		gb.stop()
	case 0x20:
		gb.JumpRelativeI8(!gb.getZFlag())
	case 0x30:
		gb.JumpRelativeI8(!gb.getNFlag())
	case 0x01:
		gb.loadI16(&gb.B, &gb.C)
	case 0x11:
		gb.loadI16(&gb.D, &gb.E)
	case 0x21:
		gb.loadI16(&gb.H, &gb.L)
	case 0x31:
		gb.loadI16SP()
	case 0x02:
		gb.storeR8(gb.A, gb.getBC())
	case 0x12:
		gb.storeR8(gb.A, gb.getDE())
	case 0x22:
		gb.storeR8(gb.A, getWordInc(&gb.H, &gb.L))
	case 0x32:
		gb.storeR8(gb.A, getWordDec(&gb.H, &gb.L))
	case 0x03:
		gb.incR16(&gb.B, &gb.C)
	case 0x13:
		gb.incR16(&gb.D, &gb.E)
	case 0x23:
		gb.incR16(&gb.H, &gb.L)
	case 0x33:
		gb.incSP()
	case 0x04:
		gb.incR8(&gb.B)
	case 0x14:
		gb.incR8(&gb.D)
	case 0x24:
		gb.incR8(&gb.H)
	case 0x34:
		gb.incM8(gb.getHL())
	case 0x05:
		gb.decR8(&gb.B)
	case 0x15:
		gb.decR8(&gb.D)
	case 0x25:
		gb.decR8(&gb.H)
	case 0x35:
		gb.decM8(gb.getHL())
	case 0x06:
		gb.loadI8(&gb.B)
	case 0x16:
		gb.loadI8(&gb.D)
	case 0x26:
		gb.loadI8(&gb.H)
	case 0x36:
		gb.storeI8()
	case 0x07:
		gb.rotateLeftCircularA()
	case 0x17:
		gb.rotateLeftA()
	case 0x27:
		gb.decimalAdjustA()
	case 0x37:
		gb.setCarryFlag(true)
	case 0x08:
		gb.storeSPI16()
	case 0x18:
		gb.JumpRelativeI8(true)
	case 0x28:
		gb.JumpRelativeI8(gb.getZFlag())
	case 0x38:
		gb.JumpRelativeI8(gb.getCFlag())
	case 0x09:
		gb.addR16R16(gb.getBC())
	case 0x19:
		gb.addR16R16(gb.getDE())
	case 0x29:
		gb.addR16R16(gb.getHL())
	case 0x39:
		gb.addR16R16(gb.SP)
	case 0x0A:
		gb.loadR8(&gb.A, gb.getBC())
	case 0x1A:
		gb.loadR8(&gb.A, gb.getDE())
	case 0x2A:
		gb.loadR8(&gb.A, getWordInc(&gb.H, &gb.L))
	case 0x3A:
		gb.loadR8(&gb.A, getWordDec(&gb.H, &gb.L))
	case 0x0B:
		gb.decR16(&gb.B, &gb.C)
	case 0x1B:
		gb.decR16(&gb.D, &gb.E)
	case 0x2B:
		gb.decR16(&gb.H, &gb.L)
	case 0x3B:
		gb.decSP()
	case 0x0C:
		gb.incR8(&gb.C)
	case 0x1C:
		gb.incR8(&gb.E)
	case 0x2C:
		gb.incR8(&gb.L)
	case 0x3C:
		gb.incR8(&gb.A)
	case 0x0D:
		gb.decR8(&gb.C)
	case 0x1D:
		gb.decR8(&gb.E)
	case 0x2D:
		gb.decR8(&gb.L)
	case 0x3D:
		gb.decR8(&gb.A)
	case 0x0E:
		gb.loadI8(&gb.C)
	case 0x1E:
		gb.loadI8(&gb.E)
	case 0x2E:
		gb.loadI8(&gb.L)
	case 0x3E:
		gb.loadI8(&gb.A)
	case 0x0F:
		gb.rotateRightCircularA()
	case 0x1F:
		gb.rotateRightA()
	case 0x2F:
		complementR8(&gb.A)
	case 0x3F:
		gb.setCarryFlag(!gb.getCFlag())
	case 0x40:
		gb.loadR8R8(&gb.B, &gb.B)
	case 0x50:
		gb.loadR8R8(&gb.D, &gb.B)
	case 0x60:
		gb.loadR8R8(&gb.H, &gb.B)
	case 0x70:
		gb.storeR8(gb.B, gb.getHL())
	case 0x41:
		gb.loadR8R8(&gb.B, &gb.C)
	case 0x51:
		gb.loadR8R8(&gb.D, &gb.C)
	case 0x61:
		gb.loadR8R8(&gb.H, &gb.C)
	case 0x71:
		gb.storeR8(gb.C, gb.getHL())
	case 0x42:
		gb.loadR8R8(&gb.B, &gb.D)
	case 0x52:
		gb.loadR8R8(&gb.D, &gb.D)
	case 0x62:
		gb.loadR8R8(&gb.H, &gb.D)
	case 0x72:
		gb.storeR8(gb.D, gb.getHL())
	case 0x43:
		gb.loadR8R8(&gb.B, &gb.E)
	case 0x53:
		gb.loadR8R8(&gb.D, &gb.E)
	case 0x63:
		gb.loadR8R8(&gb.H, &gb.E)
	case 0x73:
		gb.storeR8(gb.E, gb.getHL())
	case 0x44:
		gb.loadR8R8(&gb.B, &gb.H)
	case 0x54:
		gb.loadR8R8(&gb.D, &gb.H)
	case 0x64:
		gb.loadR8R8(&gb.H, &gb.H)
	case 0x74:
		gb.storeR8(gb.H, gb.getHL())
	case 0x45:
		gb.loadR8R8(&gb.B, &gb.L)
	case 0x55:
		gb.loadR8R8(&gb.D, &gb.L)
	case 0x65:
		gb.loadR8R8(&gb.H, &gb.L)
	case 0x75:
		gb.storeR8(gb.L, gb.getHL())
	case 0x46:
		gb.loadR8(&gb.B, gb.getHL())
	case 0x56:
		gb.loadR8(&gb.D, gb.getHL())
	case 0x66:
		gb.loadR8(&gb.H, gb.getHL())
	case 0x76:
		gb.halt()
	case 0x47:
		gb.loadR8R8(&gb.B, &gb.A)
	case 0x57:
		gb.loadR8R8(&gb.D, &gb.A)
	case 0x67:
		gb.loadR8R8(&gb.H, &gb.A)
	case 0x77:
		gb.storeR8(gb.A, gb.getHL())
	case 0x48:
		gb.loadR8R8(&gb.C, &gb.B)
	case 0x58:
		gb.loadR8R8(&gb.E, &gb.B)
	case 0x68:
		gb.loadR8R8(&gb.L, &gb.B)
	case 0x78:
		gb.loadR8R8(&gb.A, &gb.B)
	case 0x49:
		gb.loadR8R8(&gb.C, &gb.C)
	case 0x59:
		gb.loadR8R8(&gb.E, &gb.C)
	case 0x69:
		gb.loadR8R8(&gb.L, &gb.C)
	case 0x79:
		gb.loadR8R8(&gb.A, &gb.C)
	case 0x4A:
		gb.loadR8R8(&gb.C, &gb.D)
	case 0x5A:
		gb.loadR8R8(&gb.E, &gb.D)
	case 0x6A:
		gb.loadR8R8(&gb.L, &gb.D)
	case 0x7A:
		gb.loadR8R8(&gb.A, &gb.D)
	case 0x4B:
		gb.loadR8R8(&gb.C, &gb.E)
	case 0x5B:
		gb.loadR8R8(&gb.E, &gb.E)
	case 0x6B:
		gb.loadR8R8(&gb.L, &gb.E)
	case 0x7B:
		gb.loadR8R8(&gb.A, &gb.E)
	case 0x4C:
		gb.loadR8R8(&gb.C, &gb.H)
	case 0x5C:
		gb.loadR8R8(&gb.E, &gb.H)
	case 0x6C:
		gb.loadR8R8(&gb.L, &gb.H)
	case 0x7C:
		gb.loadR8R8(&gb.A, &gb.H)
	case 0x4D:
		gb.loadR8R8(&gb.C, &gb.L)
	case 0x5D:
		gb.loadR8R8(&gb.E, &gb.L)
	case 0x6D:
		gb.loadR8R8(&gb.L, &gb.L)
	case 0x7D:
		gb.loadR8R8(&gb.A, &gb.L)
	case 0x4E:
		gb.loadR8(&gb.C, gb.getHL())
	case 0x5E:
		gb.loadR8(&gb.E, gb.getHL())
	case 0x6E:
		gb.loadR8(&gb.L, gb.getHL())
	case 0x7E:
		gb.loadR8(&gb.A, gb.getHL())
	case 0x4F:
		gb.loadR8R8(&gb.C, &gb.A)
	case 0x5F:
		gb.loadR8R8(&gb.E, &gb.A)
	case 0x6F:
		gb.loadR8R8(&gb.L, &gb.A)
	case 0x7F:
		gb.loadR8R8(&gb.A, &gb.A)
	case 0x80:
		gb.addR8(gb.B)
	case 0x90:
		gb.subR8(gb.B)
	case 0xA0:
		gb.andR8(gb.B)
	case 0xB0:
		gb.orR8(gb.B)
	case 0x81:
		gb.addR8(gb.C)
	case 0x91:
		gb.subR8(gb.C)
	case 0xA1:
		gb.andR8(gb.C)
	case 0xB1:
		gb.orR8(gb.C)
	case 0x82:
		gb.addR8(gb.D)
	case 0x92:
		gb.subR8(gb.D)
	case 0xA2:
		gb.andR8(gb.D)
	case 0xB2:
		gb.orR8(gb.D)
	case 0x83:
		gb.addR8(gb.E)
	case 0x93:
		gb.subR8(gb.E)
	case 0xA3:
		gb.andR8(gb.E)
	case 0xB3:
		gb.orR8(gb.E)
	case 0x84:
		gb.addR8(gb.H)
	case 0x94:
		gb.subR8(gb.H)
	case 0xA4:
		gb.andR8(gb.H)
	case 0xB4:
		gb.orR8(gb.H)
	case 0x85:
		gb.addR8(gb.L)
	case 0x95:
		gb.subR8(gb.L)
	case 0xA5:
		gb.andR8(gb.L)
	case 0xB5:
		gb.orR8(gb.L)
	case 0x86:
		gb.aluM8(gb.addR8)
	case 0x96:
		gb.aluM8(gb.subR8)
	case 0xA6:
		gb.aluM8(gb.andR8)
	case 0xB6:
		gb.aluM8(gb.orR8)
	case 0x87:
		gb.addR8(gb.A)
	case 0x97:
		gb.subR8(gb.A)
	case 0xA7:
		gb.andR8(gb.A)
	case 0xB7:
		gb.orR8(gb.A)
	case 0x88:
		gb.adcR8(gb.B)
	case 0x98:
		gb.sbcR8(gb.B)
	case 0xA8:
		gb.xorR8(gb.B)
	case 0xB8:
		gb.cpR8(gb.B)
	case 0x89:
		gb.adcR8(gb.C)
	case 0x99:
		gb.sbcR8(gb.C)
	case 0xA9:
		gb.xorR8(gb.C)
	case 0xB9:
		gb.cpR8(gb.C)
	case 0x8A:
		gb.adcR8(gb.D)
	case 0x9A:
		gb.sbcR8(gb.D)
	case 0xAA:
		gb.xorR8(gb.D)
	case 0xBA:
		gb.cpR8(gb.D)
	case 0x8B:
		gb.adcR8(gb.E)
	case 0x9B:
		gb.sbcR8(gb.E)
	case 0xAB:
		gb.xorR8(gb.E)
	case 0xBB:
		gb.cpR8(gb.E)
	case 0x8C:
		gb.adcR8(gb.H)
	case 0x9C:
		gb.sbcR8(gb.H)
	case 0xAC:
		gb.xorR8(gb.H)
	case 0xBC:
		gb.cpR8(gb.H)
	case 0x8D:
		gb.adcR8(gb.L)
	case 0x9D:
		gb.sbcR8(gb.L)
	case 0xAD:
		gb.xorR8(gb.L)
	case 0xBD:
		gb.cpR8(gb.L)
	case 0x8E:
		gb.aluM8(gb.adcR8)
	case 0x9E:
		gb.aluM8(gb.sbcR8)
	case 0xAE:
		gb.aluM8(gb.xorR8)
	case 0xBE:
		gb.aluM8(gb.cpR8)
	case 0x8F:
		gb.adcR8(gb.A)
	case 0x9F:
		gb.sbcR8(gb.A)
	case 0xAF:
		gb.xorR8(gb.A)
	case 0xBF:
		gb.cpR8(gb.A)
	case 0xC0:
		gb.retCond(!gb.getZFlag())
	case 0xD0:
		gb.retCond(!gb.getCFlag())
	case 0xE0:
		gb.storeAI8()
	case 0xF0:
		gb.loadAI8()
	case 0xC1:
		gb.pop(&gb.B, &gb.C)
	case 0xD1:
		gb.pop(&gb.D, &gb.E)
	case 0xE1:
		gb.pop(&gb.H, &gb.L)
	case 0xF1:
		gb.pop(&gb.A, &gb.F)
	case 0xC2:
		gb.JumpI16(!gb.getZFlag())
	case 0xD2:
		gb.JumpI16(!gb.getCFlag())
	case 0xE2:
		gb.storeAC()
	case 0xF2:
		gb.loadAC()
	case 0xC3:
		gb.JumpI16(true)
	case 0xF3:
		gb.disableInterrupts()
	case 0xC4:
		gb.call(!gb.getZFlag())
	case 0xD4:
		gb.call(!gb.getCFlag())
	case 0xC6:
		gb.aluI8(gb.addR8)
	case 0xD6:
		gb.aluI8(gb.subR8)
	case 0xE6:
		gb.aluI8(gb.andR8)
	case 0xF6:
		gb.aluI8(gb.orR8)
	case 0xC5:
		gb.push(gb.B, gb.C)
	case 0xD5:
		gb.push(gb.D, gb.E)
	case 0xE5:
		gb.push(gb.H, gb.L)
	case 0xF5:
		gb.push(gb.A, gb.F)
	case 0xC7:
		gb.rst(0x00)
	case 0xD7:
		gb.rst(0x10)
	case 0xE7:
		gb.rst(0x20)
	case 0xF7:
		gb.rst(0x30)
	case 0xC8:
		gb.retCond(gb.getZFlag())
	case 0xD8:
		gb.retCond(gb.getCFlag())
	case 0xE8:
		gb.addSPS8SP()
	case 0xF8:
		gb.addSPS8HL()
	case 0xC9:
		gb.ret()
	case 0xD9:
		gb.retInterrupt()
	case 0xE9:
		gb.jumpHL()
	case 0xF9:
		gb.loadHLSP()
	case 0xCA:
		gb.JumpI16(gb.getZFlag())
	case 0xDA:
		gb.JumpI16(gb.getCFlag())
	case 0xEA:
		gb.storeAMI16()
	case 0xFA:
		gb.loadAMI16()
	case 0xCB:
		gb.execCBInstr()
	case 0xFB:
		gb.enableInterrupts()
	case 0xCC:
		gb.call(gb.getZFlag())
	case 0xDC:
		gb.call(gb.getCFlag())
	case 0xCD:
		gb.call(true)
	case 0xCE:
		gb.aluM8(gb.adcR8)
	case 0xDE:
		gb.aluM8(gb.sbcR8)
	case 0xEE:
		gb.aluM8(gb.xorR8)
	case 0xFE:
		gb.aluM8(gb.cpR8)
	case 0xCF:
		gb.rst(0x08)
	case 0xDF:
		gb.rst(0x18)
	case 0xEF:
		gb.rst(0x28)
	case 0xFF:
		gb.rst(0x38)
	default:
		panic(fmt.Sprintf("Opcode '%X' is not a valid opcode", opcode))
	}
}

func (gb *Gameboy) nop() {
}

func (gb *Gameboy) stop() int {
	// TODO implement the stop instruction
	panic("Stop instruction (0x10) not implemented")
}

func (gb *Gameboy) halt() {
	if gb.IE&gb.IF&0x1F != 0 {
		gb.haltBug = true
	} else {
		gb.haltMode = true
	}
}

func (gb *Gameboy) disableInterrupts() {
	gb.IME = false
}

func (gb *Gameboy) enableInterrupts() {
	if gb.EICounter == 0 {
		gb.EICounter = 2
	}
}

// JumpI16 conditional jump
func (gb *Gameboy) JumpI16(flag bool) {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if flag {
		gb.PC = getWord(hi, lo)
		gb.tick()
	}
}

// JumpRelativeI8 relative conditional jump
func (gb *Gameboy) JumpRelativeI8(flag bool) {
	im8 := gb.getImmediate()
	if flag {
		gb.PC += uint16(im8)
		gb.tick()
	}
}

func (gb *Gameboy) jumpHL() {
	gb.SP = gb.getHL()
}

func (gb *Gameboy) retCond(cond bool) {
	gb.tick()
	if cond {
		P := gb.load(gb.getAndIncSP())
		S := gb.load(gb.getAndIncSP())
		gb.PC = getWord(S, P)
		gb.tick()
	}
}

func (gb *Gameboy) ret() {
	P := gb.load(gb.getAndIncSP())
	S := gb.load(gb.getAndIncSP())
	gb.PC = getWord(S, P)
	gb.tick()
}

func (gb *Gameboy) retInterrupt() {
	gb.IME = true
	gb.ret()
}

func (gb *Gameboy) call(cond bool) {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	if cond {
		gb.tick()
		gb.write(hi, gb.getAndDecSP())
		gb.write(lo, gb.getAndDecSP())
	}
}

func (gb *Gameboy) rst(adr uint16) {
	P, C := getBytes(gb.PC)
	gb.tick()
	gb.write(P, gb.getAndDecSP())
	gb.write(C, gb.getAndDecSP())
	gb.PC = adr
}

// loadI8 load an 8-bit immediate into a register
func (gb *Gameboy) loadI8(reg *uint8) {
	*reg = gb.getImmediate()
}

// loadR8R8 copy r2 into r1
func (gb *Gameboy) loadR8R8(r1, r2 *uint8) {
	*r1 = *r2
	gb.tick()
}

// loadI16 load a 16-bit immediate into a register
func (gb *Gameboy) loadI16(hi, lo *uint8) {
	*lo = gb.getImmediate()
	*hi = gb.getImmediate()
}

// loadHLSP load the value of SP into HL
func (gb *Gameboy) loadHLSP() {
	setBytes(&gb.H, &gb.L, gb.SP)
	gb.tick()
}

func (gb *Gameboy) loadI16SP() {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.SP = getWord(hi, lo)
}

// loadR8 load an 8-bit val from memory into the given register
func (gb *Gameboy) loadR8(reg *uint8, adr uint16) {
	*reg = gb.load(adr)
}

// loadMAI16 load an 8-bit val into A from the memory address specified by the 16-bit immediate
func (gb *Gameboy) loadAMI16() {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.A = gb.load(getWord(hi, lo))
}

// loadAI8 load the content of an address in the block 0xFF00 - 0xFFFF given by an i8 into register A
func (gb *Gameboy) loadAI8() {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.A = gb.load(adr)
}

// loadAC load the content of an address in the block 0xFF00 - 0xFFFF given by register C into register A
func (gb *Gameboy) loadAC() {
	adr := uint16(0xFF00) | uint16(gb.C)
	gb.A = gb.load(adr)
}

// storeR8 store the content of register at regVal in the address specified by RegAdr.
func (gb *Gameboy) storeR8(val uint8, adr uint16) {
	gb.write(val, adr)
}

// storeSPI16 store the stack pointer in the memory address provided by the 16-bit immediate
func (gb *Gameboy) storeSPI16() {
	adr := uint16(gb.getImmediate())
	adr += uint16(gb.getImmediate()) << 8
	S, P := getBytes(gb.SP)
	gb.write(P, adr)
	gb.write(S, adr+1)
}

// storeI8 store the immediate 8-bit value into the memory address specified by HL
func (gb *Gameboy) storeI8() {
	gb.write(gb.getImmediate(), gb.getHL())
}

// storeAI8 store the content of register A in an address in the block 0xFF00 - 0xFFFF given by an i8
func (gb *Gameboy) storeAI8() {
	adr := uint16(0xFF00) | uint16(gb.getImmediate())
	gb.write(gb.A, adr)
}

// storeAMI16 store an 8-bit val from A into the memory address specified by the 16-bit immediate
func (gb *Gameboy) storeAMI16() {
	lo := gb.getImmediate()
	hi := gb.getImmediate()
	gb.write(gb.A, getWord(hi, lo))
}

// storeAC store the content of register A in an address in the block 0xFF00 - 0xFFFF given by C
func (gb *Gameboy) storeAC() {
	gb.write(gb.A, uint16(0xFF00)|uint16(gb.C))
}

func (gb *Gameboy) push(hi, lo uint8) {
	gb.tick()
	gb.write(hi, gb.getAndDecSP())
	gb.write(lo, gb.getAndDecSP())
}

// pop load a 16bit value from memory and increment the stack pointer during the load (twice in total)
func (gb *Gameboy) pop(hi, lo *uint8) {
	*lo = gb.load(gb.getAndIncSP())
	*hi = gb.load(gb.getAndIncSP())
}

// incR16 increments a combine 16-bit register.
func (gb *Gameboy) incR16(hi, lo *uint8) {
	setBytes(hi, lo, getWord(*hi, *lo)+1)
	gb.tick()
}

// incSP increments a combine 16-bit register.
func (gb *Gameboy) incSP() {
	gb.SP++
	gb.tick()
}

// incR8 increment the given 8-bit register
func (gb *Gameboy) incR8(reg *uint8) {
	gb.setNFlag(false)
	gb.setHFlag(halfCarryAddCheck8Bit(*reg, 1))
	*reg++
	gb.setZFlag(*reg == 0)
}

// incM8 increment the 8 bit value at the specified memory address
func (gb *Gameboy) incM8(adr uint16) {
	val := gb.load(adr)
	gb.setNFlag(false)
	gb.setHFlag(halfCarryAddCheck8Bit(val, 1))
	val++
	gb.setZFlag(val == 0)
	gb.write(val, adr)
}

// decR16 increments a combine 16-bit register.
func (gb *Gameboy) decR16(hi, lo *uint8) {
	setBytes(hi, lo, getWord(*hi, *lo)-1)
	gb.tick()
}

// decSP increments a combine 16-bit register.
func (gb *Gameboy) decSP() {
	gb.SP--
	gb.tick()
}

// decR8 decrement the given 8-bit register
func (gb *Gameboy) decR8(reg *uint8) {
	gb.setNFlag(true)
	gb.setHFlag(halfCarrySubCheck8Bit(*reg, 1))
	*reg--
	gb.setZFlag(*reg == 0)
}

// decM8 decrement the 8 bit value at the specified memory address
func (gb *Gameboy) decM8(adr uint16) {
	val := gb.load(adr)
	gb.setNFlag(true)
	gb.setHFlag(halfCarrySubCheck8Bit(val, 1))
	val--
	gb.setZFlag(val == 0)
	gb.write(val, adr)
}

// addR16R16 add the contents of one 16-bit register pair to the register HL
func (gb *Gameboy) addR16R16(val uint16) {
	HL := gb.getHL()
	gb.setHFlag(halfCarryAddCheck16Bit(HL, val))
	gb.setCFlag(HL+val < HL)
	gb.setNFlag(false)
	setBytes(&gb.H, &gb.L, HL+val)
	gb.tick()
}

// addSPS8SP add the signed 2's complement immediate to the stack pointer and write it to HL
func (gb *Gameboy) addSPS8SP() {
	gb.SP = gb.addSPS8Internal()
	gb.tick()
}

// addSPS8HL add the signed 2's complement immediate to the stack pointer and write it to HL
func (gb *Gameboy) addSPS8HL() {
	setBytes(&gb.H, &gb.L, gb.addSPS8Internal())
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
	gb.tick()
	return gb.SP - uint16(val)
}

// addR8 add the 8-bit value of a register to A
func (gb *Gameboy) addR8(val uint8) {
	a := gb.A
	gb.setFlags(a+val == 0, false, halfCarryAddCheck8Bit(a, val), a+val < a)
	gb.A += val
}

// adcR8 add the 8-bit value of a register to A
func (gb *Gameboy) adcR8(val uint8) {
	if gb.getCFlag() {
		gb.addR8(val + 1)
	}
	gb.addR8(val)
}

// subR8 subtract the 8-bit value of a register from A
func (gb *Gameboy) subR8(val uint8) {
	a := gb.A
	gb.setFlags(a+val == 0, false, halfCarrySubCheck8Bit(a, val), a-val > a)
	gb.A -= val
}

// sbcR8 subtract the 8-bit value of a register from A
func (gb *Gameboy) sbcR8(val uint8) {
	if gb.getCFlag() {
		gb.subR8(val + 1)
	}
	gb.subR8(val)
}

// andR8 logical AND the 8-bit value of a register with A
func (gb *Gameboy) andR8(val uint8) {
	gb.A &= val
	gb.setFlags(gb.A == 0, false, true, false)
}

// orR8 logical OR the 8-bit value of a register with A
func (gb *Gameboy) orR8(val uint8) {
	gb.A |= val
	gb.setFlags(gb.A == 0, false, false, false)
}

// xorR8 logical XOR the 8-bit value of a register with A
func (gb *Gameboy) xorR8(val uint8) {
	gb.A ^= val
	gb.setFlags(gb.A == 0, false, false, false)
}

// cpR8 compare the 8-bit value of a register with A
func (gb *Gameboy) cpR8(val uint8) {
	gb.setFlags(gb.A+val == 0, false, halfCarrySubCheck8Bit(gb.A, val), gb.A-val > gb.A)
}

// aluR8Def function definition of an 8-bit alu function
type aluR8Def func(uint8)

// aluI8 executes an 8-bit alu function with the given immediate
func (gb *Gameboy) aluI8(aluFunc aluR8Def) {
	aluFunc(gb.getImmediate())
}

// aluM8 executes an 8-bit alu function with the value from the given memory address
func (gb *Gameboy) aluM8(aluFunc aluR8Def) {
	val := gb.load(gb.getHL())
	aluFunc(val)
}

// complementR8 bit-swap the register
func complementR8(r *uint8) {
	*r = ^*r
}

// rotateLeftCircularA circular rotate register A left
func (gb *Gameboy) rotateLeftCircularA() {
	gb.A = gb.rotateLeftCircularInternal(gb.A)
}

// rotateLeftA rotate register A left
func (gb *Gameboy) rotateLeftA() {
	gb.A = gb.rotateLeftInternal(gb.A)
}

// rotateRightCircularA circular rotate register A left
func (gb *Gameboy) rotateRightCircularA() {
	gb.A = gb.rotateRightCircularInternal(gb.A)
}

// rotateRightA rotate register A left
func (gb *Gameboy) rotateRightA() {
	gb.A = gb.rotateRightInternal(gb.A)
}

// setCarryFlag sets the carry flag and unsets N and C
func (gb *Gameboy) setCarryFlag(val bool) {
	gb.setNFlag(false)
	gb.setHFlag(false)
	gb.setCFlag(val)
}

/*
decimalAdjustA decimal-adjusts the number

this is nuts
*/
func (gb *Gameboy) decimalAdjustA() {
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
}
