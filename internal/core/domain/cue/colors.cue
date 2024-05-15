package cue

#Color: {
	Red:   uint & >=0 & <=255
	Green: uint & >=0 & <=255
	Blue:  uint & >=0 & <=255
	White: uint & >=0 & <=255
	Gain:  uint & >=0 & <=100
}

#White: {
	Temp:       uint & >=3000 & <=6500
	Brightness: uint & >=0 & <=100
}