package solvepage

import "charm.land/bubbles/v2/key"

type SolvePageKeyMap struct {
	Clone      key.Binding
	FocusFiles key.Binding
	FocusLangs key.Binding
}

func MakeSolvePageKeyMap() SolvePageKeyMap {
	return SolvePageKeyMap{
		Clone:      key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "clone project")),
		FocusFiles: key.NewBinding(key.WithKeys("["), key.WithHelp("[", "focus file tree")),
		FocusLangs: key.NewBinding(key.WithKeys("]"), key.WithHelp("]", "focus language list")),
	}
}

func (k SolvePageKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.FocusFiles, k.FocusLangs, k.Clone}
}

func (k SolvePageKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.FocusFiles, k.FocusLangs, k.Clone}}
}

func (m SolvePageModel) ActiveHelp() SolvePageKeyMap {
	return m.keys
}
