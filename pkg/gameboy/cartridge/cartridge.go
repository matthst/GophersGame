package Cartridge

import (
	"fmt"
	"os"
	"path/filepath"
)

type Cartridge interface {
	Write(val uint8, adr uint16)
	Load(adr uint16) uint8
}

// getROMSize returns the number of ROM banks based on the header byte 0x0148
func getROMSize(val uint8) uint16 {
	return 32 * (1 << val)
}

// getRAMSize returns the number of RAM bank based on the header byte 0x0147
func getRAMSize(val uint8) uint16 {
	switch val {
	case 0x02:
		return 1
	case 0x03:
		return 4
	case 0x04:
		return 16
	case 0x05:
		return 8
	default:
		return 0
	}
}

func createOrLoadSaveFile(romPath string, ramSize int) (string, []uint8) {
	savePath := romPath[0:len(romPath)-len(filepath.Ext(romPath))] + ".sav"
	ram := make([]uint8, ramSize)

	if _, err := os.Stat(savePath); err == nil {
		saveFile, _ := os.OpenFile(savePath, os.O_RDONLY, 0644)
		_, err := saveFile.Read(ram)
		if err != nil {
			println(err)
			panic("Could not read from save file.")
		}
	} else {
		saveFile, _ := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0644)
		_, err := saveFile.Write(ram)
		if err != nil {
			println(err)
			panic("Could not write to save file.")
		}
		fmt.Printf("File does not exist\n")
	}
	return savePath, ram
}

func writeToSaveFile(savePath string, ram []uint8) {
	saveFile, _ := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	_, _ = saveFile.Write(ram)
}
