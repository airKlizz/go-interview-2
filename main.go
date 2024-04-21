package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	// Init the MQTT client
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetUsername("mynewgoproject")
	client := MQTT.NewClient(opts)
	if token := client.Connect(); !token.WaitTimeout(time.Second) || token.Error() != nil {
		log.Fatal(fmt.Errorf("failed to init mqtt client: %w", token.Error()))
	}

	// Publish message to the bulb topic
	data := `
	{
		"mode": "color",    
		"red": 255,           
		"green": 0,         
		"blue": 0,        
		"gain": 100,        
		"brightness": 0,  
		"white": 0,         
		"temp": 0,       
		"effect": 0,        
		"turn": "on",       
		"transition": 500  
	}
	`
	token := client.Publish("shellies/shellycolorbulb-mock/color/0/set", 0, false, data)
	if !token.WaitTimeout(time.Second) || token.Error() != nil {
		log.Fatal(fmt.Errorf("failed to publish data to mqtt: %w", token.Error()))
	}

	log.Println("message successfully sent to mqtt")
}
