package gameboy

// get the combined register
func getWord(r *[2]uint8) uint16 {
	val := uint16(r[0]) | (uint16(r[1]) << 8)
	return val
}

func setWord(r *[2]uint8, val uint16) {
	r[0] = uint8(val)
	r[1] = uint8(val >> 8)
}

func getWordInc(r *[2]uint8) uint16 {
	val := getWord(r)
	setWord(r, val+1)
	return val
}

func getWordDec(r *[2]uint8) uint16 {
	val := getWord(r)
	setWord(r, val-1)
	return val
}
