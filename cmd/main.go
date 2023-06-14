package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matthst/gophersgame/pkg/gameboy"
	"log"
	"os"
)

const RenderScale = 1
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
	screen.DrawImage(gameboy.Vid.RenderImage, &g.op)
}

func updateTile(eTile *ebiten.Image, data [16]uint8) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	cartridge, _ := os.ReadFile("testing/blargh/cpu_instrs/individual/03-op sp,hl.gb")
	gameboy.Bootstrap(cartridge)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Scale(RenderScale, RenderScale)
	myOp.GeoM.Translate(10, 10)
	myOp.Filter = ebiten.FilterLinear

	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("GophersGame")
	if err := ebiten.RunGame(&Game{op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
