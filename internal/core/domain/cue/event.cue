package cue

Event: {
	Target: string
	Device: "light"
}

Event:
	{
		Action: "on"
	} |
	{
		Action: "off"
	} |
	{
		Action: "change_color"
		Args: {
			ChangeColorArgs: {
				Color: #Color
			}
		}
	} |
	{
		Action: "change_white"
		Args: {
			ChangeWhiteArgs: {
				White: #White
			}
		}
	}