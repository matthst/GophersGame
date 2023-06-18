package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matthst/gophersgame/pkg/gameboy"
	"github.com/matthst/gophersgame/pkg/gameboy/video"
	"log"
	"os"
	"time"
)

const RenderScale = 3
const Border = 20
const ViewportWidth = 160
const ViewportHeight = 144

type Game struct {
	op ebiten.DrawImageOptions

	Lightswitch bool
}

func (g *Game) Update() error {
	start := time.Now()
	gameboy.RunOneTick()
	elapsed := time.Since(start)
	log.Printf("Tick took %s", elapsed)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(video.RenderImage, &g.op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return (160 + Border) * RenderScale, (144 + Border) * RenderScale
}

func main() {
	args := os.Args[1:]
	romPath := args[0]
	//romPath := "test_roms/blargg/cpu_instrs/cpu_instrs.gb"
	cartridge, _ := os.ReadFile(romPath)
	gameboy.Bootstrap(cartridge, romPath, nil)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Scale(RenderScale, RenderScale)
	myOp.GeoM.Translate(Border*RenderScale/2, Border*RenderScale/2)
	myOp.Filter = ebiten.FilterNearest

	ebiten.SetWindowSize((160+Border)*RenderScale, (144+Border)*RenderScale)
	ebiten.SetWindowTitle("GophersGame")

	if err := ebiten.RunGame(&Game{op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
