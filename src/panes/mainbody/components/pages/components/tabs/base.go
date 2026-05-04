// Package tabs: tab manager for main body panels
package tabs

import (
	"strings"

	"github.com/lbarto12/leetui/src/lib/viewmodel"

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
	content = lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		lipgloss.JoinHorizontal(lipgloss.Top, t...),
		m.styles.tabGap(m).Render(strings.Repeat(" ", max(0, m.Dims.Width-lipgloss.Width(content)-1))),
	)

	return tea.NewView(content)
}

func (m TabsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	case tea.KeyPressMsg:
		return m.HandleKeyPress(msg)
	}
	return m, nil
}
