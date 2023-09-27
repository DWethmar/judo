package skeleton

import (
	"fmt"

	"github.com/dwethmar/judo/animation"
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/direction"
	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	_ components.Drawer  = (*Graphic)(nil)
	_ components.Updater = (*Graphic)(nil)
)

type Graphic struct {
	entity    *entity.Entity
	pX, pY    int64
	animation *animation.Skeleton
}

// Update implements components.Updater.
func (g *Graphic) Update() error {
	if g.animation.Current == nil {
		g.animation.Current = g.animation.IdleDown
	}

	var animation *animation.Animation
	hasMoved := g.pX != g.entity.X || g.pY != g.entity.Y
	dir := direction.GetDirection(g.pX, g.pY, g.entity.X, g.entity.Y)

	if hasMoved {
		switch dir {
		case direction.Top, direction.TopLeft, direction.TopRight:
			animation = g.animation.WalkUp
		case direction.Bottom, direction.BottomLeft, direction.BottomRight:
			animation = g.animation.WalkDown
		case direction.Left:
			animation = g.animation.WalkLeft
		case direction.Right:
			animation = g.animation.WalkRight
		default:
			animation = g.animation.IdleDown
		}
	} else {
		switch dir {
		case direction.Top, direction.TopLeft, direction.TopRight:
			animation = g.animation.IdleUp
		case direction.Bottom, direction.BottomLeft, direction.BottomRight:
			animation = g.animation.IdleDown
		case direction.Left:
			animation = g.animation.IdleLeft
		case direction.Right:
			animation = g.animation.IdleRight
		default:
			animation = g.animation.IdleDown
		}
	}

	if g.animation.Current != animation {
		g.animation.Current.Reset()
		g.animation.Current = animation
	}

	// last
	g.animation.Current.Next()
	g.pX, g.pY = g.entity.X, g.entity.Y

	return nil
}

// Priority implements components.Drawer.
func (*Graphic) Priority() int { return 0 }

// Draw implements components.Drawer.
func (g *Graphic) Draw(screen *ebiten.Image) error {
	fmt.Println("draw graphic")

	f := g.animation.Current.Frame().(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.entity.X), float64(g.entity.Y))
	screen.DrawImage(f, op)

	return nil
}

// NewGraphic creates a new Graphic.
func NewGraphic(e *entity.Entity) *Graphic {
	return &Graphic{
		entity:    e,
		animation: animation.NewSkeleton(),
	}
}
