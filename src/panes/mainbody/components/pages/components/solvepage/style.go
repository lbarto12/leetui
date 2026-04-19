package solvepage

import "charm.land/lipgloss/v2"

type Style struct {
	border   func(focused bool) lipgloss.Style
	title    lipgloss.Style
	selected lipgloss.Style
}

func MakeStyles() Style {
	return Style{
		border: func(focused bool) lipgloss.Style {
			borderColor := lipgloss.Color("62")
			if focused {
				borderColor = lipgloss.Color("205")
			}
			return lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor)
		},
		title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("99")).
			Padding(0, 1),
		selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true),
	}
}
