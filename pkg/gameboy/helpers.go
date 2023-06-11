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

func rotateLeftInternal(val uint8) uint8 {
	result := (val << 1) + getCarryValue()
	fReg = 0b0000_0000 + ((val >> 7) << 4)
	return result
}

func rotateLeftCircularInternal(val uint8) uint8 {
	fReg = 0b0000_0000 + ((val >> 7) << 4)
	return (val << 1) + (val >> 7)
}

func rotateRightInternal(val uint8) uint8 {
	result := (val >> 1) + (getCarryValue() << 7)
	fReg = 0b0000_0000 + ((val & 0x1) << 4)
	return result
}

func rotateRightCircularInternal(val uint8) uint8 {
	fReg = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) + (val << 7)
}

func shiftLeftInternal(val uint8) uint8 {
	fReg = 0b0000_0000 + ((val >> 7) << 4)
	return val << 1
}

func shiftRightInternal(val uint8) uint8 {
	fReg = 0b0000_0000 + ((val & 0x1) << 4)
	return (val >> 1) | (val & 0b1000_0000)
}

func shiftRightMSBResetInternal(val uint8) uint8 {
	fReg = 0b0000_0000 + ((val & 0x1) << 4)
	return val >> 1
}

func swapInternal(val uint8) uint8 {
	fReg = 0
	setZFlag(val == 0)
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

func getHL() uint16 {
	return getWord(hReg, lReg)
}

func getAndIncSP() uint16 {
	val := SP
	SP++
	return val
}
func getAndDecSP() uint16 {
	val := SP
	SP--
	return val
}

func setFlags(Z, N, H, C bool) {
	setZFlag(Z)
	setNFlag(N)
	setHFlag(H)
	setCFlag(C)
}

func getZFlag() bool {
	return fReg&0b1000_0000 != 0
}

func setZFlag(val bool) {
	if val {
		fReg |= 0b1000_0000
	} else {
		fReg &= 0b0111_1111
	}
}

func getNFlag() bool {
	return fReg&0b0100_0000 != 0
}

func setNFlag(val bool) {
	if val {
		fReg |= 0b0100_0000
	} else {
		fReg &= 0b1011_1111
	}
}

func getHFlag() bool {
	return fReg&0b0010_0000 != 0
}

func setHFlag(val bool) {
	if val {
		fReg |= 0b0010_0000
	} else {
		fReg &= 0b1101_1111
	}
}

func getCFlag() bool {
	return fReg&0b0001_0000 != 0
}

func setCFlag(val bool) {
	if val {
		fReg |= 0b0001_0000
	} else {
		fReg &= 0b1110_1111
	}
}

func getCarryValue() uint8 {
	return (fReg >> 4) % 0b0000_0001
}
