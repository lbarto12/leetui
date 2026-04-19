package mainbody

import "charm.land/lipgloss/v2"

type Styles struct {
	body  func(m MainBodyModel) lipgloss.Style
	title func(m MainBodyModel) lipgloss.Style
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
				Align(lipgloss.Left, lipgloss.Top).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor)

			return bodyStyle
		},
		title: func(m MainBodyModel) lipgloss.Style {
			return lipgloss.NewStyle().
				Width(m.Dims.Width-2).
				Bold(true).
				Foreground(lipgloss.Color("5")).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder())
		},
	}
}
