// Package descriptionpage: Definition for main body description page
package descriptionpage

import (
	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/viewport"
	"charm.land/glamour/v2"
	tea "charm.land/bubbletea/v2"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
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

func (m *DescriptionPageModel) SetContent(html string) {
	m.html = html

	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		m.viewport.SetContent(html)
		return
	}

	rendered, err := glamour.Render(markdown, "dark")
	if err != nil {
		m.viewport.SetContent(markdown)
		return
	}

	m.viewport.SetContent(rendered)
}

func (m *DescriptionPageModel) SetSize(width int, height int) {
	m.ViewModel.SetSize(width, height)
	m.viewport.SetWidth(width)
	m.viewport.SetHeight(height)
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
