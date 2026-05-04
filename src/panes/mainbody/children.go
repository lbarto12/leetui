package mainbody

import (
	"github.com/lbarto12/leetui/src/lib/chup"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages"

	tea "charm.land/bubbletea/v2"
)

type Children struct {
	pages pages.MainBodyPagesModel
}

func (m MainBodyModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Batch(
		chup.Forward(&m.children.pages, msg),
	)
}
