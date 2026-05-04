package descriptionpage

import "charm.land/bubbles/v2/key"

type DescriptionPageKeyMap struct{}

func (k DescriptionPageKeyMap) ShortHelp() []key.Binding  { return nil }
func (k DescriptionPageKeyMap) FullHelp() [][]key.Binding { return nil }

func (m DescriptionPageModel) ActiveHelp() DescriptionPageKeyMap {
	return DescriptionPageKeyMap{}
}
