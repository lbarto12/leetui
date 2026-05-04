// Package mainbody: Mainbody of application
package mainbody

import (
	"fmt"

	"github.com/lbarto12/leetui/src/lib/graphqlapi/models"
	"github.com/lbarto12/leetui/src/lib/viewmodel"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	figure "github.com/common-nighthawk/go-figure"
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
			Focused: false,
			Dims: viewmodel.ViewModelDims{
				Width:  0,
				Height: 0,
			},
		},
		styles:          MakeStyles(),
		selectedProblem: "Select a problem to see it here",
		loadingSpinner:  s,
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

func (m MainBodyModel) hasProblem() bool {
	return m.selectedProblemDetails.ID != ""
}

func (m MainBodyModel) View() tea.View {
	bodyStyle := m.styles.body(m)
	innerWidth := m.Dims.Width - bodyStyle.GetHorizontalBorderSize()
	innerHeight := m.Dims.Height - bodyStyle.GetVerticalBorderSize()

	var view string
	if !m.hasProblem() && !m.loading {
		logo := figure.NewFigure("LeetUI", "larry3d", true).String()
		view = bodyStyle.Render(lipgloss.Place(
			innerWidth, innerHeight,
			lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).Render(logo),
		))
	} else if m.loading {
		view = bodyStyle.Render(lipgloss.Place(
			innerWidth, innerHeight,
			lipgloss.Center, lipgloss.Center,
			m.loadingSpinner.View()+" Getting problem...",
		))
	} else {
		titleStyle := m.styles.title(m)
		content := m.children.pages.View().Content
		view = bodyStyle.Render(lipgloss.JoinVertical(
			lipgloss.Top,
			titleStyle.Render(fmt.Sprintf("%s - %s", m.selectedProblemDetails.ID, m.selectedProblemDetails.Title)),
			content,
		))
	}

	return tea.NewView(view)
}
