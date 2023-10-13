package game

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/systems"
	"github.com/dwethmar/judo/systems/scaling"
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
			return fmt.Errorf("error updating system: %w", err)
		}
	}

	entities := append([]*entity.Entity{g.root}, List(g.root)...)

	if err := Update(entities); err != nil {
		return fmt.Errorf("error updating entities: %w", err)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scale := 0.0
	if s := scaling.FromList(g.systems); s != nil {
		scale = s.Scale
	}

	entities := append([]*entity.Entity{g.root}, List(g.root)...)

	offScreen := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

	if err := Draw(entities, offScreen); err != nil {
		g.logger.Error(err.Error())
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(scale, scale)

	screenX := screen.Bounds().Dx()
	screenY := screen.Bounds().Dy()

	offScreenX := offScreen.Bounds().Dx()
	offScreenY := offScreen.Bounds().Dy()

	// Calculate the offsets to center the image
	offsetX := (float64(screenX) - float64(offScreenX)*scale) / 2.0
	offsetY := (float64(screenY) - float64(offScreenY)*scale) / 2.0

	// Apply the offsets
	opt.GeoM.Translate(offsetX, offsetY)

	screen.DrawImage(offScreen, opt)

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
