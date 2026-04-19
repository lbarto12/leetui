package descriptionpage

import (
	"charm.land/glamour/v2"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func (m *DescriptionPageModel) SetSize(width int, height int) {
	m.ViewModel.SetSize(width, height)
	m.viewport.SetWidth(width)
	m.viewport.SetHeight(height)
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
