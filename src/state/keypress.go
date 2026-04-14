package state

import (
	"leetui/src/state/focus"

	tea "charm.land/bubbletea/v2"
)

func (s AppState) HandleUpdateKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
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
		return s.HandleChildInput(msg)
	}
	return s, nil
}
