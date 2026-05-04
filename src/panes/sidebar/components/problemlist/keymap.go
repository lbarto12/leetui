package problemlist

import "charm.land/bubbles/v2/key"

type ProblemListKeyMap struct {
	Up   key.Binding
	Down key.Binding
	Open key.Binding
}

func MakeProblemListKeyMap() ProblemListKeyMap {
	return ProblemListKeyMap{
		Up:   key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
		Down: key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
		Open: key.NewBinding(key.WithKeys("enter", "l"), key.WithHelp("⏎/l", "open problem")),
	}
}

func (k ProblemListKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Open}
}

func (k ProblemListKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Up, k.Down, k.Open}}
}

func (m ProblemlistViewModel) ActiveHelp() ProblemListKeyMap {
	return m.keys
}
