package floor

import (
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/matrix"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const FloorType = "floor.Floor"

const (
	GrassTile int32 = iota
	StoneTile
)

// check if Behavior implements components.Updater
var (
	_ components.Updater = (*Floor)(nil)
	_ components.Drawer  = (*Floor)(nil)
)

type Floor struct {
	initialized bool
	entity      *entity.Entity
	matrix      matrix.Matrix
}

// Init implements components.Updater.
func (b *Floor) Init() error {
	b.initialized = true
	return nil
}

// Initialized implements components.Updater.
func (b *Floor) Initialized() bool { return b.initialized }

// Type implements components.Updater.
func (*Floor) Type() string { return FloorType }

// Draw implements components.Drawer.
func (b *Floor) Draw(screen *ebiten.Image) error {
	matrix.Iterate(b.matrix, func(x, y int, value int32) {

	})

	return nil
}

// Update implements components.Updater.
func (b *Floor) Update() error {
	matrix.Iterate(b.matrix, func(x, y int, value int32) {

	})
	return nil
}

// NewFloor creates a new Behavior.
func NewFloor(entity *entity.Entity) *Floor {
	return &Floor{
		entity: entity,
		matrix: matrix.New(10, 10),
	}
}
