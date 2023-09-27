package assets

import (
	"bytes"
	"image"

	_ "embed"
	"image/png"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	skeletonMoveData, _   = staticSpritesFS.ReadFile("sprites/skeleton_move.png")
	skeletonMoveImg, _    = png.Decode(bytes.NewReader(skeletonMoveData))
	ebitenSkeletonMoveImg = ebiten.NewImageFromImage(skeletonMoveImg)
)

const SkeletonMoveSpriteSize int = 64

var (
	skeletonMoveCells = CreateCells(9, 4, SkeletonMoveSpriteSize, SkeletonMoveSpriteSize)
)

var (
	SkeletonUp1Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[0])
	SkeletonUp2Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[1])
	SkeletonUp3Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[2])
	SkeletonUp4Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[3])
	SkeletonUp5Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[4])
	SkeletonUp6Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[5])
	SkeletonUp7Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[6])
	SkeletonUp8Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[7])
	SkeletonUp9Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[8])

	SkeletonLeft1Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[9])
	SkeletonLeft2Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[10])
	SkeletonLeft3Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[11])
	SkeletonLeft4Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[12])
	SkeletonLeft5Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[13])
	SkeletonLeft6Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[14])
	SkeletonLeft7Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[15])
	SkeletonLeft8Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[16])
	SkeletonLeft9Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[17])

	SkeletonDown1Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[18])
	SkeletonDown2Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[19])
	SkeletonDown3Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[20])
	SkeletonDown4Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[21])
	SkeletonDown5Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[22])
	SkeletonDown6Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[23])
	SkeletonDown7Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[24])
	SkeletonDown8Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[25])
	SkeletonDown9Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[26])

	SkeletonRight1Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[27])
	SkeletonRight2Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[28])
	SkeletonRight3Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[29])
	SkeletonRight4Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[30])
	SkeletonRight5Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[31])
	SkeletonRight6Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[32])
	SkeletonRight7Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[33])
	SkeletonRight8Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[34])
	SkeletonRight9Sprite = ebitenSkeletonMoveImg.SubImage(skeletonMoveCells[35])
)

var (
	SkeletonMoveUpFrames = []image.Image{
		// SkeletonUp1Sprite, // idle
		SkeletonUp2Sprite,
		SkeletonUp3Sprite,
		SkeletonUp4Sprite,
		SkeletonUp5Sprite,
		SkeletonUp6Sprite,
		SkeletonUp7Sprite,
		SkeletonUp8Sprite,
		SkeletonUp9Sprite,
	}
	SkeletonMoveLeftFrames = []image.Image{
		// SkeletonLeft1Sprite, // idle
		SkeletonLeft2Sprite,
		SkeletonLeft3Sprite,
		SkeletonLeft4Sprite,
		SkeletonLeft5Sprite,
		SkeletonLeft6Sprite,
		SkeletonLeft7Sprite,
		SkeletonLeft8Sprite,
		SkeletonLeft9Sprite,
	}
	SkeletonMoveDownFrames = []image.Image{
		// SkeletonDown1Sprite, // idle
		SkeletonDown2Sprite,
		SkeletonDown3Sprite,
		SkeletonDown4Sprite,
		SkeletonDown5Sprite,
		SkeletonDown6Sprite,
		SkeletonDown7Sprite,
		SkeletonDown8Sprite,
		SkeletonDown9Sprite,
	}
	SkeletonMoveRightFrames = []image.Image{
		// SkeletonRight1Sprite, // idle
		SkeletonRight2Sprite,
		SkeletonRight3Sprite,
		SkeletonRight4Sprite,
		SkeletonRight5Sprite,
		SkeletonRight6Sprite,
		SkeletonRight7Sprite,
		SkeletonRight8Sprite,
		SkeletonRight9Sprite,
	}
)
