package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/components"
)

type Gameboy struct {
	Regs      *registers
	video     *video
	audio     *audio
	Cartridge components.Cart
	input     components.Input
	WRAM      components.WRAM
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
		gb.Cartridge = components.RomOnly{Rom: file}
	default:
		panic(fmt.Sprintf("Opcode '%X' not implemented", file[0x0148]))
	}

	return gb
}

// write to the memory bank
func (gb *Gameboy) write(val uint8, adr uint16) {
	//IO Registers

	//RAM and

	switch {
	case adr < 0x8000: // Cartridge ROM
		gb.Cartridge.Write(val, adr)
	case adr < 0xA000: // VRAM
		gb.video.write(val, adr)
	case adr < 0xC000: // Cartridge RAM
		gb.Cartridge.Write(val, adr)
	case adr < 0xD000:
	case adr < 0xE000:
	case adr < 0xFE00:
	case adr < 0xD000:
	case adr < 0xD000:
	}

}

// load from the memory bank
func (gb *Gameboy) load(adr uint16) uint8 {
	return 0 // TODO
}
