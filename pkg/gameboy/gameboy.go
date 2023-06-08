package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/components"
	"github.com/matthst/gophersgame/pkg/gameboy/video"
)

type Gameboy struct {
	video     video.Video
	audio     components.Audio
	cartridge components.Cartridge
	wram      components.WRAM
	input     components.Input
	timer     components.Timer

	PC, SP                           uint16
	A, F, B, C, D, E, H, L           uint8
	EICounter, IE, IF                uint8
	IME, haltMode, haltBug, stopMode bool
}

func bootstrap(file []uint8) Gameboy {

	gb := Gameboy{PC: 0x0100, SP: 0xFFFE, A: 0x01, F: 0x00, B: 0x00, C: 0x13, D: 0x00, E: 0xD8, H: 0x01, L: 0x4D}

	switch file[0x0147] {
	case 0x00:
		gb.cartridge = components.RomOnly{Rom: file}
	default:
		panic(fmt.Sprintf("Opcode '%X' not implemented", file[0x0148]))
	}

	return gb
}

func (gb *Gameboy) tick() {

}

func (gb *Gameboy) runInstructionCycle() {

	gb.interruptServiceRoutine()

	if !gb.haltMode {
		gb.execNextInstr()
	}

	if gb.EICounter > 0 {
		if gb.EICounter == 1 {
			gb.IME = true
		}
		gb.EICounter--
	}
}

func (gb *Gameboy) interruptServiceRoutine() {
	if gb.IME {
		availableInterrupts := gb.IE & gb.IF & 0x1F
		var interruptAddress uint16
		switch availableInterrupts {
		case 0b0000_0000:
			return
		case 0b0000_0001:
			interruptAddress = 0x40
			gb.IF &= 0b000_0001
		case 0b0000_0010:
			interruptAddress = 0x48
			gb.IF &= 0b000_0010
		case 0b0000_0100:
			interruptAddress = 0x50
			gb.IF &= 0b000_0100
		case 0b0000_1000:
			interruptAddress = 0x58
			gb.IF &= 0b000_1000
		case 0b0001_0000:
			interruptAddress = 0x60
			gb.IF &= 0b001_0000
		default:
			return
		}
		gb.IME = false
		if gb.haltMode {
			gb.haltMode = false
			gb.tick()
			gb.tick()
			gb.tick()
			gb.tick()
		}
		gb.tick()
		gb.rst(interruptAddress)
	} else if gb.haltMode && gb.IE&gb.IF&0x1F != 0 {
		gb.haltMode = false
		gb.tick()
		gb.tick()
		gb.tick()
		gb.tick()
	}
}

func (gb *Gameboy) getImmediate() uint8 {
	val := gb.load(gb.PC)
	if !gb.haltBug {
		gb.PC++
	}
	gb.tick()
	return val
}

// write to the memory controller
func (gb *Gameboy) write(val uint8, adr uint16) {
	switch {
	case adr < 0x8000: // cartridge ROM
		gb.cartridge.Write(val, adr)
	case adr < 0xA000: // VRAM
		gb.video.Write(val, adr)
	case adr < 0xD000: // cartridge RAM
		gb.cartridge.Write(val, adr)
	case adr < 0xE000:
		gb.wram.Write(val, adr)
	case adr < 0xFE00: //ECHO Ram
		gb.wram.Write(val, adr-0x2000)
	case adr < 0xFEA0: // OAM
		gb.video.Write(val, adr)
	case adr < 0xFF00: //OAM corruption bug
		return // TODO implement OAM corruption bug

	// I/O Registers
	case adr == 0xFF00: // input
		gb.input.Write(val, adr)
	case adr < 0xFF03: // serial port
		return // TODO implement serial port
	case adr < 0xFF0F: // timer control
		gb.timer.Write(val, adr)
	case adr == 0xFF0F: // IF flag
		gb.IF = val
	case adr < 0xFF40: // audio + wave RAM
		gb.audio.Write(val, adr)
	case adr == 0xFF4D: //CG
		return // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		gb.video.Write(val, adr)
	case adr == 0xFF70:
		return // TODO [CGB] WRAM bank switch
	case adr >= 0xFF80:
		gb.wram.Write(val, adr)
	case adr == 0xFFFF:
		gb.IE = val
	default:
		panic(fmt.Sprintf("CPU tried to access memory address '%X', but no implementation exists.", adr))
	}

	gb.tick()
}

// load from the memory controller
func (gb *Gameboy) load(adr uint16) uint8 {
	defer gb.tick()

	switch {
	case adr < 0x8000: // cartridge ROM
		return gb.cartridge.Load(adr)
	case adr < 0xA000: // VRAM
		return gb.video.Load(adr)
	case adr < 0xD000: // cartridge RAM
		return gb.cartridge.Load(adr)
	case adr < 0xE000:
		return gb.wram.Load(adr)
	case adr < 0xFE00: //ECHO Ram
		return gb.wram.Load(adr - 0x2000)
	case adr < 0xFEA0: // OAM
		return gb.video.Load(adr)
	case adr < 0xFF00: //OAM corruption bug
		return 0 // TODO implement OAM corruption bug

	// I/O Registers
	case adr == 0xFF00: // input
		return gb.input.Load(adr)
	case adr < 0xFF03: // serial port
		return 1 // TODO implement serial port
	case adr < 0xFF0F: // timer control
		return gb.timer.Load(adr)
	case adr == 0xFF0F: // IF flag
		return gb.IF
	case adr < 0xFF40: // audio + wave RAM
		return gb.audio.Load(adr)
	case adr == 0xFF4D: //CG
		return 1 // TODO: [CGB] KEY1 Prepare Speed Switch
	case adr < 0xFF70: // LCD Control, VRAM stuff and more CGB Flags
		return gb.video.Load(adr)
	case adr == 0xFF70:
		return 1 // TODO [CGB] WRAM bank switch
	case adr >= 0xFF80:
		gb.wram.Load(adr)
	case adr == 0xFFFF:
		return gb.IE
	}

	panic(fmt.Sprintf("CPU tried to access memory address '%X', but no implementation exists.", adr))
}
