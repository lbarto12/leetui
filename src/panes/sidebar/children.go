package sidebar

import (
	"leetui/src/lib/graphqlapi"
	"leetui/src/panes/sidebar/components/problemlist"
	"leetui/src/panes/sidebar/focus"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

func (m SidebarModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case graphqlapi.ProblemsLoadedMsg:
		listModel, _ := m.problemlist.Update(msg)
		m.problemlist = listModel.(problemlist.ProblemlistViewModel)

	case spinner.TickMsg:
		lm, cmd := m.problemlist.Update(msg)
		m.problemlist = lm.(problemlist.ProblemlistViewModel)
		return m, cmd

	case cursor.BlinkMsg:
		if m.focusedChild == focus.SearchBar {
			var cmd tea.Cmd
			m.search, cmd = m.search.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}
