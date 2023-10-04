package main

import (
	"log"
	"log/slog"

	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/entity/blueprint"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	root := entity.NewEntity()

	root.AddChild(blueprint.NewFloor())
	root.AddChild(blueprint.NewSkeleton())

	if err := ebiten.RunGame(&Game{
		logger: slog.Default(),
		root:   root,
	}); err != nil {
		log.Fatal(err)
	}
}
