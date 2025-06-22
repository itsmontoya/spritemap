package main

// Credits for the test sprite:
// https://kenney.nl/assets/pixel-platformer

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/itsmontoya/spritemap"
)

func main() {
	var (
		g   testGame
		err error
	)

	if g.s, err = spritemap.NewFromFile("./sprite.png", 18, 1); err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(128, 128)
	ebiten.SetWindowTitle("SpriteMap")
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

type testGame struct {
	s *spritemap.Spritemap
}

func (g *testGame) Update() error { return nil }

func (g *testGame) Draw(screen *ebiten.Image) {
	img, err := g.s.GetByRowAndColumn(8, 1)
	if err != nil {
		panic(err)
	}

	screen.DrawImage(img, nil)
}

func (g *testGame) Layout(outsideWidth, outsideHeight int) (int, int) { return 128, 128 }
