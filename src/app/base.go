// Package app: Defines general app state. Sobmodules may include state for stateful components
package app

import (
	"log"

	"leetui/src/app/focus"
	"leetui/src/lib/chup"
	"leetui/src/lib/common/cmds"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type AppState struct {
	focused  string
	mainbody mainbody.MainBodyModel
	sidebar  sidebar.SidebarModel

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

		// Sidebar keeps its own width, gets full height
		sidebarWidth := s.sidebar.GetSize().Width
		sidebarMsg := tea.WindowSizeMsg{Width: sidebarWidth, Height: msg.Height}
		chup.Forward(&s.sidebar, sidebarMsg)

		// Mainbody gets remaining width and full height
		mainWidth := msg.Width - sidebarWidth
		if s.sidebar.IsCollapsed() {
			mainWidth = msg.Width
		}
		s.mainbody.SetSize(mainWidth, msg.Height)
		mainMsg := tea.WindowSizeMsg{Width: mainWidth, Height: msg.Height}
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
	var left string
	mainWidth := s.width

	if !s.sidebar.IsCollapsed() {
		left = s.sidebar.View().Content
		mainWidth = s.width - s.sidebar.Dims.Width
	}

	s.mainbody.SetSize(mainWidth, s.height)
	right := s.mainbody.View().Content

	v := tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	v.AltScreen = true
	return v
}
