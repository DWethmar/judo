package entity

type Option func(*Entity)

func WithPosition(x, y int64) Option {
	return func(e *Entity) {
		e.X = x
		e.Y = y
	}
}
