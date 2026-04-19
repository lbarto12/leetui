package solvepage

import (
	"leetui/src/panes/mainbody/components/pages/components/solvepage/focus"

	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Children struct {
	filePicker filepicker.Model
	langList   list.Model
}

// PassToChildren forwards messages to all children (for non-keypress messages).
func (m SolvePageModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	fp, fpcmd := m.children.filePicker.Update(msg)
	m.children.filePicker = fp

	ll, llcmd := m.children.langList.Update(msg)
	m.children.langList = ll

	return m, tea.Batch(
		fpcmd,
		llcmd,
	)
}

// PassToFocusedChild forwards messages only to the focused child.
func (m SolvePageModel) PassToFocusedChild(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.focusedChild {
	case focus.FilePicker:
		fp, cmd := m.children.filePicker.Update(msg)
		m.children.filePicker = fp
		return m, cmd
	case focus.LangList:
		ll, cmd := m.children.langList.Update(msg)
		m.children.langList = ll
		return m, cmd
	}
	return m, nil
}
