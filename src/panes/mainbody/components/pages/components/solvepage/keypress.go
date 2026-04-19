package solvepage

import tea "charm.land/bubbletea/v2"

func (m SolvePageModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	default:
		return m.PassToChildren(msg)
	}
}
