package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"time"
)

const RENDER_SCALE = 1
const VIEWPORT_WIDTH = 160
const VIEWPORT_HEIGHT = 144

type Game struct {
	renderOutput *ebiten.Image
	op           ebiten.DrawImageOptions

	Lightswitch bool
}

func (g *Game) Update() error {
	var myColor color.Color
	if g.Lightswitch {
		myColor = color.White
	} else {
		myColor = color.Black
	}
	start := time.Now()
	for x := 0; x < VIEWPORT_HEIGHT; x++ {
		for y := 0; y < VIEWPORT_WIDTH; y++ {
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
	renderOutputImage := ebiten.NewImage(VIEWPORT_WIDTH, VIEWPORT_HEIGHT)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Scale(RENDER_SCALE, RENDER_SCALE)
	myOp.GeoM.Translate(64, 64)
	myOp.Filter = ebiten.FilterNearest

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{renderOutput: renderOutputImage, op: *myOp}); err != nil {
		log.Fatal(err)
	}
}
