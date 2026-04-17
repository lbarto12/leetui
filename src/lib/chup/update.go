// Package chup: tools for propogating messages to children of components
package chup

import tea "charm.land/bubbletea/v2"

func Forward[T tea.Model](model *T, msg tea.Msg) tea.Cmd {
	updated, cmd := (*model).Update(msg)
	*model = any(updated).(T)
	return cmd
}
