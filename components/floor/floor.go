package floor

import (
	"math/rand"

	"github.com/dwethmar/judo/assets"
	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/matrix"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	FloorType = "floor.Floor"
	FloorSize = 10
)

const (
	GrassTile int32 = 0
	DirtTile  int32 = 1
)

// check if Behavior implements components.Updater
var (
	_ components.Updater = (*Floor)(nil)
	_ components.Drawer  = (*Floor)(nil)
)

type Floor struct {
	entity *entity.Entity
	matrix matrix.Matrix
}

// Init implements components.Updater.
func (b *Floor) Init() error {
	matrix.Iterate(b.matrix, func(x, y int, value int32) {
		b.matrix.Set(x, y, rand.Int31n(2))
	})

	return nil
}

// Type implements components.Updater.
func (*Floor) Type() string { return FloorType }

// Draw implements components.Drawer.
func (b *Floor) Draw(screen *ebiten.Image) error {
	var wX, wY = b.entity.WorldPosition()

	matrix.Iterate(b.matrix, func(x, y int, value int32) {
		var img *ebiten.Image
		switch value {
		case GrassTile:
			img = assets.TerrainGrassCC
		case DirtTile:
			img = assets.TerrainDirtCC
		}

		dX := wX + int32(x*assets.TerrainWidth)
		dY := wY + int32(y*assets.TerrainHeight)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(dX), float64(dY))
		screen.DrawImage(img, op)
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
		matrix: matrix.New(FloorSize, FloorSize),
	}
}
