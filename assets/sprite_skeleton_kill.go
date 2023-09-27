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
	skeletonKillData, _   = staticSpritesFS.ReadFile("sprites/skeleton_kill.png")
	skeletonKillImg, _    = png.Decode(bytes.NewReader(skeletonKillData))
	ebitenSkeletonKillImg = ebiten.NewImageFromImage(skeletonKillImg)
)

const SkeletonKillSpriteSize = 64

var (
	skeletonKillCells = CreateCells(6, 1, SkeletonKillSpriteSize, SkeletonKillSpriteSize)
)

var (
	SkeletonKill1Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[0])
	SkeletonKill2Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[1])
	SkeletonKill3Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[2])
	SkeletonKill4Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[3])
	SkeletonKill5Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[4])
	SkeletonKill6Sprite = ebitenSkeletonKillImg.SubImage(skeletonKillCells[5])
)

var (
	SkeletonKillFrames = []image.Image{
		SkeletonKill1Sprite,
		SkeletonKill2Sprite,
		SkeletonKill3Sprite,
		SkeletonKill4Sprite,
		SkeletonKill5Sprite,
		SkeletonKill6Sprite,
	}
)
