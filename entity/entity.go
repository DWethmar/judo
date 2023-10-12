package entity

import (
	"errors"
	"fmt"

	"github.com/dwethmar/judo/components"
)

type State int

const (
	Draft State = iota // Draft is the default state
	Active
)

type Entity struct {
	X, Y        int32
	state       State
	parent      *Entity
	children    []*Entity
	composition *components.Composition
	bus         *Bus
}

func (e *Entity) State() State { return e.state }

func (e *Entity) SetState(s State) error {
	if s == Draft {
		return errors.New("cannot set state to draft")
	}

	e.state = s

	return nil
}

// Bus returns the bus of the entity.
// If the entity does not have a bus, it will return the parent bus.
func (e *Entity) Bus() *Bus {
	if e.bus == nil && e.parent != nil {
		return e.parent.Bus()
	}

	return e.bus
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

func (e *Entity) AddComponent(c components.Component) { e.composition.Add(c) }
func (e *Entity) Components() []components.Component  { return e.composition.Components() }

func (e *Entity) Component(t string) components.Component {
	var c components.Component

	for _, c = range e.composition.Components() {
		if c.Type() == t {
			return c
		}
	}

	return nil
}

func (e *Entity) Parent() *Entity     { return e.parent }
func (e *Entity) Children() []*Entity { return e.children }

func (e *Entity) AddChild(child *Entity) error {
	if child.parent != nil {
		if err := child.parent.RemoveChild(child); err != nil {
			return fmt.Errorf("failed to remove child from parent: %w", err)
		}
	}

	child.parent = e
	e.children = append(e.children, child)

	// send out event
	if child.State() == Draft {
		bus := e.Bus()
		if bus == nil {
			return errors.New("entity does not have a bus")
		}

		if err := bus.Created.Send(&CreatedEvent{
			Entity: child,
		}); err != nil {
			return fmt.Errorf("failed to send created event: %w", err)
		}
	}

	return nil
}

func (e *Entity) RemoveChild(child *Entity) error {
	if child.parent == nil {
		return errors.New("entity does not have a parent")
	}

	child.parent = nil

	for i, c := range e.children {
		if c == child {
			e.children = append(e.children[:i], e.children[i+1:]...)
			return nil
		}
	}

	return errors.New("entity not found")
}

func New(options ...Option) *Entity {
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
