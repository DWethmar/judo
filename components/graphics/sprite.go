package graphics

import (
	"fmt"

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
	OffsetX, OffsetY int32
}

// Init implements components.Drawer.
func (*Sprite) Init() error { return nil }

// Initialized implements components.Drawer.
func (*Sprite) Initialized() bool { return true }

// Draw implements components.Drawer.
func (s *Sprite) Draw(screen *ebiten.Image) error {
	if s.Image == nil {
		fmt.Printf("image is nil\n")
		return nil
	}

	x, y := s.entity.WorldPosition()
	x += s.OffsetX
	y += s.OffsetY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(s.Image, op)
	return nil
}

// Type implements components.Drawer.
func (*Sprite) Type() string { return SpriteType }

// NewSprite creates a new Sprite.
func NewSprite(entity *entity.Entity) *Sprite {
	return &Sprite{
		entity: entity,
		Image:  nil,
	}
}
