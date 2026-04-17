// Package sidebar: sidebar def
package sidebar

import (
	"leetui/src/lib/graphqlapi"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/sidebar/components"
	"leetui/src/panes/sidebar/components/problemlist"
	"leetui/src/panes/sidebar/focus"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/spinner"
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
	case tea.KeyMsg:
		switch m.focusedChild {
		case focus.SearchBar:
			switch msg.String() {
			case "enter":
				m.search.Blur()
				m.focusedChild = focus.ProblemList
				return m, nil
			default:
				var cmd1, cmd2 tea.Cmd

				m.search, cmd1 = m.search.Update(msg)

				searchMsg := problemlist.SearchQueryMsg{Query: m.search.Value()}
				var listModel tea.Model
				listModel, cmd2 = m.problemlist.Update(searchMsg)
				m.problemlist = listModel.(problemlist.ProblemlistViewModel)
				return m, tea.Batch(cmd1, cmd2)
			}

		case focus.ProblemList:
			switch msg.String() {
			case "esc":
				m.search.Focus()
				m.focusedChild = focus.SearchBar
				return m, nil
			default:
				pl, cmd := m.problemlist.Update(msg)
				m.problemlist = pl.(problemlist.ProblemlistViewModel)
				return m, cmd
			}
		}

	case graphqlapi.ProblemsLoadedMsg:
		listModel, _ := m.problemlist.Update(msg)
		m.problemlist = listModel.(problemlist.ProblemlistViewModel)

	case spinner.TickMsg:
		lm, cmd := m.problemlist.Update(msg)
		m.problemlist = lm.(problemlist.ProblemlistViewModel)
		return m, cmd

	case cursor.BlinkMsg:
		if m.focusedChild == focus.SearchBar {
			var cmd tea.Cmd
			m.search, cmd = m.search.Update(msg)
			return m, cmd
		}
	}

	return m, nil
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

func (m *SidebarModel) SetFocused(f bool) { // @override
	m.Focused = f
	m.search.Focus()
}

func (m *SidebarModel) ToggleCollapse() {
	m.collapsed = !m.collapsed
}

func (m *SidebarModel) IsCollapsed() bool {
	return m.collapsed
}
