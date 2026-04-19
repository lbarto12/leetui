package solvepage

type LangItem struct {
	lang     string
	langSlug string
}

func (l LangItem) Title() string       { return l.lang }
func (l LangItem) Description() string { return "" }
func (l LangItem) FilterValue() string { return l.lang }
