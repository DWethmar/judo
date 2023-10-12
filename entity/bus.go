package entity

import "github.com/dwethmar/judo/event"

type Bus struct {
	Created event.Bus[*CreatedEvent]
}
