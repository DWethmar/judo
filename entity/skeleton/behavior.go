package skeleton

import (
	"fmt"

	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

// check if Behavior implements components.Updater
var _ components.Updater = (*Behavior)(nil)

type Behavior struct {
	entity *entity.Entity
}

// Priority implements components.Updater.
func (*Behavior) Priority() int { return 0 }

// Update implements components.Updater.
func (b *Behavior) Update() error {
	fmt.Printf("Behavior.Update\n")

	var x, y int64
	var speed int64 = 1

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		fmt.Printf("KeyLeft\n")
		x = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		fmt.Printf("KeyRight\n")
		x = speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		fmt.Printf("KeyUp\n")
		y = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		fmt.Printf("KeyDown\n")
		y = speed
	}

	if x != 0 || y != 0 {
		fmt.Printf("Move\n")
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
