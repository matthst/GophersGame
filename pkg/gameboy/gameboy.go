package gameboy

import (
	"fmt"
	"github.com/matthst/gophersgame/pkg/gameboy/cartridge"
	_ "github.com/matthst/gophersgame/pkg/gameboy/cartridge"
)

type Gameboy struct {
	Regs      *registers
	Cartridge cartridge.Cart
}

func bootstrap(file []uint8) Gameboy {

	gb := Gameboy{}

	switch file[0x0148] {
	case 0x00:
		gb.Cartridge = cartridge.RomOnly{Rom: file}
	default:
		panic(fmt.Sprintf("Opcode '%X' not implemented", file[0x0148]))
	}

	return gb
}

// write to the memory bank
func (gb *Gameboy) write(val uint8, adr uint16) {

}

func (gb *Gameboy) load(adr uint16) uint8 {
	return 0 // TODO
}
