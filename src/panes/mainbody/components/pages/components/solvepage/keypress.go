package solvepage

import tea "charm.land/bubbletea/v2"

func (m SolvePageModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	default:
		model, cmd := m.PassToChildren(msg)
		m = model.(SolvePageModel)

		// Check if a directory was selected
		if didSelect, path := m.children.filePicker.DidSelectFile(msg); didSelect {
			m.selectedDir = path
		}

		return m, cmd
	}
}
