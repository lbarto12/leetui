package state

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (s AppState) View() tea.View {
	var left string
	mainWidth := s.width

	if !s.sidebar.IsCollapse() {
		left = s.sidebar.View().Content
		mainWidth = s.width - s.sidebar.Width
	}

	s.mainbody.SetSize(mainWidth, s.height)
	right := s.mainbody.View().Content

	v := tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	v.AltScreen = true
	return v
}
