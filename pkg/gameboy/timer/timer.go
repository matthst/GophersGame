package Timer

var (
	clockByte = [4]uint16{0x200, 0x008, 0x020, 0x080}

	DividerClk, nextDividerClk uint16

	timerFreqSel, tima, tma, tac uint8
	timerEnable, timerOverflow   bool
)

func Load(adr uint16) uint8 {
	switch adr {
	case 0xFF04:
		return uint8(DividerClk >> 8)
	case 0xFF05:
		return tima
	case 0xFF06:
		return tma
	case 0xFF07:
		return 0b1111_1000 | tac
	}
	return 1
}

func Write(val uint8, adr uint16) {
	switch adr {
	case 0xFF04:
		nextDividerClk = 0
	case 0xFF05:
		tima = val
		timerOverflow = false
	case 0xFF06:
		tma = val
	case 0xFF07:
		tac = val
		timerEnable = tac&0b100 != 0
		timerFreqSel = tac & 0b11
	}
}

func Cycle() uint8 {
	returnVal := uint8(0)
	if timerOverflow {
		timerOverflow = false
		tima = tma
		returnVal = 0b100
	}

	if timerEnable && nextDividerClk&clockByte[timerFreqSel] == 0 && (DividerClk^nextDividerClk)&clockByte[timerFreqSel] != 0 {
		tima++
		if tima == 0 {
			timerOverflow = true // Overflow handling is done one cycle after
		}
	}

	DividerClk = nextDividerClk
	nextDividerClk += 4
	return returnVal
}
