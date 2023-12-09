package app

import (
	"log"

	"github.com/KartoonYoko/alice-skill/internal/config"
	"github.com/KartoonYoko/alice-skill/internal/controller"
)

func Run() {
	controller := controller.New(config.New())

	if err := controller.Serve(); err != nil {
		log.Fatal(controller.Serve())
	}
}
