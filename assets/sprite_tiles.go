package assets

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	tilesData, _   = staticSpritesFS.ReadFile("sprites/tiles.png")
	tilesImg, _    = png.Decode(bytes.NewReader(symbolsData))
	ebitenTilesImg = ebiten.NewImageFromImage(tilesImg)
)
