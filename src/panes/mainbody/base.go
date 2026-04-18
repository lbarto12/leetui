// Package mainbody: Mainbody of application
package mainbody

import (
	"fmt"

	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/mainbody/components/pages"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyModel struct {
	viewmodel.ViewModel
	styles                 Styles
	selectedProblem        string // Simple ID
	selectedProblemDetails models.ProblemDetails
	children               Children
}

func MakeMainBodyModel() MainBodyModel {
	return MainBodyModel{
		ViewModel: viewmodel.ViewModel{
			Focused: true,
			Dims: viewmodel.ViewModelDims{
				Width:  0,
				Height: 0,
			},
		},
		styles:          MakeStyles(),
		selectedProblem: "Select a problem to see it here",
		children: Children{
			pages: pages.MakeMainBodyPagesModel(),
		},
	}
}

func (m MainBodyModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.pages.Init(),
	)
}

func (m MainBodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.HandleUpdate(msg)
	}
}

func (m MainBodyModel) View() tea.View {
	bodyStyle := m.styles.body(m)
	// marshalled, _ := json.Marshal(m.selectedProblemDetails)

	titleStyle := m.styles.title(m)

	view := bodyStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,

		// Title bar
		titleStyle.Render(fmt.Sprintf("%s - %s", m.selectedProblemDetails.ID, m.selectedProblemDetails.Title)),

		// Page
		m.children.pages.View().Content,
	))

	return tea.NewView(view)
}
