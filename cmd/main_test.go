package main

import (
	"github.com/matthst/gophersgame/pkg/gameboy"
	"os"
	"strings"
	"testing"
)

var (
	BlarggPath = "test_roms/blargg/"
)

func TestBlargg(t *testing.T) {
	testCases := []struct {
		name, file string
		maxTicks   int32
	}{
		{"01-special", "cpu_instrs/01-special.gb", 1000},
		{"02-interrupts", "cpu_instrs/02-interrupts.gb", 1000},
		{"03-op sp,hl", "cpu_instrs/03-op sp,hl.gb", 1000},
		{"04-op r,imm", "cpu_instrs/04-op r,imm.gb", 1000},
		{"05-op rp", "cpu_instrs/05-op rp.gb", 1000},
		{"06-ld r,r", "cpu_instrs/06-ld r,r.gb", 1000},
		{"07-jr,jp,call,ret,rst", "cpu_instrs/07-jr,jp,call,ret,rst.gb", 1000},
		{"08-misc instrs", "cpu_instrs/08-misc instrs.gb", 1000},
		{"09-op r,r", "cpu_instrs/09-op r,r.gb", 1000},
		{"10-bit ops", "cpu_instrs/10-bit ops.gb", 5000},
		{"11-op a,(hl)", "cpu_instrs/11-op a,(hl).gb", 5000},
		{"halt_bug", "halt_bug.gb", 5000},
		//{"instr_timing", "instr_timing/instr_timing.gb", 5000},
		//{"interrupt_time", "interrupt_time/interrupt_time.gb", 5000},
		//{"01-read_timing", "mem_timing/individual/01-read_timing.gb", 5000},
		//{"02-write_timing", "mem_timing/individual/02-write_timing.gb", 5000},
		//{"03-modify_timing", "mem_timing/individual/03-modify_timing.gb", 5000}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, fileErr := os.ReadFile(BlarggPath + tc.file)
			if fileErr != nil {
				t.Error(fileErr)
				t.FailNow()
			}

			gameboy.Bootstrap(file)
			sBuilder := strings.Builder{}
			gameboy.SerialC.StringBuffer = &sBuilder

			for tickCount := int32(0); ; tickCount++ {
				gameboy.RunOneTick()
				serialResult := sBuilder.String()
				if strings.Contains(serialResult, "Passed") {
					break
				} else if strings.Contains(serialResult, "Failed") {
					t.Errorf(serialResult)
					break
				}
				if tickCount > tc.maxTicks {
					t.Errorf("Test exceeded %d ticks and timed out.", tc.maxTicks)
					break
				}
			}
		})
	}
}
