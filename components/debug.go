package components

import "github.com/hajimehoshi/ebiten/v2"

type Debugger interface {
	Component
	Debug(screen *ebiten.Image) error
}
