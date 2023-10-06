package graphics

import (
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
)

const (
	CameraType = "Camera"
)

var (
	_ components.Component = (*Camera)(nil)
)

type Camera struct {
	entity *entity.Entity
}

// Init implements components.Component.
func (*Camera) Init() error { return nil }

// Type implements components.Component.
func (*Camera) Type() string { return CameraType }

// Viewport returns the viewport entity.
func (c *Camera) viewport() *Viewport {
	e := c.entity.Parent()
	for e != nil {
		if c := e.Component(ViewportType); c != nil {
			return c.(*Viewport)
		}

		e = e.Parent()
	}

	return nil
}

func (c *Camera) Update() error {
	viewport := c.viewport()
	if viewport == nil {
		return nil
	}

	viewport.Follow(c.entity.X, c.entity.Y)
	return nil
}

func NewCamera(e *entity.Entity) *Camera {
	return &Camera{
		entity: e,
	}
}
