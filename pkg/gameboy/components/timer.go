package components

type Timer struct {
	dividerRegister, timerCounter, timerModulo, timerControl uint8
	divIndex, timaIndex, TimaClock                           int
	timerEnable, timerOverflow                               bool
}

var (
	clockSelect = [4]int{256, 4, 16, 64}
)

func (t *Timer) Load(adr uint16) uint8 {
	switch adr {
	case 0xFF04:
		return t.dividerRegister
	case 0xFF05:
		return t.timerCounter
	case 0xFF06:
		return t.timerModulo
	case 0xFF07:
		return t.timerControl
	}
	return 1
}

func (t *Timer) Write(val uint8, adr uint16) {
	switch adr {
	case 0xFF04:
		t.dividerRegister = 0 // TODO find out if writing to TIMA does anything
	case 0xFF06:
		t.timerModulo = val
	case 0xFF07:
		t.timerControl = val & 0b111
		t.updateTAC()
	}
}

func (t *Timer) Cycle() uint8 {
	interruptResult := uint8(0)

	t.divIndex++
	t.timaIndex++
	if t.divIndex == 64 {
		t.divIndex = 0
		t.dividerRegister++
	}

	if t.timerOverflow {
		t.timerOverflow = false
		t.timerCounter = t.timerModulo
		interruptResult = 0b100
	}

	if t.timerEnable && t.timaIndex > t.TimaClock {
		t.timaIndex = 0
		t.timerCounter++

		if t.timerCounter == 0 { //TIMA overflow
			t.timerOverflow = true
		}
	}
	return interruptResult
}

func (t *Timer) updateTAC() {
	t.timerEnable = t.timerControl&0b100 == 1
	t.TimaClock = clockSelect[t.timerControl&0b11]
}
