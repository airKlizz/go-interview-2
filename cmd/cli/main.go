package main

import (
	"fmt"
	"log"
	"mynewgoproject/internal/adapter/driven/light"
	"mynewgoproject/internal/adapter/driving/server"
	"mynewgoproject/internal/core/service"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	mqttBroker := os.Getenv("MQTT_BROKER")
	mqttPort, err := strconv.Atoi(os.Getenv("MQTT_PORT"))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to convert MQTT_PORT to int: %w", err))
	}
	mqttUsername := os.Getenv("MQTT_USERNAME")
	shellyLightTopic := os.Getenv("SHELLY_LIGHT_TOPIC")

	bulb, err := light.NewShellyMqtt(mqttBroker, uint(mqttPort), mqttUsername, shellyLightTopic)
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
