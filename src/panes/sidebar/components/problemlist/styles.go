package problemlist

import (
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
)

type Styles struct {
	tableStyles   table.Styles
	selectedStyle lipgloss.Style
	spinnerStyle  lipgloss.Style
}

func MakeStyles() Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	selectedStyle := s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	// Base styles start with no highlight since table starts unfocused
	s.Selected = lipgloss.NewStyle()

	return Styles{
		tableStyles:   s,
		selectedStyle: selectedStyle,
		spinnerStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
	}
}
