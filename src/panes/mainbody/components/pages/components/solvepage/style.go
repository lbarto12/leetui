package solvepage

import "charm.land/lipgloss/v2"

type Style struct {
	border   func(m SolvePageModel) lipgloss.Style
	selected lipgloss.Style
}

func MakeStyles() Style {
	return Style{
		border: func(m SolvePageModel) lipgloss.Style {
			return lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))
		},
		selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true),
	}
}
