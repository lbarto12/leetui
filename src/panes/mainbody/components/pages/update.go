package pages

import (
	"github.com/lbarto12/leetui/src/lib/common/cmds"
	"github.com/lbarto12/leetui/src/lib/chup"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// tabsHeight returns the rendered height of the tabs component.
func (m MainBodyPagesModel) tabsHeight() int {
	return lipgloss.Height(m.children.tabs.View().Content)
}

func (m MainBodyPagesModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmds.SetMainBodyPageMsg:
		m.selectedPage = msg.Page
		return m, nil
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)

		// Forward full size to tabs
		chup.Forward(&m.children.tabs, msg)

		// Subtract tabs height and separator (1 line) for the page content
		pageMsg := tea.WindowSizeMsg{
			Width:  msg.Width,
			Height: msg.Height - m.tabsHeight() - 1,
		}
		chup.Forward(&m.children.descriptionPage, pageMsg)
		chup.Forward(&m.children.solvePage, pageMsg)
		return m, nil
	default:
		return m.PassToChildren(msg)
	}
}
