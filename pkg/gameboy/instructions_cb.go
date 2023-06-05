package gameboy

func (gb *gameboy) execCBInstr() int {

	opcode := gb.getImmediate() //fetch opcode

	switch opcode {

	}

	return -1
}

// rotateLeftCircular circular rotate a register left
func (regs *registers) rotateLeftCircular(reg *uint8) int {
	regs.rotateLeftCircularInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateLeft rotate the given register left
func (regs *registers) rotateLeft(reg *uint8) int {
	regs.rotateLeftInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateRightCircular circular rotate a register left
func (regs *registers) rotateRightCircular(reg *uint8) int {
	regs.rotateRightCircularInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}

// rotateRight rotate the given register left
func (regs *registers) rotateRight(reg *uint8) int {
	regs.rotateRightInternal(reg)
	regs.setZ(*reg == 0)
	return 1
}
