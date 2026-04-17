package sidebar

import tea "charm.land/bubbletea/v2"

func (m SidebarModel) HandleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch m.focusedChild {
	case focus.SearchBar:
		switch msg.String() {
		case "enter":
			m.search.Blur()
			m.focusedChild = focus.ProblemList
			return m, nil
		default:
			var cmd1, cmd2 tea.Cmd

			m.search, cmd1 = m.search.Update(msg)

			searchMsg := problemlist.SearchQueryMsg{Query: m.search.Value()}
			var listModel tea.Model
			listModel, cmd2 = m.problemlist.Update(searchMsg)
			m.problemlist = listModel.(problemlist.ProblemlistViewModel)
			return m, tea.Batch(cmd1, cmd2)
		}

	case focus.ProblemList:
		switch msg.String() {
		case "esc":
			m.search.Focus()
			m.focusedChild = focus.SearchBar
			return m, nil
		default:
			pl, cmd := m.problemlist.Update(msg)
			m.problemlist = pl.(problemlist.ProblemlistViewModel)
			return m, cmd
		}
	}
	return m, nil
}
