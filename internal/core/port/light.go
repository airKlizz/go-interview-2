package port

import (
	"context"
	"mynewgoproject/internal/core/domain"
)

type Light interface {
	SwitchOn(ctx context.Context) error
	SwitchOff(ctx context.Context) error
	ChangeColor(ctx context.Context, color *domain.Color) error
	ChangeWhite(ctx context.Context, white *domain.White) error
}
