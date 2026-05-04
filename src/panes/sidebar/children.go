package sidebar

import (
	"github.com/lbarto12/leetui/src/lib/chup"

	tea "charm.land/bubbletea/v2"
)

func (m SidebarModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	// Manual for external component :(
	var cmd tea.Cmd
	m.search, cmd = m.search.Update(msg)
	cmds = append(cmds, cmd)

	cmds = append(cmds, chup.Forward(&m.problemlist, msg))

	return m, tea.Batch(cmds...)
}
