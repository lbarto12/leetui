package mainbody

import tea "charm.land/bubbletea/v2"

func (m MainBodyModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	if !m.hasProblem() || m.loading {
		return m, nil
	}

	switch msg.String() {
	default:
		return m.PassToChildren(msg)
	}
}
