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

	httpServer := server.NewHttpServer(s)

	log.Println("http server runs on 0.0.0.0:8080")
	if err := httpServer.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to run http server: %w", err))
	}
}
