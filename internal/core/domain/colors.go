package domain

type Color struct {
	Red   int32 `validate:"min=0,max=255"`
	Green int32 `validate:"min=0,max=255"`
	Blue  int32 `validate:"min=0,max=255"`
	White int32 `validate:"min=0,max=255"`
	Gain  int32 `validate:"min=0,max=100"`
}

type White struct {
	Temp       int32 `validate:"min=3000,max=6500"`
	Brightness int32 `validate:"min=0,max=100"`
}
