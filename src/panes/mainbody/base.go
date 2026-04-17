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

func (m MainBodyModel) Init() tea.Cmd {
	return nil
}

func (m MainBodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}

func (m MainBodyModel) View() tea.View {
	borderColor := lipgloss.Color("62")
	if m.IsFocused() {
		borderColor = lipgloss.Color("205")
	}

	style := lipgloss.NewStyle().
		Width(m.Dims.Width).
		Height(m.Dims.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render("Main Body"))
}
