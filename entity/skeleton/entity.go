package skeleton

import "github.com/dwethmar/judo/entity"

func NewSkeleton(x, y int64) *entity.Entity {
	return entity.NewEntity(
		entity.WithPosition(x, y),
		func(e *entity.Entity) { e.AddComponent(NewBehavior(e)) },
		func(e *entity.Entity) { e.AddComponent(NewGraphic(e)) },
	)
}
