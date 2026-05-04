package sidebar

import (
	"github.com/lbarto12/leetui/src/panes/sidebar/focus"

	"charm.land/lipgloss/v2"
)

type Styles struct {
	body        func(m SidebarModel) lipgloss.Style
	searchStyle func(m SidebarModel) lipgloss.Style
}

func MakeStyles() Styles {
	return Styles{
		body: func(m SidebarModel) lipgloss.Style {
			borderColor := lipgloss.Color("62")
			if m.IsFocused() {
				borderColor = lipgloss.Color("205")
			}
			return lipgloss.NewStyle().
				Width(m.Dims.Width).
				Height(m.Dims.Height).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor)
		},
		searchStyle: func(m SidebarModel) lipgloss.Style {
			if m.focusedChild == focus.SearchBar {
				return lipgloss.NewStyle().
					Background(lipgloss.Color("57")).
					Foreground(lipgloss.Color("7")).
					Bold(true)
			}
			return lipgloss.NewStyle()
		},
	}
}
