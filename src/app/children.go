package app

import (
	"leetui/src/lib/graphqlapi"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

func (s AppState) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case graphqlapi.ProblemsLoadedMsg:
		sb, _ := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)

	case cursor.BlinkMsg:
		sb, cmd := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)
		return s, cmd

	// All loading spinners
	case spinner.TickMsg:
		sb, cmd1 := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)

		mb, cmd2 := s.mainbody.Update(msg)
		s.mainbody = mb.(mainbody.MainBodyModel)

		return s, tea.Batch(cmd1, cmd2)
	}
	return s, nil
}
