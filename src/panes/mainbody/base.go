// Package mainbody: Mainbody of application
package mainbody

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyModel struct {
	focused bool
	width   int
	height  int
}

func MakeMainBodyModel() MainBodyModel {
	return MainBodyModel{
		focused: true,
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
	if model.focused {
		borderColor = lipgloss.Color("205")
	}

	style := lipgloss.NewStyle().
		Width(model.width).
		Height(model.height).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render("Main Body"))
}

func (model *MainBodyModel) SetSize(w, h int) {
	model.width = w
	model.height = h
}

func (model *MainBodyModel) SetFocused(f bool) {
	model.focused = f
}
