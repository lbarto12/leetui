package pages

import (
	"leetui/src/lib/chup"
	"leetui/src/panes/mainbody/components/pages/components/descriptionpage"
	"leetui/src/panes/mainbody/components/pages/components/tabs"

	tea "charm.land/bubbletea/v2"
)

type Children struct {
	tabs            tabs.TabsModel
	descriptionPage descriptionpage.DescriptionPageModel
}

func (m MainBodyPagesModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Batch(
		chup.Forward(&m.children.tabs, msg),
		chup.Forward(&m.children.descriptionPage, msg),
	)
}
