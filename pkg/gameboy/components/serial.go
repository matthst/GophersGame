package components

import (
	"strings"
)

type Serial struct {
	clockCounter int

	serialTransferBuffer, serialTransferControl, incomingValue, outgoingValue, bitCounter uint8
	transferActive, ClockSpeed, ShiftClock                                                bool

	StringBuilder *strings.Builder
}

var (
	clockModulo   = 128         // TODO [CGB] add clockspeed change
	incomingValue = uint8(0xFF) // TODO [MISC] actually inject data
)

func (s *Serial) Load(adr uint16) uint8 {
	if adr == 0xFF01 {
		return s.serialTransferBuffer
	} else if adr == 0xFF02 {
		return s.serialTransferControl
	}
	return 0xFF
}

func (s *Serial) Write(val uint8, adr uint16) {
	if adr == 0xFF01 {
		s.serialTransferBuffer = val
	} else if adr == 0xFF02 {
		s.serialTransferControl = val & 0b1000_0001 // TODO [CGB] add clock speed change
		s.updateFlags()
	}
}

func (s *Serial) Cycle() uint8 {
	if s.transferActive {
		s.clockCounter++
		if s.clockCounter == clockModulo {
			s.clockCounter = 0
			s.serialTransferBuffer = (s.serialTransferBuffer << 1) | ((incomingValue >> s.bitCounter) & 0b1)
			s.bitCounter++
			if s.bitCounter == 8 {
				s.bitCounter = 0
				s.transferActive = false
				s.serialTransferControl &= 0b1
				if s.StringBuilder != nil {
					s.StringBuilder.WriteByte(s.outgoingValue)
				} else {
					print(string(s.outgoingValue))
				}
				return 0b1000
			}
		}
	}
	return 0x00
}

func (s *Serial) updateFlags() {
	newTransferFlag := s.serialTransferControl&0x80 != 0
	if !s.transferActive && newTransferFlag {
		s.transferActive = true
		s.clockCounter = 0
		s.bitCounter = 0
		s.outgoingValue = s.serialTransferBuffer
	}
	if s.transferActive && !newTransferFlag {
		s.transferActive = false
	}
}
