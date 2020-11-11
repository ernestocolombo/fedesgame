package main

import (
	"fmt"
	"log"
	"time"

	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img *ebiten.Image
	W   int
	H   int
	X   float64
	Y   float64
}

type Asteroid struct {
	Sprite
	Rot      float64
	RotSpeed float64
	Speed    float64
}

type Ship struct {
	Sprite
	Speed float64
}

type Game struct {
	BgImg                *ebiten.Image
	W                    int
	H                    int
	AsteroidInitialSpeed float64
	Asteroid             *Asteroid
	Ship                 *Ship
	Level                int64
	FlyBys               int64
	Score                int64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Ship.X -= g.Ship.Speed
		if g.Ship.X < 0 {
			g.Ship.X = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Ship.X += g.Ship.Speed
		if g.Ship.X > float64(g.W-g.Ship.W) {
			g.Ship.X = float64(g.W - g.Ship.W)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Ship.Y -= g.Ship.Speed
		if g.Ship.Y < 0 {
			g.Ship.Y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Ship.Y += g.Ship.Speed
		if g.Ship.Y > float64(g.H-g.Ship.H) {
			g.Ship.Y = float64(g.H - g.Ship.H)
		}
	}

	if g.Asteroid.Y > float64(g.H) {
		g.Asteroid.X = float64(rand.Intn(g.W - g.Asteroid.W))
		g.Asteroid.Y = float64(-1 * g.Asteroid.H)
		g.FlyBys++
		g.Score += g.Level * 10
		if g.FlyBys >= 10 {
			g.Level++
			g.Asteroid.Speed = float64(g.Level+1) * .5 * g.AsteroidInitialSpeed
			g.FlyBys = 0
		}
	} else {
		g.Asteroid.Y += g.Asteroid.Speed
		// log.Printf("X: %f Y: %f\n", g.Asteroid.X, g.Asteroid.Y)
		g.Asteroid.Rot += g.Asteroid.RotSpeed
	}

	if g.GameOver() {
		return fmt.Errorf("Game Over! Your score is %d", g.Score)
	}

	return nil
}

func (g *Game) GameOver() bool {
	shipX2 := g.Ship.X + float64(g.Ship.W)
	shipY2 := g.Ship.Y + float64(g.Ship.H)
	asteroidX2 := g.Asteroid.X + float64(g.Asteroid.W)
	asteroidY2 := g.Asteroid.Y + float64(g.Asteroid.H)
	overX1 := g.Ship.X >= g.Asteroid.X && g.Ship.X <= asteroidX2
	overX2 := shipX2 >= g.Asteroid.X && shipX2 <= asteroidX2
	overY1 := g.Ship.Y >= g.Asteroid.Y && g.Ship.Y <= asteroidY2
	overY2 := shipY2 >= g.Asteroid.Y && shipY2 <= asteroidY2
	log.Printf("%f %f %f %f", g.Ship.X, shipX2, g.Asteroid.X, asteroidX2)
	return (overX1 || overX2) && (overY1 || overY2)
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(
		float64(g.W)/float64(g.BgImg.Bounds().Dx()),
		float64(g.H)/float64(g.BgImg.Bounds().Dy()),
	)
	screen.DrawImage(g.BgImg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Ship.X, g.Ship.Y)
	screen.DrawImage(g.Ship.Img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(-g.Asteroid.W)/2, float64(-g.Asteroid.H)/2)
	op.GeoM.Rotate(g.Asteroid.Rot)
	op.GeoM.Translate(g.Asteroid.X, g.Asteroid.Y)
	screen.DrawImage(g.Asteroid.Img, op)

	// Debug infos
	debugMsg := fmt.Sprintf("Your score is %d at level %d", g.Score, g.Level)
	ebitenutil.DebugPrint(screen, debugMsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// log.Printf("Layout %d %d", outsideWidth, outsideHeight)
	g.W = outsideWidth
	g.H = outsideHeight
	return g.W, g.H
}

func NewGame(screenWidth int, screenHeight int) *Game {
	rand.Seed(time.Now().UnixNano())

	bgImg, _, err := ebitenutil.NewImageFromFile("res/background.png")
	if err != nil {
		log.Fatalf("could not load background image: %v", err)
	}

	shipImg, _, err := ebitenutil.NewImageFromFile("res/millennium.png")
	if err != nil {
		log.Fatalf("could not load millennium image: %v", err)
	}
	shipW := shipImg.Bounds().Dx()
	shipH := shipImg.Bounds().Dy()
	ship := &Ship{
		Sprite: Sprite{
			Img: shipImg,
			W:   shipW,
			H:   shipH,
			X:   float64(screenWidth-shipW) / 2,
			Y:   float64(screenHeight-shipH) / 2,
		},
		Speed: 6,
	}

	origAsteroidImg, _, err := ebitenutil.NewImageFromFile("res/asteroid.png")
	if err != nil {
		log.Fatalf("could not load asteroid image: %v", err)
	}
	opAsteroid := &ebiten.DrawImageOptions{}
	var asteroidScale float64 = 0.5
	opAsteroid.GeoM.Scale(asteroidScale, asteroidScale)
	asteroidW := int(float64(origAsteroidImg.Bounds().Dx()) * asteroidScale)
	asteroidH := int(float64(origAsteroidImg.Bounds().Dy()) * asteroidScale)
	asteroidImg := ebiten.NewImage(asteroidW, asteroidH)
	asteroidImg.DrawImage(origAsteroidImg, opAsteroid)
	asteroid := &Asteroid{
		Sprite: Sprite{
			Img: asteroidImg,
			W:   asteroidW,
			H:   asteroidH,
			X:   float64(rand.Intn(screenWidth - asteroidW)),
			Y:   -1 * float64(asteroidH),
		},
		Rot:      0,
		RotSpeed: .1,
		Speed:    1,
	}

	return &Game{
		BgImg:                bgImg,
		Asteroid:             asteroid,
		AsteroidInitialSpeed: 6,
		Ship:                 ship,
		FlyBys:               0,
		Level:                1,
		Score:                0,
		H:                    screenHeight,
		W:                    screenWidth,
	}
}

func main() {
	title := "This my game and it's called Asteroid Field Ship!"
	screenWidth := 1000
	screenHeight := 800
	// ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowDecorated(true)
	if err := ebiten.RunGame(NewGame(screenWidth, screenHeight)); err != nil {
		log.Fatal(err)
	}
}
