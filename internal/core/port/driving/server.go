package driving

import (
	"context"
	"mynewgoproject/internal/core/domain"
)

type Server interface {
	LightOn(ctx context.Context, name string) error
	LightOff(ctx context.Context, name string) error
	LightChangeColor(ctx context.Context, name string, color *domain.Color) error
	LightChangeWhite(ctx context.Context, name string, white *domain.White) error
}
