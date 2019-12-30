package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var millenniumImg *ebiten.Image
var millenniumX float64 = 365
var millenniumSpeed float64 = 6

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
	op.GeoM.Translate(millenniumX, 252)
	screen.DrawImage(millenniumImg, op)

	millenniumX = millenniumX + millenniumSpeed

	if millenniumX >= 730 || millenniumX <= 0 {
		millenniumSpeed = millenniumSpeed * -1
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 800, 600, 1, "This my game and it's called Asteroid Field!"); err != nil {
		log.Fatal(err)
	}
}
