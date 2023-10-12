package game

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	logger  *slog.Logger
	systems []systems.System
	root    *entity.Entity
}

func New(logger *slog.Logger, root *entity.Entity, systems []systems.System) *Game {
	return &Game{
		logger:  logger,
		systems: systems,
		root:    root,
	}
}

func (g *Game) Update() error {
	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return fmt.Errorf("error updating system %T: %w", s, err)
		}
	}

	entities := List(g.root)

	if err := Update(entities); err != nil {
		return fmt.Errorf("error updating entities: %w", err)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	entities := List(g.root)

	if err := Draw(entities, screen); err != nil {
		g.logger.Error(err.Error())
	}

	if err := Debug(entities, screen); err != nil {
		g.logger.Error(err.Error())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func Update(entities []*entity.Entity) error {
	for _, e := range entities {
		for _, c := range e.Components() {
			if updater, ok := c.(components.Updater); ok {
				if err := updater.Update(); err != nil {
					return fmt.Errorf("error updating entity: %w", err)
				}
			}
		}
	}

	return nil
}

func Draw(entities []*entity.Entity, screen *ebiten.Image) error {
	for _, e := range entities {
		for _, c := range e.Components() {
			if drawer, ok := c.(components.Drawer); ok {
				if err := drawer.Draw(screen); err != nil {
					return fmt.Errorf("error drawing entity: %w", err)
				}
			}
		}
	}

	return nil
}

func Debug(entities []*entity.Entity, screen *ebiten.Image) error {
	for _, e := range entities {
		for _, c := range e.Components() {
			if drawer, ok := c.(components.Debugger); ok {
				if err := drawer.Debug(screen); err != nil {
					return fmt.Errorf("error debugging entity: %w", err)
				}
			}
		}
	}

	return nil
}

// List returns a list of all entities in the tree.
func List(e *entity.Entity) []*entity.Entity {
	list := []*entity.Entity{e}

	for _, c := range e.Children() {
		list = append(list, List(c)...)
	}

	return list
}
