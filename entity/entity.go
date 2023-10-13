package entity

import (
	"errors"
	"fmt"
	"sort"

	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/systems"
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
	systems     map[string]systems.System
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

func (e *Entity) WorldPosition() (x, y int32) {
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
			return errors.New("entity and its parents do not have a bus")
		}

		if err := bus.CreatedEntity.Send(&CreatedEvent{
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

// AddSystem adds a system to the entity.
func (e *Entity) AddSystem(s systems.System) {
	if e.systems == nil {
		e.systems = map[string]systems.System{}
	}

	var added bool
	if s == nil {
		delete(e.systems, s.Type())
	} else {
		e.systems[s.Type()] = s
		added = true
	}

	if added {
		if bus := e.Bus(); bus != nil {
			e.Bus().AddedSystem.Send(&AddedSystemEvent{
				Entity: e,
				System: s,
			})
		} else {
			fmt.Printf("entity does not have a bus: %v\n", e)
		}
	}
}

// Systems returns the systems of the entity and combines them with the parent systems.
func (e *Entity) inheritedSystems() map[string]systems.System {
	s := map[string]systems.System{}
	if e.parent != nil {
		for k, v := range e.parent.inheritedSystems() {
			s[k] = v
		}
	}

	if e.systems != nil {
		for k, v := range e.systems {
			s[k] = v
		}
	}

	return s
}

func (e *Entity) Systems() []systems.System {
	s := []systems.System{}
	for _, v := range e.inheritedSystems() {
		s = append(s, v)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Type() < s[j].Type()
	})

	return s
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
