package components

import "github.com/hajimehoshi/ebiten/v2"

type Drawer interface {
	Component
	Draw(screen *ebiten.Image) error
}
