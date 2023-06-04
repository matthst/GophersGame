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

// load a 16-bit immediate into a register
func (gb *gameboy) loadI16(arr *[2]uint8) {
	gb.clockCycle()
	arr[0] = gb.getNextVal()
	gb.clockCycle()
	arr[1] = gb.getNextVal()
	gb.clockCycle()
}

// store the content of register at regVal in the address specified by RegAdr
// function definition is provided to make incrementing of registers at the right time possible
// values are fetched with pointers so their value is read at the right time
func (gb *gameboy) storeR8(regVal *uint8, regAdr *[2]uint8, getWord getWordDef) {
	gb.clockCycle()
	gb.mem.store(*regVal, getWord(regAdr))
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
