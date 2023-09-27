package entity

import (
	"sort"

	"github.com/dwethmar/judo/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	X, Y       int64
	components []components.Component
}

func (e *Entity) AddComponent(c components.Component) {
	e.components = append(e.components, c)
	sort.Slice(e.components, func(i, j int) bool {
		return e.components[i].Priority() < e.components[j].Priority()
	})
}

func (e *Entity) Draw(screen *ebiten.Image) error {
	for _, c := range e.components {
		if drawer, ok := c.(components.Drawer); ok {
			if err := drawer.Draw(screen); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Entity) Update() error {
	for _, c := range e.components {
		if updater, ok := c.(components.Updater); ok {
			if err := updater.Update(); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewEntity(options ...Option) *Entity {
	e := &Entity{}

	for _, option := range options {
		option(e)
	}

	return e
}
