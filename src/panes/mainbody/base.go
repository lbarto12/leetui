// Package mainbody: Mainbody of application
package mainbody

import (
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyModel struct {
	viewmodel.ViewModel
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
	}
}

func (model MainBodyModel) Init() tea.Cmd {
	return nil
}

func (model MainBodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model MainBodyModel) View() tea.View {
	borderColor := lipgloss.Color("62")
	if model.IsFocused() {
		borderColor = lipgloss.Color("205")
	}

	style := lipgloss.NewStyle().
		Width(model.Dims.Width).
		Height(model.Dims.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render("Main Body"))
}
