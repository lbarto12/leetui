package mainbody

import (
	"fmt"

	"leetui/src/lib/common/cmds"
	"leetui/src/lib/graphqlapi"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m MainBodyModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)

		// Account for body border and title bar
		bodyStyle := m.styles.body(m)
		borderH := bodyStyle.GetVerticalBorderSize()
		borderW := bodyStyle.GetHorizontalBorderSize()
		titleHeight := lipgloss.Height(m.styles.title(m).Render(
			fmt.Sprintf("%s - %s", m.selectedProblemDetails.ID, m.selectedProblemDetails.Title),
		))

		childMsg := tea.WindowSizeMsg{
			Width:  msg.Width - borderW,
			Height: msg.Height - borderH - titleHeight,
		}
		return m.PassToChildren(childMsg)
	case cmds.SelectProblemMsg:
		m.selectedProblem = msg.ProblemID
		return m, nil
	case graphqlapi.ProblemDetailsLoadedMsg:
		if msg.Err != nil {
			m.selectedProblem = "invalid problem selected"
			return m, nil
		}
		m.selectedProblemDetails = *msg.Details
		m.children.pages.SetProblemDetails(*msg.Details)
		return m, nil
	default:
		_ = msg // Ignore
		return m.PassToChildren(msg)
	}
}
