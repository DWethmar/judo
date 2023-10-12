package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/judo/entity"
	"github.com/dwethmar/judo/entity/blueprint"
	"github.com/dwethmar/judo/game"
	"github.com/dwethmar/judo/systems"
	"github.com/dwethmar/judo/systems/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	bus := &entity.Bus{}

	// systems
	entities := entities.New(logger, bus)

	// entities
	root := entity.New(entity.WithBus(bus))
	viewport := blueprint.NewViewport()

	if err := root.AddChild(viewport); err != nil {
		log.Fatal(err)
	}

	if err := viewport.AddChild(blueprint.NewFloor()); err != nil {
		log.Fatal(err)
	}

	if err := viewport.AddChild(blueprint.NewSkeleton()); err != nil {
		log.Fatal(err)
	}

	// game
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(
		game.New(slog.Default(), root, []systems.System{entities}),
	); err != nil {
		log.Fatal(err)
	}
}
