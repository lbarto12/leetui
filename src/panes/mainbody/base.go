// Package mainbody: Mainbody of application
package mainbody

import (
	"fmt"

	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/mainbody/components/pages"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyModel struct {
	viewmodel.ViewModel
	styles                 Styles
	selectedProblem        string // Simple ID
	selectedProblemDetails models.ProblemDetails
	loading                bool
	loadingSpinner         spinner.Model
	children               Children
}

func MakeMainBodyModel() MainBodyModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return MainBodyModel{
		ViewModel: viewmodel.ViewModel{
			Focused: true,
			Dims: viewmodel.ViewModelDims{
				Width:  0,
				Height: 0,
			},
		},
		styles:         MakeStyles(),
		selectedProblem: "Select a problem to see it here",
		loadingSpinner: s,
		children: Children{
			pages: pages.MakeMainBodyPagesModel(),
		},
	}
}

func (m MainBodyModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.pages.Init(),
	)
}

func (m MainBodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.HandleUpdate(msg)
	}
}

func (m MainBodyModel) View() tea.View {
	bodyStyle := m.styles.body(m)
	titleStyle := m.styles.title(m)

	var content string
	if m.loading {
		content = lipgloss.Place(
			m.Dims.Width, m.Dims.Height,
			lipgloss.Center, lipgloss.Center,
			m.loadingSpinner.View()+" Getting problem...",
		)
	} else {
		content = m.children.pages.View().Content
	}

	view := bodyStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		titleStyle.Render(fmt.Sprintf("%s - %s", m.selectedProblemDetails.ID, m.selectedProblemDetails.Title)),
		content,
	))

	return tea.NewView(view)
}
