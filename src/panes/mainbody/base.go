// Package mainbody: Mainbody of application
package mainbody

import (
	"encoding/json"

	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
)

type MainBodyModel struct {
	viewmodel.ViewModel
	styles                 Styles
	selectedProblem        string // Simple ID
	selectedProblemDetails models.ProblemDetails
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
	}
}

func (m MainBodyModel) Init() tea.Cmd {
	return nil
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
	marshalled, _ := json.Marshal(m.selectedProblemDetails)

	style := m.styles.body(m)

	return tea.NewView(style.Render(string(marshalled)))
}
