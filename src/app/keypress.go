package app

import (
	"leetui/src/app/focus"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"

	tea "charm.land/bubbletea/v2"
)

func (s AppState) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return s, tea.Quit
	case "ctrl+h":
		if !s.sidebar.IsCollapsed() {
			s.focused = focus.Sidebar
			s.mainbody.SetFocused(false)
			s.sidebar.SetFocused(true)
		}
	case "ctrl+l":
		s.focused = focus.Main
		s.sidebar.SetFocused(false)
		s.mainbody.SetFocused(true)
	case "ctrl+e":
		s.sidebar.ToggleCollapse()
		if !s.sidebar.IsCollapsed() {
			s.focused = focus.Sidebar
			s.mainbody.SetFocused(false)
			s.sidebar.SetFocused(true)
		} else {
			s.focused = focus.Main
			s.sidebar.SetFocused(false)
			s.mainbody.SetFocused(true)
		}

	default:
		switch s.focused {
		case focus.Main:
			mb, cmd := s.mainbody.Update(msg)
			s.mainbody = mb.(mainbody.MainBodyModel)
			return s, cmd
		case focus.Sidebar:
			sb, cmd := s.sidebar.Update(msg)
			s.sidebar = sb.(sidebar.SidebarModel)
			return s, cmd
		}
	}
	return s, nil
}
