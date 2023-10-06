package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	terrainHeight = 32
	terrainWidth  = 32
)

var (
	terrainImg   *ebiten.Image
	terrainCells = CreateCells(10, 7, terrainWidth, terrainHeight)
)

var (
	TerrainGrassSprite             *ebiten.Image
	TerrainGrassDirtTopLeftSprite  *ebiten.Image
	TerrainGrassDirtTopSprite      *ebiten.Image
	TerrainGrassDirtTopRightSprite *ebiten.Image
	TerrainGrassDirtLeftSprite     *ebiten.Image
	TerrainGrassDirtRightSprite    *ebiten.Image
	TerrainGrassDirtBottomLeft     *ebiten.Image
	TerrainGrassDirtBottom         *ebiten.Image
	TerrainGrassDirtBottomRight    *ebiten.Image
	TerrainDirtSprite              *ebiten.Image
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/terrain.png")
	if err != nil {
		panic(err)
	}

	terrainImg = ebiten.NewImageFromImage(img)

	TerrainGrassSprite = terrainImg.SubImage(terrainCells[0][0]).(*ebiten.Image)
	TerrainGrassDirtTopLeftSprite = terrainImg.SubImage(terrainCells[0][1]).(*ebiten.Image)
	TerrainGrassDirtTopSprite = terrainImg.SubImage(terrainCells[1][1]).(*ebiten.Image)
	TerrainGrassDirtTopRightSprite = terrainImg.SubImage(terrainCells[2][1]).(*ebiten.Image)
	TerrainGrassDirtLeftSprite = terrainImg.SubImage(terrainCells[0][2]).(*ebiten.Image)
	TerrainGrassDirtRightSprite = terrainImg.SubImage(terrainCells[2][2]).(*ebiten.Image)
	TerrainGrassDirtBottomLeft = terrainImg.SubImage(terrainCells[0][3]).(*ebiten.Image)
	TerrainGrassDirtBottom = terrainImg.SubImage(terrainCells[1][3]).(*ebiten.Image)
	TerrainGrassDirtBottomRight = terrainImg.SubImage(terrainCells[2][3]).(*ebiten.Image)
	TerrainDirtSprite = terrainImg.SubImage(terrainCells[1][0]).(*ebiten.Image)
}
