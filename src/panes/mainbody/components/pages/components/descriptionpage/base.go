// Package descriptionpage: Definition for main body description page
package descriptionpage

import (
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
)

type DescriptionPageModel struct {
	viewmodel.ViewModel
}

func MakeDescriptionPageModel() DescriptionPageModel {
	return DescriptionPageModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
	}
}

func (m DescriptionPageModel) Init() tea.Cmd {
	return nil
}

func (m DescriptionPageModel) View() tea.View {
	return tea.NewView("Description")
}

func (m DescriptionPageModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
