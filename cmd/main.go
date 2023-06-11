package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const RenderScale = 1
const ViewportWidth = 160
const ViewportHeight = 144

type Game struct {
	renderOutput *ebiten.Image
	op           ebiten.DrawImageOptions

	Lightswitch bool
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func updateTile(eTile *ebiten.Image, data [16]uint8) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	renderOutputImage := ebiten.NewImage(ViewportWidth, ViewportHeight)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Scale(RenderScale, RenderScale)
	myOp.GeoM.Translate(64, 64)
	myOp.Filter = ebiten.FilterNearest

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GophersGame")
	if err := ebiten.RunGame(&Game{renderOutput: renderOutputImage, op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
