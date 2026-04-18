package mainbody

import "charm.land/lipgloss/v2"

type Styles struct {
	body func(m MainBodyModel) lipgloss.Style
}

func MakeStyles() Styles {
	return Styles{
		body: func(m MainBodyModel) lipgloss.Style {
			borderColor := lipgloss.Color("62")
			if m.IsFocused() {
				borderColor = lipgloss.Color("205")
			}

			bodyStyle := lipgloss.NewStyle().
				Width(m.Dims.Width).
				Height(m.Dims.Height).
				Align(lipgloss.Center, lipgloss.Center).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor)

			return bodyStyle
		},
	}
}
