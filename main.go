package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var millenniumImg *ebiten.Image

func init() {
	var err error
	millenniumImg, _, err = ebitenutil.NewImageFromFile("res/millennium.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load millennium image: %v", err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(365, 252)
	screen.DrawImage(millenniumImg, op)
	return nil
}

func main() {
	if err := ebiten.Run(update, 800, 600, 1, "This my game and it's called Asteroid Field!"); err != nil {
		log.Fatal(err)
	}
}
