package tabs

import "charm.land/lipgloss/v2"

type Styles struct {
	tab func(m TabsModel, tab int) lipgloss.Style
}

func MakeStyles() Styles {
	return Styles{
		tab: func(m TabsModel, tab int) lipgloss.Style {
			borderColor := lipgloss.Color("62")
			if m.selected == tab {
				borderColor = lipgloss.Color("205")
			}

			return lipgloss.NewStyle().
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				Padding(0, 1).
				Align(lipgloss.Center, lipgloss.Center).
				BorderForeground(borderColor)
		},
	}
}

