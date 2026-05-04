package app

import (
	"leetui/src/app/focus"
	"leetui/src/lib/keymap"

	"charm.land/bubbles/v2/key"
)

type AppKeyMap struct {
	Quit       key.Binding
	FocusLeft  key.Binding
	FocusRight key.Binding
	Toggle     key.Binding
}

func MakeAppKeyMap() AppKeyMap {
	return AppKeyMap{
		Quit:       key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		FocusLeft:  key.NewBinding(key.WithKeys("ctrl+h"), key.WithHelp("ctrl+h", "focus left")),
		FocusRight: key.NewBinding(key.WithKeys("ctrl+l"), key.WithHelp("ctrl+l", "focus right")),
		Toggle:     key.NewBinding(key.WithKeys("ctrl+e"), key.WithHelp("ctrl+e", "toggle sidebar")),
	}
}

func (k AppKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.FocusLeft, k.FocusRight, k.Toggle, k.Quit}
}

func (k AppKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.FocusLeft, k.FocusRight, k.Toggle, k.Quit}}
}

// ActiveHelp walks the focus tree to assemble the keymap for the currently
// focused leaf. The chain is: app globals → focused pane → that pane's
// focused child → ... down to the leaf.
func (s AppState) ActiveHelp() keymap.Provider {
	chain := keymap.Merged{s.keys}
	switch s.focused {
	case focus.Sidebar:
		chain = append(chain, s.sidebar.ActiveHelp())
	case focus.Main:
		chain = append(chain, s.mainbody.ActiveHelp())
	}
	return chain
}
