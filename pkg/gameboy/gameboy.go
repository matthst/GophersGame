package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/components"
	"github.com/matthst/gophersgame/pkg/gameboy/video"
	"os"
)

var (
	audioC components.Audio
	wramC  components.WRAM
	hramC  components.HRAM
	timerC components.Timer
	inputC components.Input
	cart   components.Cartridge
	Vid    video.Video

	log os.File

	mCycleCounter, mCycleOffset, opcodeExecuteCounter int

	SP                                             uint16
	PC                                             uint16
	aReg, fReg, bReg, cReg, dReg, eReg, hReg, lReg uint8
	EICounter, IE, IF                              uint8
	IME, haltMode, haltBug, stopMode               bool
)

func Bootstrap(file []uint8) {
	mCycleCounter = 0
	mCycleOffset = 0

	PC = 0x0100
	SP = 0xFFFE
	aReg = 0x01
	fReg = 0xB0
	bReg = 0x00
	cReg = 0x13
	dReg = 0x00
	eReg = 0xD8
	hReg = 0x01
	lReg = 0x4D

	logFile, _ := os.OpenFile("text.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	log = *logFile

	switch file[0x0147] {
	case 0x00, 0x01:
		cart = components.CreateRomOnly(file)
	default:
		panic(fmt.Sprintf("Cartridge Type '%X' not implemented", file[0x0147]))
	}

	Vid = video.GetDmgVideo()
	wramC = components.WRAM{}
	hramC = components.HRAM{}
}

func logLine() {
	s := fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", aReg, fReg, bReg, cReg, dReg, eReg, hReg, lReg, SP, PC, debugLoad(PC), debugLoad(PC+1), debugLoad(PC+2), debugLoad(PC+3))
	_, _ = log.WriteString(s)
}

func RunOneTick() {
	mCycleCounter = mCycleOffset
	for mCycleCounter < 17556 {
		runInstructionCycle()
	}

	mCycleOffset = mCycleCounter - 17556
}

func mCycle() {
	IF |= Vid.MCycle()
	mCycleCounter++

}

func runInstructionCycle() {
	interruptServiceRoutine()

	if !haltMode {
		logLine()
		execNextInstr()
		opcodeExecuteCounter++
	}

	if EICounter > 0 {
		if EICounter == 1 {
			IME = true
		}
		EICounter--
	}
}

func interruptServiceRoutine() {
	if IME {
		availableInterrupts := IE & IF & 0x1F
		var interruptAddress uint16
		switch {
		case availableInterrupts == 0:
			return
		case availableInterrupts&0b1 == 0b1:
			interruptAddress = 0x40
			IF &= 0b1111_1110
		case availableInterrupts&0b10 == 0b10:
			interruptAddress = 0x48
			IF &= 0b1111_1101
		case availableInterrupts&0b100 == 0b100:
			interruptAddress = 0x50
			IF &= 0b1111_1011
		case availableInterrupts&0b1000 == 0b1000:
			interruptAddress = 0x58
			IF &= 0b1111_0111
		case availableInterrupts&0b1_0000 == 0b1_0000:
			interruptAddress = 0x60
			IF &= 0b1110_1111
		default:
			return
		}
		IME = false
		if haltMode {
			haltMode = false
			mCycle()
			mCycle()
			mCycle()
			mCycle()
		}
		mCycle()
		rst(interruptAddress)
	} else if haltMode && IE&IF&0x1F != 0 {
		haltMode = false
		mCycle()
		mCycle()
		mCycle()
		mCycle()
	}
}

func getImmediate() uint8 {
	val := memConLoad(PC)
	if !haltBug {
		PC++
	}
	mCycle()
	return val
}

// memConWrite to the memory controller
func memConWrite(val uint8, adr uint16) {
	switch {
	case adr < 0x8000: // cart ROM
		cart.Write(val, adr)
	case adr < 0xA000: // VRAM
		Vid.WriteToVRAM(val, adr)
	case adr < 0xC000: // cart RAM
		cart.Write(val, adr)
	case adr < 0xE000:
		wramC.Write(val, adr)
	case adr < 0xFE00: //ECHO Ram
		wramC.Write(val, adr-0x2000)
	case adr < 0xFEA0: // OAM
		Vid.WriteToOAM(val, adr)
	case adr < 0xFF00: //OAM corruption bug
		return // TODO implement OAM corruption bug

	// I/O Registers
	case adr == 0xFF00: // input
		inputC.Write(val, adr)
	case adr < 0xFF03: // serial port
		return // TODO implement serial port
	case adr < 0xFF0F: // timer control
		timerC.Write(val, adr)
	case adr == 0xFF0F: // IF flag
		IF = val
	case adr < 0xFF40: // audio + wave RAM
		audioC.Write(val, adr)
	case adr == 0xFF4D: //CG
		return // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		Vid.WriteToIORegisters(val, adr)
	case adr == 0xFF70:
		return // TODO [CGB] WRAM bank switch
	case adr >= 0xFF80:
		hramC.Write(val, adr)
	case adr == 0xFFFF:
		IE = val
	default:
		panic(fmt.Sprintf("CPU tried to read from memory address '%X', but no implementation exists.", adr))
	}

	mCycle()
}

// memConLoad from the memory controller
func memConLoad(adr uint16) uint8 {
	defer mCycle()

	switch {
	case adr == 0xFF44: // TODO REMOVEME
		return 0x90
	case adr < 0x8000: // cartridge ROM
		return cart.Load(adr)
	case adr < 0xA000: // VRAM
		return Vid.LoadFromVRAM(adr)
	case adr < 0xC000: // cartridge RAM
		return cart.Load(adr)
	case adr < 0xE000:
		return wramC.Load(adr)
	case adr < 0xFE00: //ECHO Ram
		return wramC.Load(adr - 0x2000)
	case adr < 0xFEA0: // OAM
		return Vid.LoadFromOAM(adr)
	case adr < 0xFF00: //OAM corruption bug
		return 0 // TODO implement OAM corruption bug

	// I/O Registers
	case adr == 0xFF00: // input
		return inputC.Load(adr)
	case adr < 0xFF03: // serial port
		return 1 // TODO implement serial port
	case adr < 0xFF0F: // timer control
		return timerC.Load(adr)
	case adr == 0xFF0F: // IF flag
		return IF
	case adr < 0xFF40: // audio + wave RAM
		return audioC.Load(adr)
	case adr == 0xFF4D: //CG
		return 1 // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		return Vid.LoadFromIORegisters(adr)
	case adr == 0xFF70:
		return 1 // TODO [CGB] WRAM bank switch
	case adr >= 0xFF80:
		return hramC.Load(adr)
	case adr == 0xFFFF:
		return IE
	}

	panic(fmt.Sprintf("CPU tried to write to memory address '%X', but no implementation exists.", adr))
}
