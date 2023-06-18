package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/Input"
	Cartridge "github.com/matthst/gophersgame/pkg/gameboy/cartridge"
	"github.com/matthst/gophersgame/pkg/gameboy/components"
	"github.com/matthst/gophersgame/pkg/gameboy/timer"
	"github.com/matthst/gophersgame/pkg/gameboy/video"
	"strings"
)

var (
	audioC  components.Audio
	wramC   components.WRAM
	hramC   components.HRAM
	SerialC components.Serial
	cart    Cartridge.Cartridge
	Vid     video.Video

	oamSourceOffset, oamByteIndex uint16

	mCycleCounter, mCycleOffset int

	SP                                             uint16
	PC                                             uint16
	aReg, fReg, bReg, cReg, dReg, eReg, hReg, lReg uint8
	EICounter, IE, IF                              uint8
	IME, haltMode, haltBug, oamTransferActive      bool
)

func Bootstrap(file []uint8, romPath string, serialBuilder *strings.Builder) {
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

	EICounter = 0x00
	IE, IF = 0x00, 0xE1
	IME, haltMode, haltBug = false, false, false

	Timer.DividerClk = 0xABCC

	switch file[0x0147] {
	case 0x00, 0x01:
		cart = Cartridge.CreateMBC1(file, romPath, false, false)
	case 0x02:
		cart = Cartridge.CreateMBC1(file, romPath, true, false)
	case 0x03:
		cart = Cartridge.CreateMBC1(file, romPath, true, true)
	case 0x05:
		cart = Cartridge.CreateMBC2(file, romPath, false)
	case 0x06:
		cart = Cartridge.CreateMBC2(file, romPath, true)
	default:
		panic(fmt.Sprintf("Cartridge Type '%X' not implemented", file[0x0147]))
	}

	audioC = components.Audio{}
	Vid = video.GetDmgVideo()
	wramC = components.WRAM{}
	hramC = components.HRAM{}
	SerialC = components.Serial{StringBuilder: serialBuilder}

}

func RunOneTick() {
	Input.RunTick()
	mCycleCounter = mCycleOffset
	for mCycleCounter < 17556 {
		runInstructionCycle()
	}

	mCycleOffset = mCycleCounter - 17556
}

func mCycle() {
	IF |= Vid.MCycle()
	IF |= Timer.Cycle()
	IF |= SerialC.Cycle()
	IF |= Input.Cycle()
	oamTransferCycle()
	mCycleCounter++
}

func runInstructionCycle() {
	interruptServiceRoutine()

	if !haltMode {
		execNextInstr()
	} else {
		mCycle()
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
	val := loadAndCycle(PC)
	if !haltBug {
		PC++
	}
	mCycle()
	return val
}

// writeAndCycle to the memory controller
func writeAndCycle(val uint8, adr uint16) {
	defer mCycle()
	writeToMemoryBus(val, adr)
}

// loadAndCycle from the memory controller
func loadAndCycle(adr uint16) uint8 {
	defer mCycle()
	return loadFromMemoryBus(adr)
}

func oamTransferCycle() {
	if oamTransferActive {
		val := loadFromMemoryBus(0xFE00 + oamByteIndex)
		writeToMemoryBus(val, oamSourceOffset+oamByteIndex)
		oamByteIndex++
		if oamByteIndex == 160 {
			oamTransferActive = false
			oamByteIndex = 0
		}
	}
}

func startOAMTransfer(val uint8) {
	if !oamTransferActive {
		oamTransferActive = true
		oamSourceOffset = uint16(val) << 8
		oamByteIndex = 0
	}
}

func writeToMemoryBus(val uint8, adr uint16) {
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
		return

	// I/O Registers
	case adr == 0xFF00: // Input
		Input.Write(val)
	case adr == 0xFF46:
		startOAMTransfer(val)
	case adr == 0xFF0F: // IF
		IF = (val & 0x1F) | 0b1110_0000
	case adr == 0xFFFF: // IE
		IE = val
	case adr == 0xFF4D: // KEY1
		return // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF03: // serial port
		SerialC.Write(val, adr)
	case adr < 0xFF0F: // timer control
		Timer.Write(val, adr)
	case adr < 0xFF40: // audio + wave RAM
		audioC.Write(val, adr)
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		Vid.WriteToIORegisters(val, adr)
	case adr == 0xFF70:
		return // TODO [CGB] WRAM bank switch
	case adr >= 0xFF80:
		hramC.Write(val, adr)
	default:
		// panic(fmt.Sprintf("CPU tried to read from memory address '%X', but no implementation exists.", adr))
	}
}

func loadFromMemoryBus(adr uint16) uint8 {
	switch {

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
	case adr == 0xFF00: // Input
		return Input.Load()
	case adr == 0xFF0F: // IF
		return IF | 0b1110_0000
	case adr == 0xFFFF: // IE
		return IE
	case adr == 0xFF70: // KEY1
		return 1 // TODO [CGB] WRAM bank switch
	case adr < 0xFF03: // serial port
		return SerialC.Load(adr)
	case adr < 0xFF0F: // timer control
		return Timer.Load(adr)
	case adr < 0xFF40: // audio + wave RAM
		return audioC.Load(adr)
	case adr == 0xFF4D: //CG
		return 1 // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		return Vid.LoadFromIORegisters(adr)
	case adr >= 0xFF80:
		return hramC.Load(adr)

	}
	// fmt.Printf("CPU tried to write to memory address '%X', but no implementation exists. \n", adr)
	return 0x00
}
