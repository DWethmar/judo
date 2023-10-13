package debug

import (
	"fmt"
	"image/color"

	"github.com/dwethmar/judo/assets"
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/systems/scaling"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	_ components.Debugger = (*Position)(nil)
)

type Position struct {
	entity *entity.Entity
}

// Debug implements components.Debugger.
func (p *Position) Debug(screen *ebiten.Image) error {
	scl := scaling.FromList(p.entity.Systems())

	x, y := p.entity.X, p.entity.Y
	sx, sy := scl.ScreenPosition(p.entity.WorldPosition())

	text.Draw(screen, fmt.Sprintf("x: %d, y: %d", x, y), assets.GetVGAFonts(2), int(sx), int(sy), color.White)
	return nil
}

// Init implements components.Component.
func (*Position) Init() error { return nil }

// Type implements components.Component.
func (*Position) Type() string { return "debug.Position" }

// NewPosition creates a new Position.
func NewPosition(entity *entity.Entity) *Position {
	return &Position{
		entity: entity,
	}
}
