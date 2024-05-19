package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvent_Validate(t *testing.T) {
	tests := map[string]struct {
		event       Event
		wantValid   bool
		wantReasons []string
	}{
		"OK valid change color event": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeColorArgs: &ChangeColorArgs{Color: &Color{Blue: 200}}},
			},
			wantValid: true,
		},
		"KO wrong device": {
			event: Event{
				Target: "target",
				Device: "lightt",
				Action: "on",
			},
			wantValid:   false,
			wantReasons: []string{"Key: 'Event.Device' Error:Field validation for 'Device' failed on the 'oneof' tag"},
		},
		"KO wrong action": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "onn",
			},
			wantValid:   false,
			wantReasons: []string{"Key: 'Event.Action' Error:Field validation for 'Action' failed on the 'oneof' tag"},
		},
		"KO wrong args for event": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeWhiteArgs: &ChangeWhiteArgs{}},
			},
			wantValid:   false,
			wantReasons: []string{"missing args for action"},
		},
		"KO invalid change color event because Blue value 400": {
			event: Event{
				Target: "target",
				Device: "light",
				Action: "change_color",
				Args:   &Args{ChangeColorArgs: &ChangeColorArgs{Color: &Color{Blue: 400}}},
			},
			wantValid:   false,
			wantReasons: []string{"Key: 'Event.Args.ChangeColorArgs.Color.Blue' Error:Field validation for 'Blue' failed on the 'max' tag"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			valid, reasons := tt.event.Validate()
			require.Equal(t, tt.wantValid, valid)
			require.Equal(t, tt.wantReasons, reasons)
		})
	}
}
