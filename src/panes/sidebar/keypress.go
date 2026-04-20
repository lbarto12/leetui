package sidebar

import (
	"leetui/src/lib/common/cmds"
	"leetui/src/panes/sidebar/components/problemlist"
	"leetui/src/panes/sidebar/focus"

	tea "charm.land/bubbletea/v2"
)

func (m SidebarModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch m.focusedChild {
	case focus.SearchBar:
		switch msg.String() {
		case "enter":
			if m.problemlist.HasProblemsAvailable() {
				m.focusedChild = focus.ProblemList
				m.search.Blur()

				m.problemlist.SetFocused(true)
			}
			return m, nil
		default:
			var cmd1, cmd2 tea.Cmd

			previousQuery := m.search.Value()
			m.search, cmd1 = m.search.Update(msg)

			// If value is the same, do not re-evaluate
			if previousQuery != m.search.Value() {
				searchMsg := cmds.SearchQueryMsg{Query: m.search.Value()}
				var listModel tea.Model
				listModel, cmd2 = m.problemlist.Update(searchMsg)
				m.problemlist = listModel.(problemlist.ProblemlistViewModel)
			}

			return m, tea.Batch(cmd1, cmd2)
		}

	case focus.ProblemList:
		switch msg.String() {
		case "esc":
			m.search.Focus()
			m.focusedChild = focus.SearchBar
			m.problemlist.SetFocused(false)
			m.problemlist.GotoTop()
			return m, nil
		default:
			pl, cmd := m.problemlist.Update(msg)
			m.problemlist = pl.(problemlist.ProblemlistViewModel)
			return m, cmd
		}
	}
	return m, nil
}
