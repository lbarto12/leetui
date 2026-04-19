// Package solvepage: solvepage def
package solvepage

import (
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
)

type SolvePageModel struct {
	viewmodel.ViewModel
	style Style
}

func MakeSolvePageModel() SolvePageModel {
	return SolvePageModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		style: MakeStyles(),
	}
}

func (m SolvePageModel) Init() tea.Cmd {
	return nil
}

func (m SolvePageModel) View() tea.View {
	return tea.NewView("Solve")
}

func (m SolvePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
