package domain

type Event struct {
	Target string
	Device Device
	Action Action
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
