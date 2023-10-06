package assets

import (
	"image"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	skeletonHeight = 64
	skeletonWidth  = 64
)

var (
	SkeletonImg   *ebiten.Image
	skeletonCells = CreateCells(9, 4, skeletonWidth, skeletonHeight)
)

var (
	SkeletonUp1Sprite *ebiten.Image // idle
	SkeletonUp2Sprite *ebiten.Image
	SkeletonUp3Sprite *ebiten.Image
	SkeletonUp4Sprite *ebiten.Image
	SkeletonUp5Sprite *ebiten.Image
	SkeletonUp6Sprite *ebiten.Image
	SkeletonUp7Sprite *ebiten.Image
	SkeletonUp8Sprite *ebiten.Image
	SkeletonUp9Sprite *ebiten.Image

	SkeletonLeft1Sprite *ebiten.Image // idle
	SkeletonLeft2Sprite *ebiten.Image
	SkeletonLeft3Sprite *ebiten.Image
	SkeletonLeft4Sprite *ebiten.Image
	SkeletonLeft5Sprite *ebiten.Image
	SkeletonLeft6Sprite *ebiten.Image
	SkeletonLeft7Sprite *ebiten.Image
	SkeletonLeft8Sprite *ebiten.Image
	SkeletonLeft9Sprite *ebiten.Image

	SkeletonDown1Sprite *ebiten.Image // idle
	SkeletonDown2Sprite *ebiten.Image
	SkeletonDown3Sprite *ebiten.Image
	SkeletonDown4Sprite *ebiten.Image
	SkeletonDown5Sprite *ebiten.Image
	SkeletonDown6Sprite *ebiten.Image
	SkeletonDown7Sprite *ebiten.Image
	SkeletonDown8Sprite *ebiten.Image
	SkeletonDown9Sprite *ebiten.Image

	SkeletonRight1Sprite *ebiten.Image // idle
	SkeletonRight2Sprite *ebiten.Image
	SkeletonRight3Sprite *ebiten.Image
	SkeletonRight4Sprite *ebiten.Image
	SkeletonRight5Sprite *ebiten.Image
	SkeletonRight6Sprite *ebiten.Image
	SkeletonRight7Sprite *ebiten.Image
	SkeletonRight8Sprite *ebiten.Image
	SkeletonRight9Sprite *ebiten.Image
)

var (
	SkeletonMoveUpFrames    = []image.Image{}
	SkeletonMoveLeftFrames  = []image.Image{}
	SkeletonMoveDownFrames  = []image.Image{}
	SkeletonMoveRightFrames = []image.Image{}
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/skeleton_move.png")
	if err != nil {
		panic(err)
	}

	SkeletonImg = ebiten.NewImageFromImage(img)

	SkeletonUp1Sprite = SkeletonImg.SubImage(skeletonCells[0][0]).(*ebiten.Image)
	SkeletonUp2Sprite = SkeletonImg.SubImage(skeletonCells[1][0]).(*ebiten.Image)
	SkeletonUp3Sprite = SkeletonImg.SubImage(skeletonCells[2][0]).(*ebiten.Image)
	SkeletonUp4Sprite = SkeletonImg.SubImage(skeletonCells[3][0]).(*ebiten.Image)
	SkeletonUp5Sprite = SkeletonImg.SubImage(skeletonCells[4][0]).(*ebiten.Image)
	SkeletonUp6Sprite = SkeletonImg.SubImage(skeletonCells[5][0]).(*ebiten.Image)
	SkeletonUp7Sprite = SkeletonImg.SubImage(skeletonCells[6][0]).(*ebiten.Image)
	SkeletonUp8Sprite = SkeletonImg.SubImage(skeletonCells[7][0]).(*ebiten.Image)
	SkeletonUp9Sprite = SkeletonImg.SubImage(skeletonCells[8][0]).(*ebiten.Image)

	SkeletonLeft1Sprite = SkeletonImg.SubImage(skeletonCells[0][1]).(*ebiten.Image)
	SkeletonLeft2Sprite = SkeletonImg.SubImage(skeletonCells[1][1]).(*ebiten.Image)
	SkeletonLeft3Sprite = SkeletonImg.SubImage(skeletonCells[2][1]).(*ebiten.Image)
	SkeletonLeft4Sprite = SkeletonImg.SubImage(skeletonCells[3][1]).(*ebiten.Image)
	SkeletonLeft5Sprite = SkeletonImg.SubImage(skeletonCells[4][1]).(*ebiten.Image)
	SkeletonLeft6Sprite = SkeletonImg.SubImage(skeletonCells[5][1]).(*ebiten.Image)
	SkeletonLeft7Sprite = SkeletonImg.SubImage(skeletonCells[6][1]).(*ebiten.Image)
	SkeletonLeft8Sprite = SkeletonImg.SubImage(skeletonCells[7][1]).(*ebiten.Image)
	SkeletonLeft9Sprite = SkeletonImg.SubImage(skeletonCells[8][1]).(*ebiten.Image)

	SkeletonDown1Sprite = SkeletonImg.SubImage(skeletonCells[0][2]).(*ebiten.Image)
	SkeletonDown2Sprite = SkeletonImg.SubImage(skeletonCells[1][2]).(*ebiten.Image)
	SkeletonDown3Sprite = SkeletonImg.SubImage(skeletonCells[2][2]).(*ebiten.Image)
	SkeletonDown4Sprite = SkeletonImg.SubImage(skeletonCells[3][2]).(*ebiten.Image)
	SkeletonDown5Sprite = SkeletonImg.SubImage(skeletonCells[4][2]).(*ebiten.Image)
	SkeletonDown6Sprite = SkeletonImg.SubImage(skeletonCells[5][2]).(*ebiten.Image)
	SkeletonDown7Sprite = SkeletonImg.SubImage(skeletonCells[6][2]).(*ebiten.Image)
	SkeletonDown8Sprite = SkeletonImg.SubImage(skeletonCells[7][2]).(*ebiten.Image)
	SkeletonDown9Sprite = SkeletonImg.SubImage(skeletonCells[8][2]).(*ebiten.Image)

	SkeletonRight1Sprite = SkeletonImg.SubImage(skeletonCells[0][3]).(*ebiten.Image)
	SkeletonRight2Sprite = SkeletonImg.SubImage(skeletonCells[1][3]).(*ebiten.Image)
	SkeletonRight3Sprite = SkeletonImg.SubImage(skeletonCells[2][3]).(*ebiten.Image)
	SkeletonRight4Sprite = SkeletonImg.SubImage(skeletonCells[3][3]).(*ebiten.Image)
	SkeletonRight5Sprite = SkeletonImg.SubImage(skeletonCells[4][3]).(*ebiten.Image)
	SkeletonRight6Sprite = SkeletonImg.SubImage(skeletonCells[5][3]).(*ebiten.Image)
	SkeletonRight7Sprite = SkeletonImg.SubImage(skeletonCells[6][3]).(*ebiten.Image)
	SkeletonRight8Sprite = SkeletonImg.SubImage(skeletonCells[7][3]).(*ebiten.Image)
	SkeletonRight9Sprite = SkeletonImg.SubImage(skeletonCells[8][3]).(*ebiten.Image)

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
}
