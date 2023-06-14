package gameboy

import "fmt"

func execNextInstr() {

	opcode := getImmediate() //fetch opcode

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */
	switch opcode {
	case 0x00:
		nop()
	case 0x10:
		stop()
	case 0x20:
		jumpRelativeI8(!getZFlag())
	case 0x30:
		jumpRelativeI8(!getCFlag())
	case 0x01:
		loadI16(&bReg, &cReg)
	case 0x11:
		loadI16(&dReg, &eReg)
	case 0x21:
		loadI16(&hReg, &lReg)
	case 0x31:
		loadI16SP()
	case 0x02:
		storeR8(aReg, getWord(bReg, cReg))
	case 0x12:
		storeR8(aReg, getWord(dReg, eReg))
	case 0x22:
		storeR8(aReg, getWordInc(&hReg, &lReg))
	case 0x32:
		storeR8(aReg, getWordDec(&hReg, &lReg))
	case 0x03:
		incR16(&bReg, &cReg)
	case 0x13:
		incR16(&dReg, &eReg)
	case 0x23:
		incR16(&hReg, &lReg)
	case 0x33:
		incSP()
	case 0x04:
		incR8(&bReg)
	case 0x14:
		incR8(&dReg)
	case 0x24:
		incR8(&hReg)
	case 0x34:
		incM8(getHL())
	case 0x05:
		decR8(&bReg)
	case 0x15:
		decR8(&dReg)
	case 0x25:
		decR8(&hReg)
	case 0x35:
		decM8(getHL())
	case 0x06:
		loadI8(&bReg)
	case 0x16:
		loadI8(&dReg)
	case 0x26:
		loadI8(&hReg)
	case 0x36:
		storeI8()
	case 0x07:
		rotateLeftCircularA()
	case 0x17:
		rotateLeftA()
	case 0x27:
		decimalAdjustA()
	case 0x37:
		setCarryFlag(true)
	case 0x08:
		storeSPI16()
	case 0x18:
		jumpRelativeI8(true)
	case 0x28:
		jumpRelativeI8(getZFlag())
	case 0x38:
		jumpRelativeI8(getCFlag())
	case 0x09:
		addR16R16(getWord(bReg, cReg))
	case 0x19:
		addR16R16(getWord(dReg, eReg))
	case 0x29:
		addR16R16(getHL())
	case 0x39:
		addR16R16(SP)
	case 0x0A:
		loadR8(&aReg, getWord(bReg, cReg))
	case 0x1A:
		loadR8(&aReg, getWord(dReg, eReg))
	case 0x2A:
		loadR8(&aReg, getWordInc(&hReg, &lReg))
	case 0x3A:
		loadR8(&aReg, getWordDec(&hReg, &lReg))
	case 0x0B:
		decR16(&bReg, &cReg)
	case 0x1B:
		decR16(&dReg, &eReg)
	case 0x2B:
		decR16(&hReg, &lReg)
	case 0x3B:
		decSP()
	case 0x0C:
		incR8(&cReg)
	case 0x1C:
		incR8(&eReg)
	case 0x2C:
		incR8(&lReg)
	case 0x3C:
		incR8(&aReg)
	case 0x0D:
		decR8(&cReg)
	case 0x1D:
		decR8(&eReg)
	case 0x2D:
		decR8(&lReg)
	case 0x3D:
		decR8(&aReg)
	case 0x0E:
		loadI8(&cReg)
	case 0x1E:
		loadI8(&eReg)
	case 0x2E:
		loadI8(&lReg)
	case 0x3E:
		loadI8(&aReg)
	case 0x0F:
		rotateRightCircularA()
	case 0x1F:
		rotateRightA()
	case 0x2F:
		complementR8(&aReg)
	case 0x3F:
		setCarryFlag(!getCFlag())
	case 0x40:
		loadR8R8(&bReg, &bReg)
	case 0x50:
		loadR8R8(&dReg, &bReg)
	case 0x60:
		loadR8R8(&hReg, &bReg)
	case 0x70:
		storeR8(bReg, getHL())
	case 0x41:
		loadR8R8(&bReg, &cReg)
	case 0x51:
		loadR8R8(&dReg, &cReg)
	case 0x61:
		loadR8R8(&hReg, &cReg)
	case 0x71:
		storeR8(cReg, getHL())
	case 0x42:
		loadR8R8(&bReg, &dReg)
	case 0x52:
		loadR8R8(&dReg, &dReg)
	case 0x62:
		loadR8R8(&hReg, &dReg)
	case 0x72:
		storeR8(dReg, getHL())
	case 0x43:
		loadR8R8(&bReg, &eReg)
	case 0x53:
		loadR8R8(&dReg, &eReg)
	case 0x63:
		loadR8R8(&hReg, &eReg)
	case 0x73:
		storeR8(eReg, getHL())
	case 0x44:
		loadR8R8(&bReg, &hReg)
	case 0x54:
		loadR8R8(&dReg, &hReg)
	case 0x64:
		loadR8R8(&hReg, &hReg)
	case 0x74:
		storeR8(hReg, getHL())
	case 0x45:
		loadR8R8(&bReg, &lReg)
	case 0x55:
		loadR8R8(&dReg, &lReg)
	case 0x65:
		loadR8R8(&hReg, &lReg)
	case 0x75:
		storeR8(lReg, getHL())
	case 0x46:
		loadR8(&bReg, getHL())
	case 0x56:
		loadR8(&dReg, getHL())
	case 0x66:
		loadR8(&hReg, getHL())
	case 0x76:
		halt()
	case 0x47:
		loadR8R8(&bReg, &aReg)
	case 0x57:
		loadR8R8(&dReg, &aReg)
	case 0x67:
		loadR8R8(&hReg, &aReg)
	case 0x77:
		storeR8(aReg, getHL())
	case 0x48:
		loadR8R8(&cReg, &bReg)
	case 0x58:
		loadR8R8(&eReg, &bReg)
	case 0x68:
		loadR8R8(&lReg, &bReg)
	case 0x78:
		loadR8R8(&aReg, &bReg)
	case 0x49:
		loadR8R8(&cReg, &cReg)
	case 0x59:
		loadR8R8(&eReg, &cReg)
	case 0x69:
		loadR8R8(&lReg, &cReg)
	case 0x79:
		loadR8R8(&aReg, &cReg)
	case 0x4A:
		loadR8R8(&cReg, &dReg)
	case 0x5A:
		loadR8R8(&eReg, &dReg)
	case 0x6A:
		loadR8R8(&lReg, &dReg)
	case 0x7A:
		loadR8R8(&aReg, &dReg)
	case 0x4B:
		loadR8R8(&cReg, &eReg)
	case 0x5B:
		loadR8R8(&eReg, &eReg)
	case 0x6B:
		loadR8R8(&lReg, &eReg)
	case 0x7B:
		loadR8R8(&aReg, &eReg)
	case 0x4C:
		loadR8R8(&cReg, &hReg)
	case 0x5C:
		loadR8R8(&eReg, &hReg)
	case 0x6C:
		loadR8R8(&lReg, &hReg)
	case 0x7C:
		loadR8R8(&aReg, &hReg)
	case 0x4D:
		loadR8R8(&cReg, &lReg)
	case 0x5D:
		loadR8R8(&eReg, &lReg)
	case 0x6D:
		loadR8R8(&lReg, &lReg)
	case 0x7D:
		loadR8R8(&aReg, &lReg)
	case 0x4E:
		loadR8(&cReg, getHL())
	case 0x5E:
		loadR8(&eReg, getHL())
	case 0x6E:
		loadR8(&lReg, getHL())
	case 0x7E:
		loadR8(&aReg, getHL())
	case 0x4F:
		loadR8R8(&cReg, &aReg)
	case 0x5F:
		loadR8R8(&eReg, &aReg)
	case 0x6F:
		loadR8R8(&lReg, &aReg)
	case 0x7F:
		loadR8R8(&aReg, &aReg)
	case 0x80:
		addR8(bReg)
	case 0x90:
		subR8(bReg)
	case 0xA0:
		andR8(bReg)
	case 0xB0:
		orR8(bReg)
	case 0x81:
		addR8(cReg)
	case 0x91:
		subR8(cReg)
	case 0xA1:
		andR8(cReg)
	case 0xB1:
		orR8(cReg)
	case 0x82:
		addR8(dReg)
	case 0x92:
		subR8(dReg)
	case 0xA2:
		andR8(dReg)
	case 0xB2:
		orR8(dReg)
	case 0x83:
		addR8(eReg)
	case 0x93:
		subR8(eReg)
	case 0xA3:
		andR8(eReg)
	case 0xB3:
		orR8(eReg)
	case 0x84:
		addR8(hReg)
	case 0x94:
		subR8(hReg)
	case 0xA4:
		andR8(hReg)
	case 0xB4:
		orR8(hReg)
	case 0x85:
		addR8(lReg)
	case 0x95:
		subR8(lReg)
	case 0xA5:
		andR8(lReg)
	case 0xB5:
		orR8(lReg)
	case 0x86:
		aluM8(addR8)
	case 0x96:
		aluM8(subR8)
	case 0xA6:
		aluM8(andR8)
	case 0xB6:
		aluM8(orR8)
	case 0x87:
		addR8(aReg)
	case 0x97:
		subR8(aReg)
	case 0xA7:
		andR8(aReg)
	case 0xB7:
		orR8(aReg)
	case 0x88:
		adcR8(bReg)
	case 0x98:
		sbcR8(bReg)
	case 0xA8:
		xorR8(bReg)
	case 0xB8:
		cpR8(bReg)
	case 0x89:
		adcR8(cReg)
	case 0x99:
		sbcR8(cReg)
	case 0xA9:
		xorR8(cReg)
	case 0xB9:
		cpR8(cReg)
	case 0x8A:
		adcR8(dReg)
	case 0x9A:
		sbcR8(dReg)
	case 0xAA:
		xorR8(dReg)
	case 0xBA:
		cpR8(dReg)
	case 0x8B:
		adcR8(eReg)
	case 0x9B:
		sbcR8(eReg)
	case 0xAB:
		xorR8(eReg)
	case 0xBB:
		cpR8(eReg)
	case 0x8C:
		adcR8(hReg)
	case 0x9C:
		sbcR8(hReg)
	case 0xAC:
		xorR8(hReg)
	case 0xBC:
		cpR8(hReg)
	case 0x8D:
		adcR8(lReg)
	case 0x9D:
		sbcR8(lReg)
	case 0xAD:
		xorR8(lReg)
	case 0xBD:
		cpR8(lReg)
	case 0x8E:
		aluM8(adcR8)
	case 0x9E:
		aluM8(sbcR8)
	case 0xAE:
		aluM8(xorR8)
	case 0xBE:
		aluM8(cpR8)
	case 0x8F:
		adcR8(aReg)
	case 0x9F:
		sbcR8(aReg)
	case 0xAF:
		xorR8(aReg)
	case 0xBF:
		cpR8(aReg)
	case 0xC0:
		retCond(!getZFlag())
	case 0xD0:
		retCond(!getCFlag())
	case 0xE0:
		storeAI8()
	case 0xF0:
		loadAI8()
	case 0xC1:
		pop(&bReg, &cReg)
	case 0xD1:
		pop(&dReg, &eReg)
	case 0xE1:
		pop(&hReg, &lReg)
	case 0xF1:
		pop(&aReg, &fReg)
	case 0xC2:
		jumpI16(!getZFlag())
	case 0xD2:
		jumpI16(!getCFlag())
	case 0xE2:
		storeAC()
	case 0xF2:
		loadAC()
	case 0xC3:
		jumpI16(true)
	case 0xF3:
		disableInterrupts()
	case 0xC4:
		call(!getZFlag())
	case 0xD4:
		call(!getCFlag())
	case 0xC6:
		aluI8(addR8)
	case 0xD6:
		aluI8(subR8)
	case 0xE6:
		aluI8(andR8)
	case 0xF6:
		aluI8(orR8)
	case 0xC5:
		push(bReg, cReg)
	case 0xD5:
		push(dReg, eReg)
	case 0xE5:
		push(hReg, lReg)
	case 0xF5:
		push(aReg, fReg)
	case 0xC7:
		rst(0x00)
	case 0xD7:
		rst(0x10)
	case 0xE7:
		rst(0x20)
	case 0xF7:
		rst(0x30)
	case 0xC8:
		retCond(getZFlag())
	case 0xD8:
		retCond(getCFlag())
	case 0xE8:
		addSPS8SP()
	case 0xF8:
		addSPS8HL()
	case 0xC9:
		ret()
	case 0xD9:
		retInterrupt()
	case 0xE9:
		jumpHL()
	case 0xF9:
		loadHLSP()
	case 0xCA:
		jumpI16(getZFlag())
	case 0xDA:
		jumpI16(getCFlag())
	case 0xEA:
		storeAMI16()
	case 0xFA:
		loadAMI16()
	case 0xCB:
		execCBInstr()
	case 0xFB:
		enableInterrupts()
	case 0xCC:
		call(getZFlag())
	case 0xDC:
		call(getCFlag())
	case 0xCD:
		call(true)
	case 0xCE:
		aluI8(adcR8)
	case 0xDE:
		aluI8(sbcR8)
	case 0xEE:
		aluI8(xorR8)
	case 0xFE:
		aluI8(cpR8)
	case 0xCF:
		rst(0x08)
	case 0xDF:
		rst(0x18)
	case 0xEF:
		rst(0x28)
	case 0xFF:
		rst(0x38)
	default:
		panic(fmt.Sprintf("Opcode '%X' is not a valid opcode", opcode))
	}
}

func nop() {
}

func stop() int {
	// TODO implement the stop instruction
	panic("Stop instruction (0x10) not implemented")
}

func halt() {
	if IE&IF&0x1F != 0 {
		haltBug = true
	} else {
		haltMode = true
	}
}

func disableInterrupts() {
	IME = false
}

func enableInterrupts() {
	if EICounter == 0 {
		EICounter = 2
	}
}

// JumpI16 conditional jump
func jumpI16(flag bool) {
	lo := getImmediate()
	hi := getImmediate()
	if flag {
		PC = getWord(hi, lo)
		mCycle()
	}
}

// JumpRelativeI8 relative conditional jump
func jumpRelativeI8(flag bool) {
	im8 := getImmediate()
	if flag {
		if im8 < 128 {
			PC += uint16(im8)
		} else {
			PC -= uint16(^im8 + 1)
		}
		mCycle()
	}
}

func jumpHL() {
	PC = getHL()
}

func retCond(cond bool) {
	mCycle()
	if cond {
		P := memConLoad(getAndIncSP())
		S := memConLoad(getAndIncSP())
		PC = getWord(S, P)
		mCycle()
	}
}

func ret() {
	C := memConLoad(getAndIncSP())
	P := memConLoad(getAndIncSP())
	PC = getWord(P, C)
	mCycle()
}

func retInterrupt() {
	IME = true
	ret()
}

func call(cond bool) {
	lo := getImmediate()
	hi := getImmediate()
	if cond {
		P, C := getBytes(PC)
		mCycle()
		SP--
		memConWrite(P, SP)
		SP--
		memConWrite(C, SP)
		PC = getWord(hi, lo)
	}
}

func rst(adr uint16) {
	P, C := getBytes(PC)
	mCycle()
	SP--
	memConWrite(P, SP)
	SP--
	memConWrite(C, SP)
	PC = adr
}

// loadI8 load an 8-bit immediate into a register
func loadI8(reg *uint8) {
	*reg = getImmediate()
}

// loadR8R8 copy r2 into r1
func loadR8R8(r1, r2 *uint8) {
	*r1 = *r2
	mCycle()
}

// loadI16 load a 16-bit immediate into a register
func loadI16(hi, lo *uint8) {
	*lo = getImmediate()
	*hi = getImmediate()
}

// loadHLSP load the contents of register pair HL into the stack pointer SP
func loadHLSP() {
	SP = getWord(hReg, lReg)
	mCycle()
}

func loadI16SP() {
	lo := getImmediate()
	hi := getImmediate()
	SP = getWord(hi, lo)
}

// loadR8 load an 8-bit val from memory into the given register
func loadR8(reg *uint8, adr uint16) {
	*reg = memConLoad(adr)
}

// loadMAI16 load an 8-bit val into aReg from the memory address specified by the 16-bit immediate
func loadAMI16() {
	lo := getImmediate()
	hi := getImmediate()
	aReg = memConLoad(getWord(hi, lo))
}

// loadAI8 load the content of an address in the block 0xFF00 - 0xFFFF given by an i8 into register A
func loadAI8() {
	adr := uint16(0xFF00) | uint16(getImmediate())
	aReg = memConLoad(adr)
}

// loadAC load the content of an address in the block 0xFF00 - 0xFFFF given by register C into register A
func loadAC() {
	adr := uint16(0xFF00) | uint16(cReg)
	aReg = memConLoad(adr)
}

// storeR8 store the content of register at regVal in the address specified by RegAdr.
func storeR8(val uint8, adr uint16) {
	memConWrite(val, adr)
}

// storeSPI16 store the stack pointer in the memory address provided by the 16-bit immediate
func storeSPI16() {
	adr := uint16(getImmediate())
	adr += uint16(getImmediate()) << 8
	s, p := getBytes(SP)
	memConWrite(p, adr)
	memConWrite(s, adr+1)
}

// storeI8 store the immediate 8-bit value into the memory address specified by HL
func storeI8() {
	memConWrite(getImmediate(), getHL())
}

// storeAI8 store the content of register A in an address in the block 0xFF00 - 0xFFFF given by an i8
func storeAI8() {
	adr := uint16(0xFF00) | uint16(getImmediate())
	memConWrite(aReg, adr)
}

// storeAMI16 store an 8-bit val from aReg into the memory address specified by the 16-bit immediate
func storeAMI16() {
	lo := getImmediate()
	hi := getImmediate()
	memConWrite(aReg, getWord(hi, lo))
}

// storeAC store the content of register A in an address in the block 0xFF00 - 0xFFFF given by A
func storeAC() {
	memConWrite(aReg, uint16(0xFF00)|uint16(cReg))
}

func push(hi, lo uint8) {
	mCycle()
	SP--
	memConWrite(hi, SP)
	SP--
	memConWrite(lo, SP)
}

// pop load a 16bit value from memory and increment the stack pointer during the load (twice in total)
func pop(hi, lo *uint8) {
	*lo = memConLoad(getAndIncSP())
	*hi = memConLoad(getAndIncSP())
}

// incR16 increments a combine 16-bit register.
func incR16(hi, lo *uint8) {
	setBytes(hi, lo, getWord(*hi, *lo)+1)
	mCycle()
}

// incSP increments a combine 16-bit register.
func incSP() {
	SP++
	mCycle()
}

// incR8 increment the given 8-bit register
func incR8(reg *uint8) {
	setNFlag(false)
	setHFlag(halfCarryAddCheck8Bit(*reg, 1))
	*reg++
	setZFlag(*reg == 0)
}

// incM8 increment the 8 bit value at the specified memory address
func incM8(adr uint16) {
	val := memConLoad(adr)
	setNFlag(false)
	setHFlag(halfCarryAddCheck8Bit(val, 1))
	val++
	setZFlag(val == 0)
	memConWrite(val, adr)
}

// decR16 increments a combine 16-bit register.
func decR16(hi, lo *uint8) {
	setBytes(hi, lo, getWord(*hi, *lo)-1)
	mCycle()
}

// decSP increments a combine 16-bit register.
func decSP() {
	SP--
	mCycle()
}

// decR8 decrement the given 8-bit register
func decR8(reg *uint8) {
	setNFlag(true)
	setHFlag(halfCarrySubCheck8Bit(*reg, 1))
	*reg--
	setZFlag(*reg == 0)
}

// decM8 decrement the 8 bit value at the specified memory address
func decM8(adr uint16) {
	val := memConLoad(adr)
	setNFlag(true)
	setHFlag(halfCarrySubCheck8Bit(val, 1))
	val--
	setZFlag(val == 0)
	memConWrite(val, adr)
}

// addR16R16 add the contents of one 16-bit register pair to the register HL
func addR16R16(val uint16) {
	HL := getHL()
	setHFlag(halfCarryAddCheck16Bit(HL, val))
	setCFlag(HL+val < HL)
	setNFlag(false)
	setBytes(&hReg, &lReg, HL+val)
	mCycle()
}

// addSPS8SP add the signed 2's complement immediate to the stack pointer and write it to HL
func addSPS8SP() {
	SP = addSPS8Internal()
	mCycle()
}

// addSPS8HL add the signed 2's complement immediate to the stack pointer and write it to HL
func addSPS8HL() {
	setBytes(&hReg, &lReg, addSPS8Internal())
}

// addSPS8 add the signed 2's complement immediate to the stack pointer and return the value
func addSPS8Internal() uint16 {
	val := getImmediate()
	P := uint8(SP)
	setZFlag(false)
	setNFlag(false)
	if val < 128 { // positive 2's complement value :=
		setHFlag(halfCarryAddCheck8Bit(P, val))
		setCFlag(P+val < P)
		return SP + uint16(val)
	}
	// negative 2's complement value
	val = ^val + 1 //get positive value from 2's complement signed number
	setHFlag(halfCarrySubCheck8Bit(P, val))
	setCFlag(P-val > P)
	mCycle()
	return SP - uint16(val)
}

// addR8 add the 8-bit value of a register to A
func addR8(val uint8) {
	setFlags(aReg+val == 0, false, halfCarryAddCheck8Bit(aReg, val), aReg+val < aReg)
	aReg += val
}

// adcR8 add the 8-bit value of a register to A
func adcR8(val uint8) {
	if getCFlag() {
		val += 1
	}
	addR8(val)
}

// subR8 subtract the 8-bit value of a register from A
func subR8(val uint8) {
	setFlags(aReg-val == 0, true, halfCarrySubCheck8Bit(aReg, val), aReg-val > aReg)
	aReg -= val
}

// sbcR8 subtract the 8-bit value of a register from A
func sbcR8(val uint8) {
	if getCFlag() {
		val += 1
	}
	subR8(val)
}

// andR8 logical AND the 8-bit value of a register with A
func andR8(val uint8) {
	aReg &= val
	setFlags(aReg == 0, false, true, false)
}

// orR8 logical OR the 8-bit value of a register with A
func orR8(val uint8) {
	aReg |= val
	setFlags(aReg == 0, false, false, false)
}

// xorR8 logical XOR the 8-bit value of a register with A
func xorR8(val uint8) {
	aReg ^= val
	setFlags(aReg == 0, false, false, false)
}

// cpR8 compare the 8-bit value of a register with A
func cpR8(val uint8) {
	setFlags(aReg-val == 0, true, halfCarrySubCheck8Bit(aReg, val), aReg-val > aReg)
}

// aluR8Def function definition of an 8-bit alu function
type aluR8Def func(uint8)

// aluI8 executes an 8-bit alu function with the given immediate
func aluI8(aluFunc aluR8Def) {
	aluFunc(getImmediate())
}

// aluM8 executes an 8-bit alu function with the value from the given memory address
func aluM8(aluFunc aluR8Def) {
	val := memConLoad(getHL())
	aluFunc(val)
}

// complementR8 bit-swap the register
func complementR8(r *uint8) {
	*r = ^*r
}

// rotateLeftCircularA circular rotate register A left
func rotateLeftCircularA() {
	aReg = rotateLeftCircularInternal(aReg)
}

// rotateLeftA rotate register aReg left
func rotateLeftA() {
	aReg = rotateLeftInternal(aReg)
}

// rotateRightCircularA circular rotate register A left
func rotateRightCircularA() {
	aReg = rotateRightCircularInternal(aReg)
}

// rotateRightA rotate register aReg left
func rotateRightA() {
	aReg = rotateRightInternal(aReg)
}

// setCarryFlag sets the carry flag and unsets N and C
func setCarryFlag(val bool) {
	setNFlag(false)
	setHFlag(false)
	setCFlag(val)
}

/*
decimalAdjustA decimal-adjusts the number

this is nuts
*/
func decimalAdjustA() {
	if !getNFlag() {
		if getCFlag() || aReg > 0x99 {
			aReg += 0x000_0060
			setCFlag(true)
		}
		if getHFlag() || (aReg&0x0f) > 0x09 {
			aReg += 0x000_0006
		}
	} else {
		if getCFlag() {
			aReg -= 0x000_0060
		}
		if getHFlag() {
			aReg += 0x000_0006
		}
	}

	setZFlag(aReg == 0)
	setHFlag(false)
}
