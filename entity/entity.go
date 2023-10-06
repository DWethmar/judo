package entity

import (
	"errors"
	"fmt"

	"github.com/dwethmar/judo/components"
)

type Entity struct {
	X, Y        int32
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

func (e *Entity) WorldPosition() (x int32, y int32) {
	x, y = e.X, e.Y

	if e.parent != nil {
		px, py := e.parent.WorldPosition()
		x += px
		y += py
	}

	return
}

func (e *Entity) AddComponent(c components.Component) {
	e.composition.Add(c)
}

func (e *Entity) Components() []components.Component {
	return e.composition.Components()
}

func (e *Entity) Component(t string) components.Component {
	var c components.Component
	for _, c = range e.composition.Components() {
		if c.Type() == t {
			return c
		}
	}

	return nil
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
