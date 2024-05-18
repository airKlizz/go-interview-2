package service

import (
	"context"
	"testing"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driven"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServer_LightChangeColor(t *testing.T) {
	type fields struct {
		lights map[string]func() driven.Light
	}
	type args struct {
		name  string
		color *domain.Color
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr bool
	}{
		"OK change color": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						m.EXPECT().
							ChangeColor(mock.Anything, &domain.Color{Blue: 200}).
							Return(nil)
						return m
					},
				},
			},
			args: args{
				name:  "bedroom",
				color: &domain.Color{Blue: 200},
			},
		},
		"KO change color with invalid color": {
			fields: fields{
				lights: map[string]func() driven.Light{
					"bedroom": func() driven.Light {
						m := driven.NewMockLight(t)
						return m
					},
				},
			},
			args: args{
				name:  "bedroom",
				color: &domain.Color{Blue: 500},
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
			s := NewServer(c)
			err := s.LightChangeColor(context.TODO(), tt.args.name, tt.args.color)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
