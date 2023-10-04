package skeleton

import (
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const BehaviorType = "skeleton.behavior"

// check if Behavior implements components.Updater
var _ components.Updater = (*Behavior)(nil)

type Behavior struct {
	entity *entity.Entity
}

// Init implements components.Updater.
func (b *Behavior) Init() error {
	return nil
}

// Type implements components.Updater.
func (*Behavior) Type() string { return BehaviorType }

// Update implements components.Updater.
func (b *Behavior) Update() error {
	var x, y int64
	var speed int64 = 1

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		x = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		x = speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		y = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		y = speed
	}

	if x != 0 || y != 0 {
		b.entity.X += x
		b.entity.Y += y
	}

	return nil
}

// NewBehavior creates a new Behavior.
func NewBehavior(entity *entity.Entity) *Behavior {
	return &Behavior{
		entity: entity,
	}
}
