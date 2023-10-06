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

func NewSkeleton(framesPerImage int) *Skeleton {
	return &Skeleton{
		IdleDown:  NewAnimator([]image.Image{assets.SkeletonDown1}, framesPerImage),
		IdleUp:    NewAnimator([]image.Image{assets.SkeletonUp1}, framesPerImage),
		IdleLeft:  NewAnimator([]image.Image{assets.SkeletonLeft1}, framesPerImage),
		IdleRight: NewAnimator([]image.Image{assets.SkeletonRight1}, framesPerImage),
		WalkDown:  NewAnimator(assets.SkeletonMoveDownFrames, framesPerImage),
		WalkUp:    NewAnimator(assets.SkeletonMoveUpFrames, framesPerImage),
		WalkLeft:  NewAnimator(assets.SkeletonMoveLeftFrames, framesPerImage),
		WalkRight: NewAnimator(assets.SkeletonMoveRightFrames, framesPerImage),
		// Kill:      NewAnimator(assets.SkeletonKillFrames, framesPerImage),
	}
}
