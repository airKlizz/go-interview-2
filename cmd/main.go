package main

import (
	"context"
	"fmt"
	"log"
	"mynewgoproject/internal/adapter/light"
	"mynewgoproject/internal/core/domain"
)

func main() {
	bulb, err := light.NewShellyMqtt("localhost", 1883, "shellymock", "shellies/shellycolorbulb-mock/")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init the bulb: %w", err))
	}
	err = bulb.ChangeColor(context.Background(), &domain.Color{
		Red:   0,
		Green: 255,
		Blue:  0,
		White: 0,
		Gain:  100,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to change color: %w", err))
	}
	log.Println("successfully changed color")
}
