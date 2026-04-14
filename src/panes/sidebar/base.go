// Package sidebar: sidebar def
package sidebar

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SidebarModel struct {
	focused   bool
	collapsed bool
	Width     int
	height    int

	// components
	search textinput.Model
}

func MakeSidebarModel() SidebarModel {
	ti := textinput.New()
	ti.Placeholder = "Find problems..."
	ti.Prompt = " "
	ti.CharLimit = 128
	ti.Focus()

	return SidebarModel{
		Width:  30,
		search: ti,
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
	if m.focused {
		borderColor = lipgloss.Color("205")
	}

	m.search.SetWidth(m.Width - 4) // account for border + padding

	searchBar := m.search.View()
	content := searchBar + "\n\n" + "No results."

	style := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)

	return tea.NewView(style.Render(content))
}

func (m *SidebarModel) SetSize(w, h int) {
	m.Width = w
	m.height = h
}

func (m *SidebarModel) SetFocused(f bool) {
	m.focused = f
	if f {
		m.search.Focus()
	} else {
		m.search.Blur()
	}
}

func (m *SidebarModel) ToggleCollapse() {
	m.collapsed = !m.collapsed
}

func (m *SidebarModel) IsCollapse() bool {
	return m.collapsed
}
