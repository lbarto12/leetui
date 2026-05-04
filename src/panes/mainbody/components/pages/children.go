package pages

import (
	"github.com/lbarto12/leetui/src/lib/chup"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/descriptionpage"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/solvepage"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/components/tabs"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/focus"

	tea "charm.land/bubbletea/v2"
)

type Children struct {
	tabs            tabs.TabsModel
	descriptionPage descriptionpage.DescriptionPageModel
	solvePage       solvepage.SolvePageModel
}

func (m MainBodyPagesModel) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmdtabs := chup.Forward(&m.children.tabs, msg)

	return m, tea.Batch(
		cmdtabs,
		chup.Forward(&m.children.descriptionPage, msg),
		chup.Forward(&m.children.solvePage, msg),
	)
}

func (m MainBodyPagesModel) PassToFocusedPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmdtabs := chup.Forward(&m.children.tabs, msg)

	var pagecmd tea.Cmd
	switch m.selectedPage {
	case focus.DescriptionPage:
		pagecmd = chup.Forward(&m.children.descriptionPage, msg)
	case focus.SolvePage:
		pagecmd = chup.Forward(&m.children.solvePage, msg)
	}

	return m, tea.Batch(
		cmdtabs,
		pagecmd,
	)
}
