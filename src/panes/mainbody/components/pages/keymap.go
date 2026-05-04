package pages

import (
	"github.com/lbarto12/leetui/src/lib/keymap"
	"github.com/lbarto12/leetui/src/panes/mainbody/components/pages/focus"

	"charm.land/bubbles/v2/key"
)

type PagesKeyMap struct {
	PrevTab key.Binding
	NextTab key.Binding
}

func MakePagesKeyMap() PagesKeyMap {
	// Note: bubbletea v2 reports shift+letter as the uppercase letter ("H"/"L")
	// — that's what the keypress handler matches on. The Help text spells it
	// out for the footer.
	return PagesKeyMap{
		PrevTab: key.NewBinding(key.WithKeys("H"), key.WithHelp("shift+h", "prev tab")),
		NextTab: key.NewBinding(key.WithKeys("L"), key.WithHelp("shift+l", "next tab")),
	}
}

func (k PagesKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.PrevTab, k.NextTab}
}

func (k PagesKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.PrevTab, k.NextTab}}
}

func (m MainBodyPagesModel) ActiveHelp() keymap.Provider {
	chain := keymap.Merged{m.keys}
	switch m.selectedPage {
	case focus.DescriptionPage:
		chain = append(chain, m.children.descriptionPage.ActiveHelp())
	case focus.SolvePage:
		chain = append(chain, m.children.solvePage.ActiveHelp())
	}
	return chain
}
