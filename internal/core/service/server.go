package service

import (
	"context"
	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driving"
)

type Server struct {
	c *Controller
}

func NewServer(c *Controller) driving.Server {
	return &Server{
		c: c,
	}
}

func (s *Server) LightChangeColor(ctx context.Context, name string, color *domain.Color) error {
	event := &domain.Event{
		Target: name,
		Device: domain.Light,
		Action: domain.ChangeColor,
		Args: &domain.Args{
			ChangeColorArgs: &domain.ChangeColorArgs{
				Color: color,
			},
		},
	}
	valid, reasons := event.Validate()
	if !valid {
		return domain.NewErrEventNotValid(reasons)
	}
	return s.c.Handle(ctx, event)
}

func (s *Server) LightChangeWhite(ctx context.Context, name string, white *domain.White) error {
	event := &domain.Event{
		Target: name,
		Device: domain.Light,
		Action: domain.ChangeWhite,
		Args: &domain.Args{
			ChangeWhiteArgs: &domain.ChangeWhiteArgs{
				White: white,
			},
		},
	}
	valid, reasons := event.Validate()
	if !valid {
		return domain.NewErrEventNotValid(reasons)
	}
	return s.c.Handle(ctx, event)
}

func (s *Server) LightOff(ctx context.Context, name string) error {
	event := &domain.Event{
		Target: name,
		Device: domain.Light,
		Action: domain.Off,
	}
	valid, reasons := event.Validate()
	if !valid {
		return domain.NewErrEventNotValid(reasons)
	}
	return s.c.Handle(ctx, event)
}

func (s *Server) LightOn(ctx context.Context, name string) error {
	event := &domain.Event{
		Target: name,
		Device: domain.Light,
		Action: domain.On,
	}
	valid, reasons := event.Validate()
	if !valid {
		return domain.NewErrEventNotValid(reasons)
	}
	return s.c.Handle(ctx, event)
}
