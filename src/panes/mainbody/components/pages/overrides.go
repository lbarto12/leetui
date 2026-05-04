package pages

import (
	"github.com/lbarto12/leetui/src/lib/graphqlapi/models"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/focus"
)

func (m *MainBodyPagesModel) SetPage(page int) {
	m.selectedPage = page
}

func (m *MainBodyPagesModel) SetProblemDetails(details models.ProblemDetails) {
	m.problemDetails = details
	m.children.descriptionPage.SetContent(details.Content)
	m.children.solvePage.SetProblemDetails(details)
}

func (m *MainBodyPagesModel) ResetSelections() {
	m.selectedPage = focus.DescriptionPage
	m.children.tabs.SetSelected(0)
	m.children.solvePage.ResetSelection()
}
