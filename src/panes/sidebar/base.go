// Package sidebar: sidebar def
package sidebar

import (
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/sidebar/components"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SidebarModel struct {
	viewmodel.ViewModel

	collapsed bool

	// components
	search textinput.Model
}

func MakeSidebarModel() SidebarModel {
	return SidebarModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
			Dims: viewmodel.ViewModelDims{
				Width:  30,
				Height: 0, // inferred elswhere
			},
		},
		collapsed: false,
		search:    components.ProblemSearchBar(),
	}
}

func (m SidebarModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SidebarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.search, cmd = m.search.Update(msg)
	return m, cmd
}

func (m SidebarModel) View() tea.View {
	borderColor := lipgloss.Color("62")
	if m.IsFocused() {
		borderColor = lipgloss.Color("205")
	}

	m.search.SetWidth(m.Dims.Width - 4) // account for border + padding

	searchBar := m.search.View()
	content := searchBar + "\n\n" + "No results."

	style := lipgloss.NewStyle().
		Width(m.Dims.Width).
		Height(m.Dims.Height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render(content))
}

func (m *SidebarModel) SetFocused(f bool) { // @override
	m.Focused = f
	if f {
		m.search.Focus()
	} else {
		m.search.Blur()
	}
}

func (m *SidebarModel) ToggleCollapse() {
	m.collapsed = !m.collapsed
}

func (m *SidebarModel) IsCollapsed() bool {
	return m.collapsed
}
