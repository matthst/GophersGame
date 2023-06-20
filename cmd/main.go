package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/matthst/gophersgame/pkg/gameboy"
	"github.com/matthst/gophersgame/pkg/gameboy/video"
	"log"
	"os"
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
	gameboy.RunOneTick()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(video.CurrentFrame, &g.op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return (160 + Border) * RenderScale, (144 + Border) * RenderScale
}

func main() {
	args := os.Args[1:]
	romPath := args[0]
	//romPath = "test_roms/mooneye/acceptance/timer/tim00_div_trigger.gb"
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
