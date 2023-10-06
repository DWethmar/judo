package skeleton

import (
	"github.com/dwethmar/judo/animation"
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/components/graphics"
	"github.com/dwethmar/judo/direction"
	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AnimationType  = "skeleton.graphics"
	AnimationSpeed = 15 // frames per image
)

var (
	_ components.Updater = (*Animation)(nil)
)

type Animation struct {
	entity    *entity.Entity
	direction direction.Direction
	pX, pY    int32 // previous x, y
	animation *animation.Skeleton
	sprite    *graphics.Sprite
	hasMoved  bool
}

func (g *Animation) Init() error {
	if g.animation.Current == nil {
		g.animation.Current = g.animation.IdleDown
	}

	if g.sprite == nil {
		g.sprite = graphics.NewSprite(g.entity)
		g.entity.AddComponent(g.sprite)
	}

	return nil
}

// Type implements components.Drawer.
func (*Animation) Type() string { return AnimationType }

func (g *Animation) setSprite(image *ebiten.Image) {
	g.sprite.Image = image
	g.sprite.OffsetX = int32(-image.Bounds().Dx() / 2)
	g.sprite.OffsetY = int32(-image.Bounds().Dy())
}

// Update implements components.Updater.
func (g *Animation) Update() error {
	var animation *animation.Animation

	if g.pX != g.entity.X || g.pY != g.entity.Y {
		g.direction = direction.GetDirection(g.pX, g.pY, g.entity.X, g.entity.Y)
		g.pX, g.pY = g.entity.X, g.entity.Y
		g.hasMoved = true
	} else {
		g.hasMoved = false
	}

	if g.hasMoved {
		switch g.direction {
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
		switch g.direction {
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

	g.setSprite(g.animation.Current.Next().(*ebiten.Image))

	return nil
}

// Priority implements components.Drawer.
func (*Animation) Priority() int { return 0 }

// NewAnimation creates a new Graphic.
func NewAnimation(e *entity.Entity) *Animation {
	return &Animation{
		entity:    e,
		animation: animation.NewSkeleton(AnimationSpeed),
	}
}
