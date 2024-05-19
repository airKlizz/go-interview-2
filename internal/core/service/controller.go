package service

import (
	"context"
	"fmt"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driven"
)

type Controller struct {
	lights map[string]driven.Light
}

func NewController() *Controller {
	return &Controller{
		lights: make(map[string]driven.Light),
	}
}

func (c *Controller) WithLight(name string, light driven.Light) *Controller {
	c.lights[name] = light
	return c
}

func (c *Controller) Handle(ctx context.Context, event *domain.Event) error {
	switch event.Device {
	case domain.Light:
		light, err := c.getLight(event.Target)
		if err != nil {
			return domain.ErrorDeviceNotFound
		}
		switch event.Action {
		case domain.On:
			err := light.SwitchOn(ctx)
			if err != nil {
				return domain.ErrorEventFailed
			}
		case domain.Off:
			err := light.SwitchOff(ctx)
			if err != nil {
				return domain.ErrorEventFailed
			}
		case domain.ChangeColor:
			err := light.ChangeColor(ctx, event.Args.ChangeColorArgs.Color)
			if err != nil {
				return domain.ErrorEventFailed
			}
		case domain.ChangeWhite:
			err := light.ChangeWhite(ctx, event.Args.ChangeWhiteArgs.White)
			if err != nil {
				return domain.ErrorEventFailed
			}
		}
	default:
		return domain.ErrorDeviceNotSupported
	}
	return nil
}

func (c *Controller) getLight(name string) (driven.Light, error) {
	if light, found := c.lights[name]; found {
		return light, nil
	}
	return nil, fmt.Errorf("light %s not found", name)
}
