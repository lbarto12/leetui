package state

import (
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"

	tea "charm.land/bubbletea/v2"
)

func (s AppState) HandleUpdateKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return s, tea.Quit
	case "ctrl+h":
		if !s.sidebar.IsCollapse() {
			s.focused = "sidebar"
			s.sidebar.SetFocused(true)
			s.mainbody.SetFocused(false)
		}

	case "ctrl+l":
		s.focused = "main"
		s.sidebar.SetFocused(false)
		s.mainbody.SetFocused(true)
	case "ctrl+e":
		s.sidebar.ToggleCollapse()
		if !s.sidebar.IsCollapse() {
			s.mainbody.SetFocused(false)
			s.sidebar.SetFocused(true)
		} else {
			s.sidebar.SetFocused(false)
			s.mainbody.SetFocused(true)
		}

	default:

		if s.focused == "sidebar" {
			updated, cmd := s.sidebar.Update(msg)
			if m, ok := updated.(sidebar.SidebarModel); ok {
				s.sidebar = m
			}
			return s, cmd
		}
		updated, cmd := s.mainbody.Update(msg)
		if m, ok := updated.(mainbody.MainBodyModel); ok {
			s.mainbody = m
		}
		return s, cmd
	}
	return s, nil
}
