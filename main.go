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

	viewport := blueprint.NewViewport()
	viewport.AddChild(blueprint.NewFloor())
	viewport.AddChild(blueprint.NewSkeleton())

	root.AddChild(viewport)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{
		logger: slog.Default(),
		root:   root,
	}); err != nil {
		log.Fatal(err)
	}
}
