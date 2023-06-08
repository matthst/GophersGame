package gameboy

///////////////////////////////////////
//    GENERAL BYTE<->WORD HELPERS    //
///////////////////////////////////////

// get Word from 2 uint8 values
func getWord(hi, lo uint8) uint16 {
	return uint16(lo) | (uint16(hi) << 8)
}

// get bytes from word, returns hi, lo
func getBytes(word uint16) (uint8, uint8) {
	return uint8(word >> 8), uint8(word)
}

func getWordInc(hi, lo *uint8) uint16 {
	val := getWord(*hi, *lo)
	setBytes(hi, lo, val+1)
	return val
}

func getWordDec(hi, lo *uint8) uint16 {
	val := getWord(*hi, *lo)
	setBytes(hi, lo, val-1)
	return val
}

func setBytes(hi, lo *uint8, val uint16) {
	*lo = uint8(val)
	*hi = uint8(val >> 8)
}

///////////////////////////////////////
//            0xCB HELPERS           //
///////////////////////////////////////

// aluR8Def function definition for the internal helpers used for the 0xCB shift functions
type shiftInternalFuncDef func(uint8) uint8

func (gb *Gameboy) rotateLeftInternal(val uint8) uint8 {
	result := (val << 1) + gb.getCarryValue()
	gb.F = 0b0000_0000 + ((val >> 7) << 4)
	return result
}

func (gb *Gameboy) rotateLeftCircularInternal(val uint8) uint8 {
	gb.F = 0b0000_0000 + ((val >> 7) << 4)
	return (val << 1) + (val >> 7)
}

func (gb *Gameboy) rotateRightInternal(val uint8) uint8 {
	result := (val >> 1) + (gb.getCarryValue() << 7)
	gb.F = 0b0000_0000 + ((val & 0x1) << 4)
	return result
}

func (gb *Gameboy) rotateRightCircularInternal(val uint8) uint8 {
	gb.F = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) + (val << 7)
}

func (gb *Gameboy) shiftLeftInternal(val uint8) uint8 {
	gb.F = 0b0000_0000 + ((val >> 7) << 4)
	return val << 1
}

func (gb *Gameboy) shiftRightInternal(val uint8) uint8 {
	gb.F = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) | (val & 0b1000_0000)
}

func (gb *Gameboy) shiftRightMSBResetInternal(val uint8) uint8 {
	gb.F = 0b0000_0000 + ((val & 0x1) << 4)
	return val >> 1
}

func (gb *Gameboy) swapInternal(val uint8) uint8 {
	gb.F = 0
	gb.setZFlag(val == 0)
	return (val >> 4) + (val << 4)
}

func halfCarryAddCheck8Bit(a, b uint8) bool {
	return (((a & 0xf) + (b & 0xf)) & 0x10) == 0x10
}

func halfCarrySubCheck8Bit(a, b uint8) bool {
	return (((a & 0xf) - (b & 0xf)) & 0x10) == 0x10
}

func halfCarryAddCheck16Bit(a, b uint16) bool {
	return (((a & 0xf00) + (b & 0xf00)) & 0x1000) == 0x1000
}

///////////////////////////////////////
//       MORE REGISTER HELPERS       //
///////////////////////////////////////

func (gb *Gameboy) getAF() uint16 {
	return getWord(gb.A, gb.F)
}

func (gb *Gameboy) getBC() uint16 {
	return gb.getBC()
}

func (gb *Gameboy) getDE() uint16 {
	return gb.getDE()
}

func (gb *Gameboy) getHL() uint16 {
	return getWord(gb.H, gb.L)
}

func (gb *Gameboy) getAndIncSP() uint16 {
	val := gb.SP
	gb.SP++
	return val
}
func (gb *Gameboy) getAndDecSP() uint16 {
	val := gb.SP
	gb.SP--
	return val
}

func (gb *Gameboy) setFlags(Z, N, H, C bool) {
	gb.setZFlag(Z)
	gb.setNFlag(N)
	gb.setHFlag(H)
	gb.setCFlag(C)
}

func (gb *Gameboy) getZFlag() bool {
	return gb.F&0b1000_0000 != 0
}

func (gb *Gameboy) setZFlag(val bool) {
	if val {
		gb.F |= 0b1000_0000
	} else {
		gb.F &= 0b0111_1111
	}
}

func (gb *Gameboy) getNFlag() bool {
	return gb.F&0b0100_0000 != 0
}

func (gb *Gameboy) setNFlag(val bool) {
	if val {
		gb.F |= 0b0100_0000
	} else {
		gb.F &= 0b1011_1111
	}
}

func (gb *Gameboy) getHFlag() bool {
	return gb.F&0b0010_0000 != 0
}

func (gb *Gameboy) setHFlag(val bool) {
	if val {
		gb.F |= 0b0010_0000
	} else {
		gb.F &= 0b1101_1111
	}
}

func (gb *Gameboy) getCFlag() bool {
	return gb.F&0b0001_0000 != 0
}

func (gb *Gameboy) setCFlag(val bool) {
	if val {
		gb.F |= 0b0001_0000
	} else {
		gb.F &= 0b1110_1111
	}
}

func (gb *Gameboy) getCarryValue() uint8 {
	return (gb.F >> 4) % 0b0000_0001
}
