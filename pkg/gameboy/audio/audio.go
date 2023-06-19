package audio

type Audio struct {
}

var (
	NR50, NR51, NR52 uint8

	apuEnable bool
)

func (a *Audio) Load(adr uint16) uint8 {
	switch adr {
	case 0xFF24: // TODO implement NR50, Master Volume & VIN panning
		return NR50
	case 0xFF25: // TODO implement NR51, sound panning
		return NR51
	case 0xFF26: // TODO implement NR52, global sound on/ff
		return NR52
	}

	return 0x00
}

func (a *Audio) Write(val uint8, adr uint16) {
	switch adr {
	case 0xFF24:
		NR50 = val
	case 0xFF25:
		NR51 = val
	case 0xFF26:
		NR52 = val
	}
}
