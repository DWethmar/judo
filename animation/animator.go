package animation

import (
	"image"
)

type Animation struct {
	frames              []image.Image
	frame               int
	framesPerImage      int
	framesPerImageCount int
}

func (a *Animation) Next() image.Image {
	if a.framesPerImageCount > a.framesPerImage {
		a.frame++

		if a.frame >= len(a.frames) {
			a.frame = 0
		}

		a.framesPerImageCount = 0
	} else {
		a.framesPerImageCount++
	}

	return a.Frame()
}

func (a *Animation) Frame() image.Image {
	return a.frames[a.frame]
}

func (a *Animation) Reset() {
	a.frame = 0
	a.framesPerImageCount = 0
}

func NewAnimator(frames []image.Image, framesPerImage int) *Animation {
	return &Animation{
		frames:         frames,
		framesPerImage: framesPerImage,
	}
}
