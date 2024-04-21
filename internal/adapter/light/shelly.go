package light

import (
	"context"
	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port"
)

type ShellyMqtt struct {
}

func NewShellyMqtt() port.Light {
	return nil
}

func (c *ShellyMqtt) ChangeColor(ctx context.Context, color *domain.Color) error {
	panic("not implemented")
}

func (c *ShellyMqtt) ChangeWhite(ctx context.Context, white *domain.White) error {
	panic("not implemented")
}

func (c *ShellyMqtt) SwitchOff(ctx context.Context) error {
	panic("not implemented")
}

func (c *ShellyMqtt) SwitchOn(ctx context.Context) error {
	panic("not implemented")
}
