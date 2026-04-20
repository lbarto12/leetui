package tabs

import "charm.land/lipgloss/v2"

type Styles struct {
	tab func(m TabsModel, tab int) lipgloss.Style
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func MakeStyles() Styles {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder := tabBorderWithBottom("┘", " ", "└")

	return Styles{
		tab: func(m TabsModel, tab int) lipgloss.Style {
			style := lipgloss.NewStyle().
				Bold(true).
				Border(inactiveTabBorder).
				Padding(0, 1).
				Align(lipgloss.Center, lipgloss.Center).
				BorderForeground(lipgloss.Color("62"))

			if m.selected == tab {
				style = style.
					Foreground(lipgloss.Color("205")).
					Border(activeTabBorder)
				// BorderForeground(lipgloss.Color("205"))
			}

			return style
		},
	}
}
