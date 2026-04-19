package solvepage

import (
	"charm.land/bubbles/v2/filepicker"
	tea "charm.land/bubbletea/v2"
)

type Children struct {
	filePicker filepicker.Model
}

func (m SolvePageModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	fp, fpcmd := m.children.filePicker.Update(msg)
	m.children.filePicker = fp

	return m, tea.Batch(
		fpcmd,
	)
}
