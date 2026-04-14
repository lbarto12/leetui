package state

import (
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"
	"leetui/src/state/focus"

	tea "charm.land/bubbletea/v2"
)

func (s AppState) HandleChildInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var updated tea.Model
	var cmd tea.Cmd
	switch s.focused {
	case focus.Main:
		updated, cmd = s.mainbody.Update(msg)
		s.mainbody = updated.(mainbody.MainBodyModel)
	case focus.Sidebar:
		updated, cmd = s.sidebar.Update(msg)
		s.sidebar = updated.(sidebar.SidebarModel)
	}
	return s, cmd
}
