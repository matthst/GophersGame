package gameboy

// for the size-2 arrays, the first val is lo, the second is hi
type registers struct {
	PC                 uint16
	AF, BC, DE, HL, SP [2]uint8
}

type getWordDef func(*[2]uint8) uint16

// get the combined register
func getWord(r *[2]uint8) uint16 {
	val := uint16(r[0]) | (uint16(r[1]) << 8)
	return val
}

func getWordInc(r *[2]uint8) uint16 {
	val := getWord(r)
	set_word(r, val+1)
	return val
}

func getWordDec(r *[2]uint8) uint16 {
	val := getWord(r)
	set_word(r, val-1)
	return val
}

func set_word(r *[2]uint8, val uint16) {
	r[0] = uint8(val)
	r[1] = uint8(val >> 8)
}

func (r *registers) getZ() bool {
	return r.AF[1]&0b1000_0000 != 0
}

func (r *registers) setZ(val bool) {
	if val {
		r.AF[1] |= 0b1000_0000
	} else {
		r.AF[1] |= 0b0111_1111
	}
}

func (r *registers) getN() bool {
	return r.AF[1]&0b0100_0000 != 0
}

func (r *registers) setN(val bool) {
	if val {
		r.AF[1] |= 0b0100_0000
	} else {
		r.AF[1] |= 0b1011_1111
	}
}

func (r *registers) getH() bool {
	return r.AF[1]&0b0010_0000 != 0
}

func (r *registers) setH(val bool) {
	if val {
		r.AF[1] |= 0b0010_0000
	} else {
		r.AF[1] |= 0b1101_1111
	}
}

func (r *registers) getC() bool {
	return r.AF[1]&0b0001_0000 != 0
}

func (r *registers) setC(val bool) {
	if val {
		r.AF[1] |= 0b0001_0000
	} else {
		r.AF[1] |= 0b1110_1111
	}
}

func (r *registers) getCarryValue() uint8 {
	if r.getC() {
		return 1
	}
	return 0
}
