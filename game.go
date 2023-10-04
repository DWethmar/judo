package main

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	logger *slog.Logger
	root   *entity.Entity
}

func (g *Game) Update() error {
	if err := RecursiveInit(g.root); err != nil {
		return fmt.Errorf("error initializing root entity: %w", err)
	}

	return RecursiveUpdate(g.root)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := RecursiveDraw(g.root, screen); err != nil {
		g.logger.Error(err.Error())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func RecursiveInit(e *entity.Entity) error {
	if err := e.Init(); err != nil {
		return fmt.Errorf("error initializing entity: %w", err)
	}

	for i, c := range e.Children() {
		if err := RecursiveInit(c); err != nil {
			return fmt.Errorf("error initializing child entity at index %d: %w", i, err)
		}
	}

	return nil
}

func RecursiveUpdate(e *entity.Entity) error {
	if err := e.Update(); err != nil {
		return fmt.Errorf("error updating entity: %w", err)
	}

	for i, c := range e.Children() {
		if err := RecursiveUpdate(c); err != nil {
			return fmt.Errorf("error updating child entity at index %d: %w", i, err)
		}
	}

	return nil
}

func RecursiveDraw(e *entity.Entity, screen *ebiten.Image) error {
	if err := e.Draw(screen); err != nil {
		return fmt.Errorf("error drawing entity: %w", err)
	}

	for i, c := range e.Children() {
		if err := RecursiveDraw(c, screen); err != nil {
			return fmt.Errorf("error drawing child entity at index %d: %w", i, err)
		}
	}

	return nil
}
