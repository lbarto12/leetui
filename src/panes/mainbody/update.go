package mainbody

import (
	"leetui/src/lib/common/cmds"
	"leetui/src/lib/graphqlapi"

	tea "charm.land/bubbletea/v2"
)

func (m MainBodyModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmds.SelectProblemMsg:
		m.selectedProblem = msg.ProblemID
		return m, nil
	case graphqlapi.ProblemDetailsLoadedMsg:
		if msg.Err != nil {
			m.selectedProblem = "invalid problem selected"
			return m, nil
		}
		m.selectedProblemDetails = *msg.Details
		return m, nil
	default:
		_ = msg // Ignore
		return m.PassToChildren(msg)
	}
}
