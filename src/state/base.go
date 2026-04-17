// Package state: Defines general app state. Sobmodules may include state for stateful components
package state

import (
	"log"

	"leetui/src/lib/graphqlapi"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"
	"leetui/src/state/focus"

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
		focused:  focus.Main,
	}

	return as
}

func (s AppState) Init() tea.Cmd {
	return nil
}

func (s AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// stats:
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return s.HandleUpdateKeypress(msg)
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		sidebarWidth := 30
		mainWidth := msg.Width - sidebarWidth
		s.sidebar.SetSize(sidebarWidth, msg.Height)
		s.mainbody.SetSize(mainWidth, msg.Height)

	case graphqlapi.ProblemsLoadedMsg:
		sb, _ := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)
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
