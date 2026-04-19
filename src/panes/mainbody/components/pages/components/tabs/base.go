// Package tabs: tab manager for main body panels
package tabs

import (
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type TabsModel struct {
	viewmodel.ViewModel
	styles   Styles
	options  []string
	selected int
}

func MakeTabsModel(options []string) TabsModel {
	return TabsModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		styles:  MakeStyles(),
		options: options,
	}
}

func (m *TabsModel) SetSelected(i int) {
	m.selected = i
}

func (m TabsModel) Init() tea.Cmd {
	return nil
}

func (m TabsModel) View() tea.View {
	t := []string{}

	for i, o := range m.options {
		t = append(t, m.styles.tab(m, i).Render(o))
	}

	content := lipgloss.JoinHorizontal(lipgloss.Top, t...)

	return tea.NewView(content)
}

func (m TabsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeyPress(msg)
	}
	return m, nil
}
