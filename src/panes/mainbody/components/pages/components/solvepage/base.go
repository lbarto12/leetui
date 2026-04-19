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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Modify child sizes here
		m.SetSize(msg.Width, msg.Height)
		childMsg := tea.WindowSizeMsg{
			Width:  m.Dims.Width,
			Height: m.Dims.Height,
		}
		return m.PassToChildren(childMsg)
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		m.PassToChildren(msg)
	}
	return m, nil
}
