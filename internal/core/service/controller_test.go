package service

import (
	"context"
	"errors"
	"testing"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driven"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Handle(t *testing.T) {
	type fields struct {
		lights map[string]func() driven.Light
	}
	type args struct {
		event *domain.Event
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr bool
	}{
		"OK switch on light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().SwitchOn(mock.Anything).Return(nil)
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.On},
			},
		},
		"OK switch off light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().SwitchOff(mock.Anything).Return(nil)
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.Off},
			},
		},
		"OK change color light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().ChangeColor(mock.Anything, &domain.Color{
							Red:   100,
							Green: 100,
							Blue:  100,
							White: 100,
							Gain:  50,
						}).Return(nil)
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.ChangeColor,
					Args: &domain.Args{ChangeColorArgs: &domain.ChangeColorArgs{Color: &domain.Color{
						Red:   100,
						Green: 100,
						Blue:  100,
						White: 100,
						Gain:  50,
					}}},
				},
			},
		},
		"OK change white light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().ChangeWhite(mock.Anything, &domain.White{
							Temp:       3000,
							Brightness: 100,
						}).Return(nil)
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.ChangeWhite,
					Args: &domain.Args{ChangeWhiteArgs: &domain.ChangeWhiteArgs{White: &domain.White{
						Temp:       3000,
						Brightness: 100,
					}}},
				},
			},
		},
		"KO light not found": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						return driven.NewMockLight(t)
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "kitchen", Device: domain.Light, Action: domain.On},
			},
			wantErr: true,
		},
		"KO device not supported": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						return driven.NewMockLight(t)
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Device("not supported"), Action: domain.On},
			},
			wantErr: true,
		},
		"KO switch on light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().SwitchOn(mock.Anything).Return(errors.New("random"))
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.On},
			},
			wantErr: true,
		},
		"KO switch off light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().SwitchOff(mock.Anything).Return(errors.New("random"))
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.Off},
			},
			wantErr: true,
		},
		"KO change color light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().ChangeColor(mock.Anything, &domain.Color{
							Red:   100,
							Green: 100,
							Blue:  100,
							White: 100,
							Gain:  50,
						}).Return(errors.New("random"))
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.ChangeColor,
					Args: &domain.Args{ChangeColorArgs: &domain.ChangeColorArgs{Color: &domain.Color{
						Red:   100,
						Green: 100,
						Blue:  100,
						White: 100,
						Gain:  50,
					}}},
				},
			},
			wantErr: true,
		},
		"KO change white light": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().ChangeWhite(mock.Anything, &domain.White{
							Temp:       3000,
							Brightness: 100,
						}).Return(errors.New("random"))
						return m
					},
				},
			},
			args: args{
				event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.ChangeWhite,
					Args: &domain.Args{ChangeWhiteArgs: &domain.ChangeWhiteArgs{White: &domain.White{
						Temp:       3000,
						Brightness: 100,
					}}},
				},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewController()
			for name, light := range tt.fields.lights {
				c.WithLight(name, light())
			}
			err := c.Handle(context.TODO(), tt.args.event)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
