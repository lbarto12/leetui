package mainbody

import "github.com/lbarto12/leetui/src/lib/keymap"

func (m MainBodyModel) ActiveHelp() keymap.Provider {
	if !m.hasProblem() || m.loading {
		return keymap.Merged{}
	}
	return m.children.pages.ActiveHelp()
}
