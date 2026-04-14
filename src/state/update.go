package state

import (
	tea "charm.land/bubbletea/v2"
)

func (s AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return s.HandleUpdateKeypress(msg)
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		sidebarWidth := 30
		mainWidth := msg.Width - sidebarWidth
		s.sidebar.SetSize(sidebarWidth, msg.Height)
		s.mainbody.SetSize(mainWidth, msg.Height)
	}

	return s, nil
}
