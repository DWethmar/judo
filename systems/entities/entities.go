package entities

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judo/entity"
)

type Entities struct {
	logger        *slog.Logger
	draftEntities []*entity.Entity
	entities      []*entity.Entity
}

func (e *Entities) HandleEntityCreated(event *entity.CreatedEvent) error {
	e.draftEntities = append(e.draftEntities, event.Entity)
	e.logger.Debug("Entity created", "entity", event.Entity)
	return nil
}

func (e *Entities) Update() error {
	for _, entity := range e.draftEntities {
		if err := entity.Init(); err != nil {
			return fmt.Errorf("error initializing entity: %w", err)
		}

		e.logger.Debug("Entity initialized", "entity", entity)
		e.entities = append(e.entities, entity)
	}

	e.draftEntities = []*entity.Entity{}

	return nil
}

func New(logger *slog.Logger, bus *entity.Bus) *Entities {
	entities := &Entities{
		logger: logger,
	}
	bus.Created.Subscribe(entities.HandleEntityCreated)

	return entities
}
