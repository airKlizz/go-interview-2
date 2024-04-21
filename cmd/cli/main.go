package main

import (
	"fmt"
	"log"
	"mynewgoproject/internal/adapter/driven/light"
	"mynewgoproject/internal/adapter/driving/server"
	"mynewgoproject/internal/core/service"
)

func main() {
	bulb, err := light.NewShellyMqtt("localhost", 1883, "shellymock", "shellies/shellycolorbulb-mock/")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init the bulb: %w", err))
	}

	c := service.NewController().WithLight("mock", bulb)

	s := service.NewServer(c)

	cliServer := server.NewCliServer(s)

	if err := cliServer.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to execute CLI: %w", err))
	}
}
