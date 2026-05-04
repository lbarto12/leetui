package problemlist

import (
	"context"

	"github.com/lbarto12/leetui/src/lib/common/cmds"
	"github.com/lbarto12/leetui/src/lib/graphqlapi"

	tea "charm.land/bubbletea/v2"
)

func (m ProblemlistViewModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "l":
		selected := m.table.SelectedRow()
		if selected == nil {
			return m, nil
		}

		cmd1 := cmds.SelectProblem(cmds.SelectProblemMsg{
			ProblemID: selected[0], // problem id
		})

		// Get problem from problemlist
		slug := ""
		for _, p := range m.problems {
			if p.ID == selected[0] {
				slug = p.TitleSlug
				break
			}
		}

		cmd2 := graphqlapi.GetProblemDetails(context.Background(), slug)

		return m, tea.Batch(cmd1, cmd2)
	default:
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
}
