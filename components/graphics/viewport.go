package graphics

import (
	"fmt"

	"github.com/dwethmar/judo/components"
	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/systems/scaling"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ViewportType = "Viewport"
)

var (
	_ components.Component = (*Viewport)(nil)
)

type Viewport struct {
	entity           *entity.Entity
	followX, followY int32
	screenWidth      int
	screenHeight     int
}

// Init implements components.Component.
func (*Viewport) Init() error {
	return nil
}

// Type implements components.Component.
func (*Viewport) Type() string { return ViewportType }

func (v *Viewport) Follow(x, y int32) {
	v.followX, v.followY = x, y
}

func (v *Viewport) Debug(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	screenSizeX, screenSizeY := ebiten.WindowSize()
	screemFullSizeX, screemFullSizeY := ebiten.ScreenSizeInFullscreen()

	scal := scaling.FromList(v.entity.Systems())

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("viewport x: %d, y: %d", v.entity.X, v.entity.Y), 0, 5)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("follow x: %d, y: %d", v.followX, v.followY), 0, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("screen bounds: width: %d, height: %d", screenWidth, screenHeight), 0, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("screen size  : width: %d, height: %d", screenSizeX, screenSizeY), 0, 60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("screen F size: width: %d, height: %d", screemFullSizeX, screemFullSizeY), 0, 80)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("scale: %f", scal.Scale), 0, 100)
	return nil
}

// Update aligns the viewport to the follow x and y
func (v *Viewport) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	v.entity.X = int32(screenWidth/2) - v.followX
	v.entity.Y = int32(screenHeight/2) - v.followY

	return nil
}

func NewViewport(e *entity.Entity) *Viewport {
	return &Viewport{
		entity: e,
	}
}
