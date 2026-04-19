// Package solvepage: solvepage def
package solvepage

import (
	"os"

	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/filepicker"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SolvePageModel struct {
	viewmodel.ViewModel
	style        Style
	children     Children
	selectedFile string
}

func MakeSolvePageModel() SolvePageModel {
	fp := filepicker.New()
	fp.AllowedTypes = []string{}
	fp.CurrentDirectory, _ = os.UserHomeDir()
	fp.DirAllowed = true

	return SolvePageModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		style: MakeStyles(),
		children: Children{
			filePicker: fp,
		},
	}
}

func (m SolvePageModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.filePicker.Init(),
	)
}

func (m SolvePageModel) View() tea.View {
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.children.filePicker.View(),
		),
	)

	return tea.NewView(content)
}

func (m SolvePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.children.filePicker.SetHeight(msg.Height)
		return m, nil
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}
