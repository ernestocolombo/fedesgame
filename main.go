package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1000
	screenHeight = 800
)

var bgImg *ebiten.Image

var millenniumImg *ebiten.Image
var millenniumX float64
var millenniumY float64
var millenniumW float64
var millenniumH float64
var millenniumSpeed float64 = 6

func init() {
	var err error
	millenniumImg, _, err = ebitenutil.NewImageFromFile("res/millennium.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load millennium image: %v", err)
	}
	millenniumW = float64(millenniumImg.Bounds().Dx())
	millenniumH = float64(millenniumImg.Bounds().Dy())
	millenniumX = (screenWidth - millenniumW) / 2
	millenniumY = (screenHeight - millenniumH) / 2

	bgImg, _, err = ebitenutil.NewImageFromFile("res/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load background image: %v", err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Get user input and update game state
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		millenniumX -= millenniumSpeed
		if millenniumX < 0 {
			millenniumX = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		millenniumX += millenniumSpeed
		if millenniumX > screenWidth-millenniumW {
			millenniumX = screenWidth - millenniumW
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		millenniumY -= millenniumSpeed
		if millenniumY < 0 {
			millenniumY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		millenniumY += millenniumSpeed
		if millenniumY > screenHeight-millenniumH {
			millenniumY = screenHeight - millenniumH
		}
	}

	// Draw scene
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(screenWidth/float64(bgImg.Bounds().Dx()), screenHeight/float64(bgImg.Bounds().Dy()))
	screen.DrawImage(bgImg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(millenniumX, millenniumY)
	screen.DrawImage(millenniumImg, op)

	// Debug infos
	debugMsg := fmt.Sprintf("Millennium X: %d, Y: %d", int64(millenniumX), int64(millenniumY))
	ebitenutil.DebugPrint(screen, debugMsg)

	return nil
}

func main() {
	title := "This my game and it's called Asteroid Field!"
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, title); err != nil {
		log.Fatal(err)
	}
}
