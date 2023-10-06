package blueprint

import (
	"github.com/dwethmar/judo/components/graphics"
	"github.com/dwethmar/judo/entity"
)

func NewViewport() *entity.Entity {
	return entity.NewEntity(
		func(e *entity.Entity) { e.AddComponent(graphics.NewViewport(e)) },
	)
}
