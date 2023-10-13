package entity

import "github.com/dwethmar/judo/systems"

type Option func(*Entity)

func WithPosition(x, y int32) Option {
	return func(e *Entity) {
		e.X = x
		e.Y = y
	}
}

func WithBus(bus *Bus) Option            { return func(e *Entity) { e.bus = bus } }
func WithSystem(s systems.System) Option { return func(e *Entity) { e.AddSystem(s) } }
