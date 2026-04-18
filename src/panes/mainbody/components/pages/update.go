package pages

import (
	"leetui/src/lib/common/cmds"

	tea "charm.land/bubbletea/v2"
)

func (m MainBodyPagesModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmds.SetMainBodyPageMsg:
		m.selectedPage = msg.Page
		return m, nil
	default:
		return m.PassToChildren(msg)
	}
}
