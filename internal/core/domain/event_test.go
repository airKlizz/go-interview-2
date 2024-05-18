package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvent_Validate(t *testing.T) {
	tests := map[string]struct {
		event   Event
		wantErr bool
	}{
		"OK valid change color event": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeColorArgs: &ChangeColorArgs{Color: &Color{Blue: 200}}},
			},
		},
		"KO wrong device": {
			event: Event{
				Target: "target",
				Device: "lightt",
				Action: "on",
			},
			wantErr: true,
		},
		"KO wrong action": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "onn",
			},
			wantErr: true,
		},
		"KO wrong args for event": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeWhiteArgs: &ChangeWhiteArgs{}},
			},
			wantErr: true,
		},
		"KO invalid change color event because Blue value 400": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeColorArgs: &ChangeColorArgs{Color: &Color{Blue: 400}}},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.event.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
