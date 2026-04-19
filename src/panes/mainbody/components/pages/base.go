// Package pages: Page swapper for multi-panel tabbed interface
package pages

import (
	"strings"

	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/mainbody/components/pages/components/descriptionpage"
	"leetui/src/panes/mainbody/components/pages/components/tabs"
	"leetui/src/panes/mainbody/components/pages/focus"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainBodyPagesModel struct {
	viewmodel.ViewModel
	children       Children
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
				"Test 1",
				"Test 2",
			}),
			descriptionPage: descriptionpage.MakeDescriptionPageModel(),
		},
		selectedPage: focus.DescriptionPage,
	}
}

func (m MainBodyPagesModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.descriptionPage.Init(),
	)
}

func (m MainBodyPagesModel) View() tea.View {
	content := "Pages"

	switch m.selectedPage {
	case focus.DescriptionPage:
		content = m.children.descriptionPage.View().Content
	}

	separator := strings.Repeat("─", m.Dims.Width)

	view := lipgloss.JoinVertical(lipgloss.Top, m.children.tabs.View().Content, separator, content)

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
