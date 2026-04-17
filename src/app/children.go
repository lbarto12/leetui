package app

import (
	"leetui/src/lib/chup"

	tea "charm.land/bubbletea/v2"
)

func (s AppState) PassToChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, tea.Batch(
		chup.Forward(&s.sidebar, msg),
		chup.Forward(&s.mainbody, msg),
	)
}
