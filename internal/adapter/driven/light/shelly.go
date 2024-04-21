package light

import (
	"context"
	"fmt"
	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driven"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type ShellyMqtt struct {
	client MQTT.Client
	topic  string
}

func NewShellyMqtt(broker string, port uint, username string, topic string) (driven.Light, error) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetUsername(username)
	var client MQTT.Client
	client = MQTT.NewClient(opts)
	if token := client.Connect(); !token.WaitTimeout(time.Second) || token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to mqtt: %w", token.Error())
	}
	return &ShellyMqtt{client: client, topic: topic}, nil
}

func (c *ShellyMqtt) ChangeColor(ctx context.Context, color *domain.Color) error {
	return c.pub(
		"/color/0/set",
		fmt.Sprintf(`
			{
				"mode": "color",    
				"red": %d,           
				"green": %d,         
				"blue": %d,        
				"gain": %d,        
				"brightness": 0,  
				"white": %d,         
				"temp": 0,       
				"effect": 0,        
				"turn": "on",       
				"transition": 500  
			}
		`, color.Red, color.Green, color.Blue, color.Gain, color.White),
	)

}

func (c *ShellyMqtt) ChangeWhite(ctx context.Context, white *domain.White) error {
	return c.pub(
		"/color/0/set",
		fmt.Sprintf(`
			{
				"mode": "white",    
				"red": 0,           
				"green": 0,         
				"blue": 0,        
				"gain": 0,        
				"brightness": %d,  
				"white": 0,         
				"temp": %d,       
				"effect": 0,        
				"turn": "on",       
				"transition": 500  
			}
		`, white.Brightness, white.Temp),
	)
}

func (c *ShellyMqtt) SwitchOff(ctx context.Context) error {
	return c.pub("/color/0/command", "off")
}

func (c *ShellyMqtt) SwitchOn(ctx context.Context) error {
	return c.pub("/color/0/command", "on")
}

func (c *ShellyMqtt) pub(path string, data any) error {
	token := c.client.Publish(c.topic+strings.TrimPrefix(path, "/"), 0, false, data)
	if !token.WaitTimeout(time.Second) || token.Error() != nil {
		return fmt.Errorf("failed to publish data to mqtt: %w", token.Error())

	}
	return nil
}
