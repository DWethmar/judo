package entity

import "github.com/dwethmar/judo/systems"

type CreatedEvent struct {
	Entity *Entity
}

type AddedSystemEvent struct {
	Entity *Entity
	System systems.System
}
