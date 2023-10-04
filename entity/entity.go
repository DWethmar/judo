package entity

import (
	"errors"
	"fmt"

	"github.com/dwethmar/judo/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	X, Y        int64
	parent      *Entity
	children    []*Entity
	composition *components.Composition
}

func (e *Entity) Init() error {
	if err := e.composition.InitializeQueue(); err != nil {
		return fmt.Errorf("failed to initialize entity: %w", err)
	}

	return nil
}

func (e *Entity) AddComponent(c components.Component) {
	e.composition.Add(c)
}

func (e *Entity) Parent() *Entity {
	return e.parent
}

func (e *Entity) setParent(parent *Entity) {
	e.parent = parent
}

func (e *Entity) Children() []*Entity {
	return e.children
}

func (e *Entity) AddChild(child *Entity) error {
	if child.parent != nil {
		return errors.New("entity already has a parent")
	}

	child.setParent(e)
	e.children = append(e.children, child)
	return nil
}

func (e *Entity) RemoveChild(child *Entity) error {
	if child.parent == nil {
		return errors.New("entity does not have a parent")
	}

	child.setParent(nil)
	for i, c := range e.children {
		if c == child {
			e.children = append(e.children[:i], e.children[i+1:]...)
			return nil
		}
	}

	return errors.New("entity not found")
}

func (e *Entity) Draw(screen *ebiten.Image) error {
	for _, c := range e.composition.Components() {
		if drawer, ok := c.(components.Drawer); ok {
			if err := drawer.Draw(screen); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Entity) Update() error {
	for _, c := range e.composition.Components() {
		if updater, ok := c.(components.Updater); ok {
			if err := updater.Update(); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewEntity(options ...Option) *Entity {
	e := &Entity{
		composition: components.NewComposition(),
		children:    []*Entity{},
	}

	if len(options) > 0 {
		for _, option := range options {
			if option != nil {
				option(e)
			}
		}
	}

	return e
}
