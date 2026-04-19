// Package descriptionpage: Definition for main body description page
package descriptionpage

import (
	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type DescriptionPageModel struct {
	viewmodel.ViewModel
	viewport viewport.Model
	html     string
}

func MakeDescriptionPageModel() DescriptionPageModel {
	vp := viewport.New()
	vp.SoftWrap = true

	return DescriptionPageModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		viewport: vp,
	}
}

func (m DescriptionPageModel) Init() tea.Cmd {
	return m.viewport.Init()
}

func (m DescriptionPageModel) View() tea.View {
	return tea.NewView(m.viewport.View())
}

func (m DescriptionPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
