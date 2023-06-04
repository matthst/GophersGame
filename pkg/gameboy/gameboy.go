package gameboy

type gameboy struct {
	regs *registers
	mem  *memory
}

func (gb *gameboy) execNextInstr() {

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */

	opcode := gb.getImmediate() //fetch opcode

	switch opcode {
	case 0x00:
		nop()
	case 0x10:
		stop()
	case 0x20:
		gb.JumpI8(!gb.regs.getZ())
	case 0x30:
		gb.JumpI8(!gb.regs.getN())
	case 0x01:
		gb.loadI16(&gb.regs.BC)
	case 0x11:
		gb.loadI16(&gb.regs.DE)
	case 0x21:
		gb.loadI16(&gb.regs.HL)
	case 0x31:
		gb.loadI16(&gb.regs.SP)
	case 0x02:
		gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.BC))
	case 0x12:
		gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.DE))
	case 0x22:
		gb.storeR8(gb.regs.AF[1], getWordInc(&gb.regs.HL))
	case 0x32:
		gb.storeR8(gb.regs.AF[1], getWordDec(&gb.regs.HL))
	case 0x03:
		gb.incR16(&gb.regs.BC)
	case 0x13:
		gb.incR16(&gb.regs.DE)
	case 0x23:
		gb.incR16(&gb.regs.HL)
	case 0x33:
		gb.incR16(&gb.regs.SP)
	case 0x04:
		gb.incR8(&gb.regs.BC[1])
	case 0x14:
		gb.incR8(&gb.regs.DE[1])
	case 0x24:
		gb.incR8(&gb.regs.HL[1])
	case 0x34:
		gb.incM8(getWord(&gb.regs.HL))
	case 0x05:
		gb.decR8(&gb.regs.BC[1])
	case 0x15:
		gb.decR8(&gb.regs.DE[1])
	case 0x25:
		gb.decR8(&gb.regs.HL[1])
	case 0x35:
		gb.decM8(getWord(&gb.regs.HL))
	case 0x06:
		gb.loadI8(&gb.regs.BC[1])
	case 0x16:
		gb.loadI8(&gb.regs.DE[1])
	case 0x26:
		gb.loadI8(&gb.regs.HL[1])
	case 0x36:
		gb.storeI8(&gb.regs.HL)
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
		gb.JumpI8(true)
	case 0x28:
		gb.JumpI8(gb.regs.getZ())
	case 0x38:
		gb.JumpI8(gb.regs.getC())
	case 0x09:
		gb.addR16R16(&gb.regs.HL, &gb.regs.BC)
	case 0x19:
		gb.addR16R16(&gb.regs.HL, &gb.regs.DE)
	case 0x29:
		gb.addR16R16(&gb.regs.HL, &gb.regs.HL)
	case 0x39:
		gb.addR16R16(&gb.regs.HL, &gb.regs.SP)
	case 0x0A:
		gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.BC))
	case 0x1A:
		gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.DE))
	case 0x2A:
		gb.loadR8(&gb.regs.AF[1], getWordInc(&gb.regs.HL))
	case 0x3A:
		gb.loadR8(&gb.regs.AF[1], getWordDec(&gb.regs.HL))
	case 0x0B:
		gb.decR16(&gb.regs.BC)
	case 0x1B:
		gb.decR16(&gb.regs.DE)
	case 0x2B:
		gb.decR16(&gb.regs.HL)
	case 0x3B:
		gb.decR16(&gb.regs.SP)
	case 0x0C:
		gb.incR8(&gb.regs.BC[0])
	case 0x1C:
		gb.incR8(&gb.regs.DE[0])
	case 0x2C:
		gb.incR8(&gb.regs.HL[0])
	case 0x3C:
		gb.incR8(&gb.regs.AF[1])
	case 0x0D:
		gb.decR8(&gb.regs.BC[0])
	case 0x1D:
		gb.decR8(&gb.regs.DE[0])
	case 0x2D:
		gb.decR8(&gb.regs.HL[0])
	case 0x3D:
		gb.decR8(&gb.regs.AF[1])
	case 0x0E:
		gb.loadI8(&gb.regs.BC[0])
	case 0x1E:
		gb.loadI8(&gb.regs.DE[0])
	case 0x2E:
		gb.loadI8(&gb.regs.HL[0])
	case 0x3E:
		gb.loadI8(&gb.regs.AF[1])
	case 0x0F:
		gb.rotateRightCircularA()
	case 0x1F:
		gb.rotateRightA()
	case 0x2F:
		gb.complementR8(&gb.regs.AF[1])
	case 0x3F:
		gb.setCarryFlag(!gb.regs.getC())
	case 0x40:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.BC[1])
	case 0x50:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.BC[1])
	case 0x60:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.BC[1])
	case 0x70:
		gb.storeR8(gb.regs.BC[1], getWord(&gb.regs.HL))
	case 0x41:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.BC[0])
	case 0x51:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.BC[0])
	case 0x61:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.BC[0])
	case 0x71:
		gb.storeR8(gb.regs.BC[0], getWord(&gb.regs.HL))
	case 0x42:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.DE[1])
	case 0x52:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.DE[1])
	case 0x62:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.DE[1])
	case 0x72:
		gb.storeR8(gb.regs.DE[1], getWord(&gb.regs.HL))
	case 0x43:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.DE[0])
	case 0x53:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.DE[0])
	case 0x63:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.DE[0])
	case 0x73:
		gb.storeR8(gb.regs.DE[0], getWord(&gb.regs.HL))
	case 0x44:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.HL[1])
	case 0x54:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.HL[1])
	case 0x64:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.HL[1])
	case 0x74:
		gb.storeR8(gb.regs.HL[1], getWord(&gb.regs.HL))
	case 0x45:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.HL[0])
	case 0x55:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.HL[0])
	case 0x65:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.HL[0])
	case 0x75:
		gb.storeR8(gb.regs.HL[0], getWord(&gb.regs.HL))
	case 0x46:
		gb.loadR8(&gb.regs.BC[1], getWord(&gb.regs.HL))
	case 0x56:
		gb.loadR8(&gb.regs.DE[1], getWord(&gb.regs.HL))
	case 0x66:
		gb.loadR8(&gb.regs.HL[1], getWord(&gb.regs.HL))
	case 0x76:
		gb.halt()
	case 0x47:
		gb.loadR8R8(&gb.regs.BC[1], &gb.regs.AF[1])
	case 0x57:
		gb.loadR8R8(&gb.regs.DE[1], &gb.regs.AF[1])
	case 0x67:
		gb.loadR8R8(&gb.regs.HL[1], &gb.regs.AF[1])
	case 0x77:
		gb.storeR8(gb.regs.AF[1], getWord(&gb.regs.HL))
	case 0x48:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.BC[1])
	case 0x58:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.BC[1])
	case 0x68:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.BC[1])
	case 0x78:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.BC[1])
	case 0x49:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.BC[0])
	case 0x59:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.BC[0])
	case 0x69:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.BC[0])
	case 0x79:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.BC[0])
	case 0x4A:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.DE[1])
	case 0x5A:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.DE[1])
	case 0x6A:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.DE[1])
	case 0x7A:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.DE[1])
	case 0x4B:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.DE[0])
	case 0x5B:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.DE[0])
	case 0x6B:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.DE[0])
	case 0x7B:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.DE[0])
	case 0x4C:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.HL[1])
	case 0x5C:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.HL[1])
	case 0x6C:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.HL[1])
	case 0x7C:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.HL[1])
	case 0x4D:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.HL[0])
	case 0x5D:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.HL[0])
	case 0x6D:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.HL[0])
	case 0x7D:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.HL[0])
	case 0x4E:
		gb.loadR8(&gb.regs.BC[0], getWord(&gb.regs.HL))
	case 0x5E:
		gb.loadR8(&gb.regs.DE[0], getWord(&gb.regs.HL))
	case 0x6E:
		gb.loadR8(&gb.regs.HL[0], getWord(&gb.regs.HL))
	case 0x7E:
		gb.loadR8(&gb.regs.AF[1], getWord(&gb.regs.HL))
	case 0x4F:
		gb.loadR8R8(&gb.regs.BC[0], &gb.regs.AF[1])
	case 0x5F:
		gb.loadR8R8(&gb.regs.DE[0], &gb.regs.AF[1])
	case 0x6F:
		gb.loadR8R8(&gb.regs.HL[0], &gb.regs.AF[1])
	case 0x7F:
		gb.loadR8R8(&gb.regs.AF[1], &gb.regs.AF[1])
	}

}

func (gb *gameboy) getImmediate() uint8 {
	return 0
}

func (gb *gameboy) clockCycle(mCycles int) {
}

func nop() {
	// TODO
}

func stop() {
	// TODO
}

func (gb *gameboy) halt() {
	// TODO
}

// JumpI8 conditional jump
func (gb *gameboy) JumpI8(flag bool) {
	im8 := gb.getImmediate()
	gb.clockCycle(2)
	if flag {
		gb.regs.PC += uint16(im8)
		gb.clockCycle(1)
	}
}

// load an 8-bit immediate into a register
func (gb *gameboy) loadI8(reg *uint8) {
	*reg = gb.getImmediate()
	gb.clockCycle(2)
}

func (gb *gameboy) loadR8R8(r1, r2 *uint8) {
	*r1 = *r2
	gb.clockCycle(1)
}

// load a 16-bit immediate into a register
func (gb *gameboy) loadI16(reg *[2]uint8) {
	reg[0] = gb.getImmediate()
	reg[1] = gb.getImmediate()
	gb.clockCycle(3)
}

func (gb *gameboy) loadR8(reg *uint8, adr uint16) {
	*reg = gb.mem.load(adr)
	gb.clockCycle(2)
}

// storeR8 store the content of register at regVal in the address specified by RegAdr.
func (gb *gameboy) storeR8(val uint8, adr uint16) {
	gb.mem.store(val, adr)
	gb.clockCycle(2)
}

// storeSPI16 store the stack pointer in the memory address provided by the 16-bit immediate
func (gb *gameboy) storeSPI16() {
	adr := uint16(gb.getImmediate())
	adr += uint16(gb.getImmediate()) << 8
	gb.mem.store(gb.regs.SP[0], adr)
	gb.mem.store(gb.regs.SP[1], adr+1)
	gb.clockCycle(5)
}

// storeI8 store the immediate 8-bit value into the memory address specified by the 16-bit register
func (gb *gameboy) storeI8(reg *[2]uint8) {
	gb.mem.store(gb.getImmediate(), getWord(reg))
	gb.clockCycle(3)
}

// incR16 increments a combine 16-bit register.
func (gb *gameboy) incR16(reg *[2]uint8) {
	setWord(reg, getWord(reg)+1)
	gb.clockCycle(2)
}

// incR8 increment the given 8-bit register
func (gb *gameboy) incR8(reg *uint8) {
	gb.regs.setN(false)
	gb.regs.setH(halfCarryAddCheck8Bit(*reg, 1))
	*reg++
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// incM8 increment the 8 bit value at the specified memory address
func (gb *gameboy) incM8(adr uint16) {
	val := gb.mem.load(adr)
	gb.regs.setN(false)
	gb.regs.setH(halfCarryAddCheck8Bit(val, 1))
	val++
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	gb.clockCycle(3)
}

// decR16 increments a combine 16-bit register.
func (gb *gameboy) decR16(reg *[2]uint8) {
	setWord(reg, getWord(reg)-1)
	gb.clockCycle(2)
}

// decR8 decrement the given 8-bit register
func (gb *gameboy) decR8(reg *uint8) {
	gb.regs.setN(true)
	gb.regs.setH(halfCarrySubCheck8Bit(*reg, 1))
	*reg--
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// decM8 decrement the 8 bit value at the specified memory address
func (gb *gameboy) decM8(adr uint16) {
	val := gb.mem.load(adr)
	gb.regs.setN(true)
	gb.regs.setH(halfCarrySubCheck8Bit(val, 1))
	val--
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	gb.clockCycle(3)
}

// addR16R16 add the contents of one 16-bit register pair to another
func (gb *gameboy) addR16R16(reg1, reg2 *[2]uint8) {
	a := getWord(reg1)
	b := getWord(reg2)
	gb.regs.setH(halfCarryAddCheck16Bit(a, b))
	gb.regs.setC(CarryAddCheck16Bit(a, b))
	gb.regs.setN(false)
	setWord(reg1, a+b)
	gb.clockCycle(2)
}

func (gb *gameboy) complementR8(r *uint8) {
	*r ^= 0
	gb.clockCycle(1)
}

// circular rotate register A left
func (gb *gameboy) rotateLeftCircularA() {
	gb.regs.rotateLeftCircularInternal(&gb.regs.AF[1])
	gb.clockCycle(1)
}

// circular rotate a register left
func (gb *gameboy) rotateLeftCircular(reg *uint8) {
	gb.regs.rotateLeftCircularInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// rotate register A left
func (gb *gameboy) rotateLeftA() {
	gb.regs.rotateLeftInternal(&gb.regs.AF[1])
	gb.clockCycle(1)
}

// rotate the given register left
func (gb *gameboy) rotateLeft(reg *uint8) {
	gb.regs.rotateLeftInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// circular rotate register A left
func (gb *gameboy) rotateRightCircularA() {
	gb.regs.rotateRightCircularInternal(&gb.regs.AF[1])
	gb.clockCycle(1)
}

// circular rotate a register left
func (gb *gameboy) rotateRightCircular(reg *uint8) {
	gb.regs.rotateRightCircularInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// rotate register A left
func (gb *gameboy) rotateRightA() {
	gb.regs.rotateRightInternal(&gb.regs.AF[1])
	gb.clockCycle(1)
}

// rotate the given register left
func (gb *gameboy) rotateRight(reg *uint8) {
	gb.regs.rotateRightInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle(1)
}

// setCarryFlag sets the carry flag and unsets N and C
func (gb *gameboy) setCarryFlag(val bool) {
	gb.regs.setN(false)
	gb.regs.setH(false)
	gb.regs.setC(val)
	gb.clockCycle(1)
}

func (gb *gameboy) decimalAdjustA() {
	regs := gb.regs
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
	gb.clockCycle(1)
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

func halfCarrySubCheck16Bit(a, b uint16) bool {
	return (((a & 0xf) - (b & 0xf)) & 0x10) == 0x10
}

func CarryAddCheck8Bit(a, b uint8) bool {
	return a+b < a
}

func CarryAddCheck16Bit(a, b uint16) bool {
	return a+b < a
}

func CarrySubCheck8Bit(a, b uint16) bool {
	return a-b > a
}

func CarrySubCheck16Bit(a, b uint16) bool {
	return a-b > a
}
