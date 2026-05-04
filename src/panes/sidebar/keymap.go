package sidebar

import (
	"leetui/src/lib/keymap"
	"leetui/src/panes/sidebar/focus"

	"charm.land/bubbles/v2/key"
)

type SidebarKeyMap struct {
	SearchSubmit key.Binding
	BackToSearch key.Binding
}

func MakeSidebarKeyMap() SidebarKeyMap {
	return SidebarKeyMap{
		SearchSubmit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("⏎", "go to results")),
		BackToSearch: key.NewBinding(key.WithKeys("esc", "backspace"), key.WithHelp("esc", "back to search")),
	}
}

// per-focus slices of the sidebar keymap

type searchBindings struct{ km SidebarKeyMap }

func (k searchBindings) ShortHelp() []key.Binding  { return []key.Binding{k.km.SearchSubmit} }
func (k searchBindings) FullHelp() [][]key.Binding { return [][]key.Binding{{k.km.SearchSubmit}} }

type listBindings struct{ km SidebarKeyMap }

func (k listBindings) ShortHelp() []key.Binding  { return []key.Binding{k.km.BackToSearch} }
func (k listBindings) FullHelp() [][]key.Binding { return [][]key.Binding{{k.km.BackToSearch}} }

func (m SidebarModel) ActiveHelp() keymap.Provider {
	switch m.focusedChild {
	case focus.SearchBar:
		return searchBindings{m.keys}
	case focus.ProblemList:
		return keymap.Merged{listBindings{m.keys}, m.problemlist.ActiveHelp()}
	}
	return keymap.Merged{}
}
