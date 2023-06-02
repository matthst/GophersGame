package gameboy

// for the size-2 arrays, the first val is lo, the second is hi
type registers struct {
	SP, PC         uint16
	AF, BC, DE, HL [2]uint8
}

func get_word(r *[2]uint8) uint16 {
	return uint16(r[0]) | (uint16(r[1]) << 8)
}

func set_word(r *[2]uint8, val uint16) {
	r[0] = uint8(val)
	r[1] = uint8(val >> 8)
}

func (r *registers) getZ() bool {
	return r.AF[1]&0b0100_0000 != 0
}

func (r *registers) setZ(val bool) {
	if val {
		r.AF[1] |= 0b0100_0000
	} else {
		r.AF[1] |= 0b1011_1111
	}

}
