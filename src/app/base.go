// Package app: Defines general app state. Sobmodules may include state for stateful components
package app

import (
	"log"

	"leetui/src/app/focus"
	"leetui/src/lib/chup"
	"leetui/src/lib/common/cmds"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"

	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type AppState struct {
	focused  string
	mainbody mainbody.MainBodyModel
	sidebar  sidebar.SidebarModel
	keys     AppKeyMap
	help     help.Model

	width  int
	height int
}

func MakeAppState() AppState {
	sb, err := sidebar.NewSidebarModel()
	if err != nil {
		log.Fatalf("could not make sidebar: %v\n", sb)
	}

	as := AppState{
		mainbody: mainbody.MakeMainBodyModel(),
		sidebar:  *sb,
		focused:  focus.Sidebar,
		keys:     MakeAppKeyMap(),
		help:     help.New(),
	}

	return as
}

func (s AppState) Init() tea.Cmd {
	return tea.Batch(s.sidebar.Init(), s.mainbody.Init())
}

func (s AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// stats:
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return s.HandleKeypress(msg)
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		s.help.SetWidth(msg.Width)

		bodyHeight := msg.Height - 1 // reserve 1 row for the help footer

		// Sidebar keeps its own width
		sidebarWidth := s.sidebar.GetSize().Width
		sidebarMsg := tea.WindowSizeMsg{Width: sidebarWidth, Height: bodyHeight}
		chup.Forward(&s.sidebar, sidebarMsg)

		// Mainbody gets remaining width
		mainWidth := msg.Width - sidebarWidth
		if s.sidebar.IsCollapsed() {
			mainWidth = msg.Width
		}
		s.mainbody.SetSize(mainWidth, bodyHeight)
		mainMsg := tea.WindowSizeMsg{Width: mainWidth, Height: bodyHeight}
		chup.Forward(&s.mainbody, mainMsg)
	case cmds.SelectProblemMsg:
		s.focused = focus.Main
		s.sidebar.SetFocused(false)
		s.mainbody.SetFocused(true)
		mb, cmd := s.mainbody.Update(msg)
		s.mainbody = mb.(mainbody.MainBodyModel)
		return s, cmd
	default:
		return s.PassToChildren(msg)
	}

	return s, nil
}

func (s AppState) View() tea.View {
	bodyHeight := s.height - 1

	var left string
	mainWidth := s.width

	if !s.sidebar.IsCollapsed() {
		s.sidebar.SetSize(s.sidebar.Dims.Width, bodyHeight)
		left = s.sidebar.View().Content
		mainWidth = s.width - s.sidebar.Dims.Width
	}

	s.mainbody.SetSize(mainWidth, bodyHeight)
	right := s.mainbody.View().Content

	body := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	footer := s.help.View(s.ActiveHelp())

	v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left, body, footer))
	v.AltScreen = true
	return v
}
