package scaling

import (
	"github.com/dwethmar/judo/input"
	"github.com/dwethmar/judo/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Type                = "scaling"
	InterpolationFactor = 0.1 // Define how quickly you move to the target scale. Closer to 1 is faster.
)

type Scaling struct {
	Scale       float64
	TargetScale float64
}

func (s *Scaling) Type() string { return Type }

func (s *Scaling) Update() error {
	_, dy := ebiten.Wheel()

	// Update target scale based on wheel input
	if dy > 0 {
		s.TargetScale += 0.1
	} else if dy < 0 {
		s.TargetScale -= 0.1
	}

	// Ensure target scale bounds
	if s.TargetScale < 0.1 {
		s.TargetScale = 0.1
	}

	// Interpolate current scale towards target scale
	s.Scale += (s.TargetScale - s.Scale) * InterpolationFactor

	return nil
}

func (s *Scaling) WorldPosition(screenX, screenY int32) (x, y int32) {
	screenWidth, screenHeight := ebiten.WindowSize()
	return input.WorldPosition(screenX, screenY, s.Scale, screenWidth, screenHeight)
}

func (s *Scaling) ScreenPosition(worldX, worldY int32) (x, y int32) {
	screenWidth, screenHeight := ebiten.WindowSize()
	return input.ScreenPosition(worldX, worldY, s.Scale, int32(screenWidth), int32(screenHeight))
}

func New() *Scaling {
	return &Scaling{
		Scale:       1,
		TargetScale: 1,
	}
}

func FromList(s []systems.System) *Scaling {
	for _, system := range s {
		if system.Type() == Type {
			return system.(*Scaling)
		}
	}
	return nil
}
