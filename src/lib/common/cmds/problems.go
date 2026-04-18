package cmds

import tea "charm.land/bubbletea/v2"

type SearchQueryMsg struct {
	Query string
}

type SelectProblemMsg struct {
	ProblemID string
}

func SelectProblem(msg SelectProblemMsg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
