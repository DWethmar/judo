package main

import (
	"log/slog"

	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	logger   *slog.Logger
	entities []*entity.Entity
}

func (g *Game) Update() error {
	for _, e := range g.entities {
		if err := e.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range g.entities {
		if err := e.Draw(screen); err != nil {
			panic(err)
		}
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
