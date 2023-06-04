package gameboy

type gameboy struct {
	regs *registers
	mem  *memory
}

func (gb *gameboy) execNextInstr() {

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */

	opcode := gb.getNextVal() //fetch opcode

	switch opcode {
	case 0x00:
		nop()
	case 0x10:
		stop()
	case 0x20:
		gb.jrc(!gb.regs.getZ())
	case 0x30:
		gb.jrc(!gb.regs.getN())
	case 0x01:
		gb.loadI16(&gb.regs.BC)
	case 0x11:
		gb.loadI16(&gb.regs.DE)
	case 0x21:
		gb.loadI16(&gb.regs.HL)
	case 0x31:
		gb.loadI16(&gb.regs.SP)
	case 0x02:
		gb.storeR8(&gb.regs.AF[1], &gb.regs.BC, getWord)
	case 0x12:
		gb.storeR8(&gb.regs.AF[1], &gb.regs.DE, getWord)
	case 0x22:
		gb.storeR8(&gb.regs.AF[1], &gb.regs.HL, getWordInc)
	case 0x32:
		gb.storeR8(&gb.regs.AF[1], &gb.regs.HL, getWordDec)
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
	}

}

func (gb *gameboy) getNextVal() uint8 {
	return 0
}

func (gb *gameboy) clockCycle() {
}

func nop() {
	// TODO
}

func stop() {
	// TODO
}

// conditional jump with immediate
func (gb *gameboy) jrc(flag bool) {
	gb.clockCycle()
	im8 := gb.getNextVal()
	gb.clockCycle()
	if flag {
		gb.regs.PC += uint16(im8)
		gb.clockCycle()
	}
}

// load an 8-bit immediate into a register
func (gb *gameboy) loadI8(reg *uint8) {
	gb.clockCycle()
	*reg = gb.getNextVal()
	gb.clockCycle()
}

// load a 16-bit immediate into a register
func (gb *gameboy) loadI16(reg *[2]uint8) {
	gb.clockCycle()
	reg[0] = gb.getNextVal()
	gb.clockCycle()
	reg[1] = gb.getNextVal()
	gb.clockCycle()
}

/*
store the content of register at regVal in the address specified by RegAdr.

	function definition is provided to make incrementing of registers at the right time possible,
	values are fetched with pointers so their value is read at the right time
*/
func (gb *gameboy) storeR8(regVal *uint8, regAdr *[2]uint8, getWordFunc getWordDef) {
	gb.clockCycle()
	gb.mem.store(*regVal, getWordFunc(regAdr))
	gb.clockCycle()
}

// store the immediate 8-bit value into the memory address specified by the 16-bit register
func (gb *gameboy) storeI8(reg *[2]uint8) {
	gb.clockCycle()
	val := gb.getNextVal()
	gb.clockCycle()
	gb.mem.store(val, getWord(reg))
	gb.clockCycle()
}

// increments a combine 16-bit register. Implements correctly how the CPU *probably* does it.
func (gb *gameboy) incR16(reg *[2]uint8) {
	reg[0]++
	gb.clockCycle()
	if reg[0] == 0 { //overflow handling
		reg[1]++
	}
	gb.clockCycle()
}

// increment the given 8-bit register
func (gb *gameboy) incR8(reg *uint8) {
	gb.regs.setN(false)
	gb.regs.setH(halfCarryAddCheck(*reg, 1))
	*reg++
	gb.regs.setZ(*reg == 0)
	gb.clockCycle()
}

// increment the 8 bit value at the specified memory address
func (gb *gameboy) incM8(adr uint16) {
	gb.clockCycle()
	val := gb.mem.load(adr)
	gb.clockCycle()
	gb.regs.setN(false)
	gb.regs.setH(halfCarryAddCheck(val, 1))
	val++
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	gb.clockCycle()
}

// decrement the given 8-bit register
func (gb *gameboy) decR8(reg *uint8) {
	gb.regs.setN(true)
	gb.regs.setH(halfCarrySubCheck(*reg, 1))
	*reg--
	gb.regs.setZ(*reg == 0)
	gb.clockCycle()
}

// decrement the 8 bit value at the specified memory address
func (gb *gameboy) decM8(adr uint16) {
	gb.clockCycle()
	val := gb.mem.load(adr)
	gb.clockCycle()
	gb.regs.setN(true)
	gb.regs.setH(halfCarrySubCheck(val, 1))
	val--
	gb.regs.setZ(val == 0)
	gb.mem.store(val, adr)
	gb.clockCycle()
}

// circular rotate register A left
func (gb *gameboy) rotateLeftCircularA() {
	gb.regs.rotateLeftCircularInternal(&gb.regs.AF[1])
	gb.clockCycle()
}

// circular rotate a register left
func (gb *gameboy) rotateLeftCircular(reg *uint8) {
	gb.regs.rotateLeftCircularInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle()
}

// rotate register A left
func (gb *gameboy) rotateLeftA() {
	gb.regs.rotateLeftCircularInternal(&gb.regs.AF[1])
	gb.clockCycle()
}

// rotate the given register left
func (gb *gameboy) rotateLeft(reg *uint8) {
	gb.regs.rotateLeftCircularInternal(reg)
	gb.regs.setZ(*reg == 0)
	gb.clockCycle()
}

func (regs *registers) rotateLeftCircularInternal(reg *uint8) {
	regs.AF[0] = 0b0000_0000 + ((*reg >> 7) << 4)
	*reg = (*reg << 1) + (*reg >> 7)
}

func (gb *gameboy) rotateLeftInternal(reg *uint8) {
	val := *reg<<1 + gb.regs.getCarryValue()
	gb.regs.AF[1] = 0b0000_0000 + ((*reg >> 7) << 4)
	*reg = val
}

func (gb *gameboy) set(val bool) {
	gb.regs.setN(false)
	gb.regs.setH(false)
	gb.regs.setC(val)
	gb.clockCycle()
}

func halfCarryAddCheck(a, b uint8) bool {
	return (((a & 0xf) + (b & 0xf)) & 0x10) == 0x10
}

func halfCarrySubCheck(a, b uint8) bool {
	return (((a & 0xf) - (b & 0xf)) & 0x10) == 0x10
}
