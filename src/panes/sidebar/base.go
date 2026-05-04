// Package sidebar: sidebar def
package sidebar

import (
	"github.com/lbarto12/leetui/src/lib/chup"
	"github.com/lbarto12/leetui/src/lib/viewmodel"
	"github.com/lbarto12/leetui/src/panes/sidebar/components"
	"github.com/lbarto12/leetui/src/panes/sidebar/components/problemlist"
	"github.com/lbarto12/leetui/src/panes/sidebar/focus"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SidebarModel struct {
	viewmodel.ViewModel

	collapsed    bool
	focusedChild string
	styles       Styles
	keys         SidebarKeyMap

	// components
	search      textinput.Model
	problemlist problemlist.ProblemlistViewModel
}

func NewSidebarModel() (*SidebarModel, error) {
	return &SidebarModel{
		ViewModel: viewmodel.ViewModel{
			Focused: true,
			Dims:    viewmodel.ViewModelDims{Width: 50, Height: 0},
		},
		collapsed:    false,
		focusedChild: focus.SearchBar,
		styles:       MakeStyles(),
		keys:         MakeSidebarKeyMap(),
		search:       components.ProblemSearchBar(),
		problemlist:  problemlist.MakeProblemListViewModel(),
	}, nil
}

func (m SidebarModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.problemlist.Init())
}

func (m SidebarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)

		// Pass inner dimensions to problemlist:
		// subtract border (2) vertically and horizontally,
		// and search bar + blank line (2) vertically
		innerWidth := msg.Width - 2
		innerHeight := msg.Height - 2 - 2
		plMsg := tea.WindowSizeMsg{Width: innerWidth, Height: innerHeight}
		chup.Forward(&m.problemlist, plMsg)
		return m, nil
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}

func (m SidebarModel) View() tea.View {
	innerWidth := m.Dims.Width - 2 // border takes 2 columns
	m.search.SetWidth(innerWidth - 3)

	m.problemlist.Dims = viewmodel.ViewModelDims{
		Width:  innerWidth,
		Height: m.Dims.Height - 2 - 2, // border + (search + blank)
	}

	// Set input styles
	inputStyle := m.search.Styles()
	inputStyle.Focused.Text = m.styles.searchStyle(m)
	inputStyle.Focused.Placeholder = m.styles.searchStyle(m)
	m.search.SetStyles(inputStyle)

	// Join view
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.search.View(),
		"",
		m.problemlist.View().Content,
	)

	return tea.NewView(m.styles.body(m).Render(content))
}
