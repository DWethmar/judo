package animation

import (
	"image"

	"github.com/dwethmar/judo/assets"
)

type Skeleton struct {
	Current   *Animation
	IdleDown  *Animation
	IdleUp    *Animation
	IdleLeft  *Animation
	IdleRight *Animation
	WalkDown  *Animation
	WalkUp    *Animation
	WalkLeft  *Animation
	WalkRight *Animation
	Kill      *Animation
}

func NewSkeleton() *Skeleton {
	framesPerImage := 15

	return &Skeleton{
		IdleDown:  NewAnimator([]image.Image{assets.SkeletonDown1Sprite}, framesPerImage),
		IdleUp:    NewAnimator([]image.Image{assets.SkeletonUp1Sprite}, framesPerImage),
		IdleLeft:  NewAnimator([]image.Image{assets.SkeletonLeft1Sprite}, framesPerImage),
		IdleRight: NewAnimator([]image.Image{assets.SkeletonRight1Sprite}, framesPerImage),
		WalkDown:  NewAnimator(assets.SkeletonMoveDownFrames, framesPerImage),
		WalkUp:    NewAnimator(assets.SkeletonMoveUpFrames, framesPerImage),
		WalkLeft:  NewAnimator(assets.SkeletonMoveLeftFrames, framesPerImage),
		WalkRight: NewAnimator(assets.SkeletonMoveRightFrames, framesPerImage),
		Kill:      NewAnimator(assets.SkeletonKillFrames, framesPerImage),
	}
}
