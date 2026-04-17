// Package state: Defines general app state. Sobmodules may include state for stateful components
package state

import (
	"log"

	"leetui/src/lib/graphqlapi"
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"
	"leetui/src/state/focus"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/spinner"
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

		// Set children sizes on resize
		s.sidebar.SetSize(s.sidebar.GetSize().Width, msg.Height)
		s.mainbody.SetSize(msg.Width-s.sidebar.GetSize().Width, msg.Height)

	// Child passthrough
	case graphqlapi.ProblemsLoadedMsg:
		sb, _ := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)

	case cursor.BlinkMsg:
		sb, cmd := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)
		return s, cmd

	// All loading spinners
	case spinner.TickMsg:
		sb, cmd1 := s.sidebar.Update(msg)
		s.sidebar = sb.(sidebar.SidebarModel)

		mb, cmd2 := s.mainbody.Update(msg)
		s.mainbody = mb.(mainbody.MainBodyModel)

		return s, tea.Batch(cmd1, cmd2)
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
