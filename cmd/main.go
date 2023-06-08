package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"time"
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
	var myColor color.Color
	if g.Lightswitch {
		myColor = color.Gray{150}
	} else {
		myColor = color.Gray{160}
	}
	start := time.Now()
	for x := 0; x < ViewportHeight; x++ {
		for y := 0; y < ViewportWidth; y++ {
			g.renderOutput.Set(x, y, myColor)
		}
	}

	g.Lightswitch = !g.Lightswitch

	elapsed := time.Since(start)
	log.Printf("Drawing took %s", elapsed)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(g.renderOutput, &g.op)

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
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{renderOutput: renderOutputImage, op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
