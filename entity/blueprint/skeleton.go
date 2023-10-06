package blueprint

import (
	"github.com/dwethmar/judo/components/graphics"
	"github.com/dwethmar/judo/components/skeleton"
	"github.com/dwethmar/judo/entity"
)

func NewSkeleton() *entity.Entity {
	return entity.NewEntity(
		func(e *entity.Entity) { e.AddComponent(skeleton.NewBehavior(e)) },
		func(e *entity.Entity) { e.AddComponent(skeleton.NewAnimation(e)) },
		func(e *entity.Entity) { e.AddComponent(graphics.NewCamera(e)) },
	)
}
