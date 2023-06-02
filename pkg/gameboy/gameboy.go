package gameboy

type gameboy struct {
	regs *registers
}

func (gb *gameboy) execNextInstr() {

	/* the opcodes are sorted in pairs of four, the pattern is clear once you look at the opcode table */

	opcode := gb.getNextVal()

	switch opcode {
	case 0x00:
		nop()
	case 0x10:
		stop()
	case 0x20:
		jr()
	case 0x30:
	case 0x01:
	case 0x11:
	case 0x21:
	case 0x31:
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

func jr(flag bool) {
	// TODO
}
