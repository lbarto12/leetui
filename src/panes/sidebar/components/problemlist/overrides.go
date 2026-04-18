package problemlist

import "charm.land/lipgloss/v2"

func (m *ProblemlistViewModel) SetFocused(focused bool) { // @override
	m.Focused = focused
	s := m.styles.tableStyles
	if m.Focused {
		m.table.Focus()
		s.Selected = m.styles.selectedStyle
	} else {
		m.table.Blur()
		s.Selected = lipgloss.NewStyle()
	}
	m.table.SetStyles(s)
}

func (m ProblemlistViewModel) HasProblemsAvailable() bool {
	return len(m.problems) > 0
}
