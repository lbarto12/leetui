// Package solvepage: solvepage def
package solvepage

import (
	"fmt"
	"os"

	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/filepicker"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SolvePageModel struct {
	viewmodel.ViewModel
	style       Style
	children    Children
	selectedDir string
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
	selectedText := "No folder selected"
	if m.selectedDir != "" {
		selectedText = m.selectedDir
	}
	selectedLine := fmt.Sprintf("Selected: %s", m.style.selected.Render(selectedText))

	borderStyle := m.style.border(m).
		Width(m.Dims.Width - 2).
		Height(m.Dims.Height - 4) // border (2) + selected line (1) + blank line (1)

	picker := borderStyle.Render(m.children.filePicker.View())

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		picker,
		selectedLine,
	)

	return tea.NewView(content)
}

func (m SolvePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		// Border (2) + selected line (1) + blank line (1)
		m.children.filePicker.SetHeight(msg.Height - 4)
		return m, nil
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}
