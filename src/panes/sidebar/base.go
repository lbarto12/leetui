// Package sidebar: sidebar def
package sidebar

import (
	"leetui/src/lib/chup"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/sidebar/components"
	"leetui/src/panes/sidebar/components/problemlist"
	"leetui/src/panes/sidebar/focus"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SidebarModel struct {
	viewmodel.ViewModel

	collapsed    bool
	focusedChild string

	// components
	search      textinput.Model
	problemlist problemlist.ProblemlistViewModel
}

func NewSidebarModel() (*SidebarModel, error) {
	return &SidebarModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
			Dims:    viewmodel.ViewModelDims{Width: 50, Height: 0},
		},
		collapsed:    false,
		focusedChild: focus.SearchBar,
		search:       components.ProblemSearchBar(),
		problemlist:  *problemlist.NewProblemListViewModel(),
	}, nil
}

func (m SidebarModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.problemlist.Init())
}

func (m SidebarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(m.GetSize().Width, msg.Height)
		chup.Forward(&m.problemlist, msg)
		return m, nil
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}

func (m SidebarModel) View() tea.View {
	borderColor := lipgloss.Color("62")
	if m.IsFocused() {
		borderColor = lipgloss.Color("205")
	}

	innerWidth := m.Dims.Width - 2 // border takes 2 columns
	m.search.SetWidth(innerWidth - 2)

	m.problemlist.Dims = viewmodel.ViewModelDims{
		Width:  innerWidth,
		Height: m.Dims.Height - 2 - 2, // border + (search + blank)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.search.View(),
		"",
		m.problemlist.View().Content,
	)

	style := lipgloss.NewStyle().
		Width(m.Dims.Width).
		Height(m.Dims.Height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render(content))
}
