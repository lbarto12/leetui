package cmds

import tea "charm.land/bubbletea/v2"

type SetMainBodyPageMsg struct {
	Page int
}

func SetMainBodyPage(page int) tea.Cmd {
	return func() tea.Msg {
		return SetMainBodyPageMsg{
			Page: page,
		}
	}
}
