package entity

import "github.com/dwethmar/judo/event"

type Bus struct {
	CreatedEntity event.Bus[*CreatedEvent]
	AddedSystem   event.Bus[*AddedSystemEvent]
}
