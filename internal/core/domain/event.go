package domain

import (
	"github.com/go-playground/validator/v10"
)

type Event struct {
	Target string `validate:"required"`
	Device Device `validate:"oneof=light"`
	Action Action `validate:"oneof=on off change_color change_white"`
	Args   *Args
}

type Device string

const (
	Light = "light"
)

type Action string

const (
	On          = "on"
	Off         = "off"
	ChangeColor = "change_color"
	ChangeWhite = "change_white"
)

type Args struct {
	OnArgs          *OnArgs
	OffArgs         *OffArgs
	ChangeColorArgs *ChangeColorArgs
	ChangeWhiteArgs *ChangeWhiteArgs
}

type OnArgs struct{}

type OffArgs struct{}

type ChangeColorArgs struct {
	Color *Color
}

type ChangeWhiteArgs struct {
	White *White
}

func (e *Event) Validate() (valid bool, reasons []string) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(e)
	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			reasons = append(reasons, fieldErr.Error())
		}
		valid = false
		return
	}

	switch e.Action {
	case ChangeColor:
		if e.Args == nil || e.Args.ChangeColorArgs == nil || e.Args.ChangeColorArgs.Color == nil {
			return false, []string{"missing args for action"}
		}
	case ChangeWhite:
		if e.Args == nil || e.Args.ChangeWhiteArgs == nil || e.Args.ChangeWhiteArgs.White == nil {
			return false, []string{"missing args for action"}
		}
	}

	return true, nil

}
