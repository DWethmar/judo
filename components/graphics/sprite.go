package graphics

import (
	"github.com/dwethmar/judo/assets"
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType = "sprite"

var (
	_ components.Drawer = (*Sprite)(nil)
)

type Sprite struct {
	entity           *entity.Entity
	Image            *ebiten.Image
	OffsetX, OffsetY int64
}

// Init implements components.Drawer.
func (*Sprite) Init() error { return nil }

// Initialized implements components.Drawer.
func (*Sprite) Initialized() bool { return true }

// Draw implements components.Drawer.
func (s *Sprite) Draw(screen *ebiten.Image) error {
	if s.Image == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.OffsetX), float64(s.OffsetY))
	op.GeoM.Translate(float64(s.entity.X), float64(s.entity.Y))
	screen.DrawImage(s.Image, op)
	return nil
}

// Type implements components.Drawer.
func (*Sprite) Type() string { return SpriteType }

// NewSprite creates a new Sprite.
func NewSprite(entity *entity.Entity) *Sprite {
	return &Sprite{
		entity: entity,
		Image:  ebiten.NewImageFromImage(assets.Symbol29Sprite),
	}
}
