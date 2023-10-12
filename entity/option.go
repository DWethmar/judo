package entity

type Option func(*Entity)

func WithPosition(x, y int32) Option {
	return func(e *Entity) {
		e.X = x
		e.Y = y
	}
}

func WithBus(bus *Bus) Option {
	return func(e *Entity) {
		e.bus = bus
	}
}
