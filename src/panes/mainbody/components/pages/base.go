// Package pages: Page swapper for multi-panel tabbed interface
package pages

import (
	"github.com/lbarto12/leetui/src/lib/graphqlapi/models"
	"github.com/lbarto12/leetui/src/lib/viewmodel"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/descriptionpage"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/solvepage"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/tabs"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/focus"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyPagesModel struct {
	viewmodel.ViewModel
	children       Children
	keys           PagesKeyMap
	selectedPage   int
	problemDetails models.ProblemDetails
}

func MakeMainBodyPagesModel() MainBodyPagesModel {
	return MainBodyPagesModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		children: Children{
			tabs: tabs.MakeTabsModel([]string{
				"Description",
				"Solve",
			}),
			descriptionPage: descriptionpage.MakeDescriptionPageModel(),
			solvePage:       solvepage.MakeSolvePageModel(),
		},
		keys:         MakePagesKeyMap(),
		selectedPage: focus.DescriptionPage,
	}
}

func (m MainBodyPagesModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.descriptionPage.Init(),
		m.children.solvePage.Init(),
	)
}

func (m MainBodyPagesModel) View() tea.View {
	content := "Pages"

	switch m.selectedPage {
	case focus.DescriptionPage:
		content = m.children.descriptionPage.View().Content
	case focus.SolvePage:
		content = m.children.solvePage.View().Content
	}

	view := lipgloss.JoinVertical(lipgloss.Top, m.children.tabs.View().Content, content)

	return tea.NewView(view)
}

func (m MainBodyPagesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.HandleUpdate(msg)
	}
}
