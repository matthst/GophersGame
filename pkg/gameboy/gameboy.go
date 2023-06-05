package gameboy

type gameboy struct {
	regs *registers
	mem  *memory
}

func (gb *gameboy) execNextInstr() int {

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */

	opcode := gb.getImmediate() //fetch opcode

	switch opcode {
	case 0x00:
		return nop()
	case 0x10:
		return stop()
	case 0x20:
		return gb.JumpI8(!gb.regs.getZ())
	case 0x30:
		return gb.JumpI8(!gb.regs.getN())
	case 0x01:
		return gb.loadI16(&gb.regs.BC)
	case 0x11:
		return gb.loadI16(&gb.regs.DE)
	case 0x21:
		return gb.loadI16(&gb.regs.HL)
	case 0x31:
		return gb.loadI16(&gb.regs.SP)
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
		return incR16(&gb.regs.SP)
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
		return gb.JumpI8(true)
	case 0x28:
		return gb.JumpI8(gb.regs.getZ())
	case 0x38:
		return gb.JumpI8(gb.regs.getC())
	case 0x09:
		return gb.regs.addR16R16(&gb.regs.HL, &gb.regs.BC)
	case 0x19:
		return gb.regs.addR16R16(&gb.regs.HL, &gb.regs.DE)
	case 0x29:
		return gb.regs.addR16R16(&gb.regs.HL, &gb.regs.HL)
	case 0x39:
		return gb.regs.addR16R16(&gb.regs.HL, &gb.regs.SP)
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
		return decR16(&gb.regs.SP)
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
	}
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

// JumpI8 conditional jump
func (gb *gameboy) JumpI8(flag bool) int {
	im8 := gb.getImmediate()
	if flag {
		gb.regs.PC += uint16(im8)
		return 3
	}
	return 2
}

// load an 8-bit immediate into a register
func (gb *gameboy) loadI8(reg *uint8) int {
	*reg = gb.getImmediate()
	return 2
}

func loadR8R8(r1, r2 *uint8) int {
	*r1 = *r2
	return 1
}

// load a 16-bit immediate into a register
func (gb *gameboy) loadI16(reg *[2]uint8) int {
	reg[0] = gb.getImmediate()
	reg[1] = gb.getImmediate()
	return 3
}

func (gb *gameboy) loadR8(reg *uint8, adr uint16) int {
	*reg = gb.mem.load(adr)
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
	gb.mem.store(gb.regs.SP[0], adr)
	gb.mem.store(gb.regs.SP[1], adr+1)
	return 5
}

// storeI8 store the immediate 8-bit value into the memory address specified by the 16-bit register
func (gb *gameboy) storeI8(reg *[2]uint8) int {
	gb.mem.store(gb.getImmediate(), getWord(reg))
	return 3
}

// incR16 increments a combine 16-bit register.
func incR16(reg *[2]uint8) int {
	setWord(reg, getWord(reg)+1)
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

// addR8R8 add the contents of one 8-bit register to another
func (gb *gameboy) addR8R8(reg1, reg2 *uint8) int {

	return 1
}

// addR16R16 add the contents of one 16-bit register pair to another
func (regs *registers) addR16R16(reg1, reg2 *[2]uint8) int {
	a := getWord(reg1)
	b := getWord(reg2)
	regs.setH(halfCarryAddCheck16Bit(a, b))
	regs.setC(CarryAddCheck16Bit(a, b))
	regs.setN(false)
	setWord(reg1, a+b)
	return 2
}

func complementR8(r *uint8) int {
	*r ^= 0
	return 1
}

// rotateLeftCircularA circular rotate register A left
func (regs *registers) rotateLeftCircularA() int {
	regs.rotateLeftCircularInternal(&regs.AF[1])
	return 1
}

// rotateLeftCircular circular rotate a register left
func (regs *registers) rotateLeftCircular(reg *uint8) int {
	regs.rotateLeftCircularInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateLeftA rotate register A left
func (regs *registers) rotateLeftA() int {
	regs.rotateLeftInternal(&regs.AF[1])
	return 1
}

// rotateLeft rotate the given register left
func (regs *registers) rotateLeft(reg *uint8) int {
	regs.rotateLeftInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateRightCircularA circular rotate register A left
func (regs *registers) rotateRightCircularA() int {
	regs.rotateRightCircularInternal(&regs.AF[1])
	return 1
}

// rotateRightCircular circular rotate a register left
func (regs *registers) rotateRightCircular(reg *uint8) int {
	regs.rotateRightCircularInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateRightA rotate register A left
func (regs *registers) rotateRightA() int {
	regs.rotateRightInternal(&regs.AF[1])
	return 1
}

// rotateRight rotate the given register left
func (regs *registers) rotateRight(reg *uint8) int {
	regs.rotateRightInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// setCarryFlag sets the carry flag and unsets N and C
func (regs *registers) setCarryFlag(val bool) int {
	regs.setN(false)
	regs.setH(false)
	regs.setC(val)
	return 1
}

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
