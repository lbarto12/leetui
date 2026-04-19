package solvepage

import (
	"leetui/src/panes/mainbody/components/pages/components/solvepage/focus"

	tea "charm.land/bubbletea/v2"
)

func (m SolvePageModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "]":
		m.focusedChild = focus.LangList
		return m, nil
	case "[":
		m.focusedChild = focus.FilePicker
		return m, nil
	default:
		model, cmd := m.PassToFocusedChild(msg)
		m = model.(SolvePageModel)

		// Check if a directory was selected
		if didSelect, path := m.children.filePicker.DidSelectFile(msg); didSelect {
			m.selectedDir = path
		}

		return m, cmd
	}
}
