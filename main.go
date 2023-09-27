package main

import (
	"log"
	"log/slog"

	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/entity/skeleton"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	entities := []*entity.Entity{
		skeleton.NewSkeleton(10, 10),
	}

	if err := ebiten.RunGame(&Game{
		logger:   slog.Default(),
		entities: entities,
	}); err != nil {
		log.Fatal(err)
	}
}
