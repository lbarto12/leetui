package sidebar

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Styles struct {
	borderColor        color.Color
	focusedBorderColor color.Color
	border             lipgloss.Border
}

func MakeStyles() Styles {
	return Styles{
		borderColor:        lipgloss.Color("62"),
		focusedBorderColor: lipgloss.Color("205"),
		border:             lipgloss.RoundedBorder(),
	}
}
