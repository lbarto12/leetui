package pages

import "leetui/src/lib/graphqlapi/models"

func (m *MainBodyPagesModel) SetPage(page int) {
	m.selectedPage = page
}

func (m *MainBodyPagesModel) SetProblemDetails(details models.ProblemDetails) {
	m.problemDetails = details
	m.children.descriptionPage.SetContent(details.Content)
}
