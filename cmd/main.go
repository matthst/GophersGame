package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matthst/gophersgame/pkg/gameboy"
	"log"
	"os"
)

const RenderScale = 2
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
	return 160 * RenderScale, 144 * RenderScale
}

func main() {
	cartridge, _ := os.ReadFile("testing/blargg/cpu_instrs/individual/04-op r,imm.gb")
	gameboy.Bootstrap(cartridge)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Scale(RenderScale, RenderScale)
	myOp.GeoM.Translate(0, 0)
	myOp.Filter = ebiten.FilterNearest

	ebiten.SetWindowSize(160*RenderScale, 144*RenderScale)
	ebiten.SetWindowTitle("GophersGame")
	if err := ebiten.RunGame(&Game{op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
