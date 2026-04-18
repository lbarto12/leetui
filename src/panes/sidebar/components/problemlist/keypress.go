package problemlist

import tea "charm.land/bubbletea/v2"

func (m ProblemlistViewModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
