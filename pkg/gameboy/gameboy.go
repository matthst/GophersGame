package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/components"
)

type Gameboy struct {
	IME    bool
	IE, IF uint8

	Regs *registers

	video     components.Video
	audio     components.Audio
	cartridge components.Cartridge
	wram      components.WRAM
	input     components.Input
	timer     components.Timer
}

func bootstrap(file []uint8) Gameboy {

	gb := Gameboy{}

	gb.Regs = &registers{
		PC: 0x0100, SP: 0xFFFE,
		AF: [2]uint8{0x00, 0x01},
		BC: [2]uint8{0x13, 0x00},
		DE: [2]uint8{0xD8, 0x00},
		HL: [2]uint8{0x4D, 0x01}}

	switch file[0x0147] {
	case 0x00:
		gb.cartridge = components.RomOnly{Rom: file}
	default:
		panic(fmt.Sprintf("Opcode '%X' not implemented", file[0x0148]))
	}

	return gb
}

// write to the memory bank
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
		return // TODO Write to IF interrupt flag
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
		// TODO find out what to do here
	}

	panic(fmt.Sprintf("CPU tried to access memory address '%X', but no implementation exists.", adr))
}

// load from the memory bank
func (gb *Gameboy) load(adr uint16) uint8 {
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
		return 1 // TODO Write to IF interrupt flag
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
		return 1 // TODO find out what to do here
	}

	panic(fmt.Sprintf("CPU tried to access memory address '%X', but no implementation exists.", adr))
}
