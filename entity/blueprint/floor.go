package blueprint

import (
	"github.com/dwethmar/judo/components/floor"
	"github.com/dwethmar/judo/entity"
)

func NewFloor() *entity.Entity {
	return entity.NewEntity(
		func(e *entity.Entity) { e.AddComponent(floor.NewFloor(e)) },
	)
}
