package tabs

import (
	"github.com/lbarto12/leetui/src/lib/common/cmds"

	tea "charm.land/bubbletea/v2"
)

func (m TabsModel) HandleKeyPress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "H":
		m.selected -= 1
		if m.selected < 0 {
			m.selected = len(m.options) - 1
		}
		return m, cmds.SetMainBodyPage(m.selected)
	case "L":
		m.selected += 1
		if m.selected >= len(m.options) {
			m.selected = 0
		}
		return m, cmds.SetMainBodyPage(m.selected)
	}
	return m, nil
}
