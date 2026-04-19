package solvepage

import tea "charm.land/bubbletea/v2"

func (m SolvePageModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
