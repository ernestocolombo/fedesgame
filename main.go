package main

import (
	"fmt"
	"log"
	"time"

	"math/rand"

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

var asteroidImg *ebiten.Image
var asteroidX float64
var asteroidY float64
var asteroidW float64
var asteroidH float64
var asteroidScale float64 = .5
var asteroidSpeed float64 = 10

func init() {
	var err error
	rand.Seed(time.Now().UnixNano())

	bgImg, _, err = ebitenutil.NewImageFromFile("res/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load background image: %v", err)
	}

	millenniumImg, _, err = ebitenutil.NewImageFromFile("res/millennium.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load millennium image: %v", err)
	}
	millenniumW = float64(millenniumImg.Bounds().Dx())
	millenniumH = float64(millenniumImg.Bounds().Dy())
	millenniumX = (screenWidth - millenniumW) / 2
	millenniumY = (screenHeight - millenniumH) / 2

	asteroidImg, _, err = ebitenutil.NewImageFromFile("res/asteroid.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not load asteroid image: %v", err)
	}
	asteroidW = float64(asteroidImg.Bounds().Dx()) * asteroidScale
	asteroidH = float64(asteroidImg.Bounds().Dy()) * asteroidScale
	asteroidX = float64(rand.Int63n(screenWidth - int64(asteroidW)))
	asteroidY = -1 * asteroidH
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

	if asteroidY > screenHeight {
		asteroidX = float64(rand.Int63n(screenWidth - int64(asteroidW)))
		asteroidY = -1 * asteroidH
	} else {
		asteroidY += asteroidSpeed
	}

	// Draw scene
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(screenWidth/float64(bgImg.Bounds().Dx()), screenHeight/float64(bgImg.Bounds().Dy()))
	screen.DrawImage(bgImg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(millenniumX, millenniumY)
	screen.DrawImage(millenniumImg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(asteroidScale, asteroidScale)
	op.GeoM.Translate(asteroidX, asteroidY)
	screen.DrawImage(asteroidImg, op)

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
