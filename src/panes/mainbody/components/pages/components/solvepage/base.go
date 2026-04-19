// Package solvepage: solvepage def
package solvepage

import (
	"fmt"
	"os"

	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"
	"leetui/src/panes/mainbody/components/pages/components/solvepage/focus"

	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SolvePageModel struct {
	viewmodel.ViewModel
	style        Style
	children     Children
	selectedDir  string
	selectedLang string
	focusedChild int
}

func MakeSolvePageModel() SolvePageModel {
	fp := filepicker.New()
	fp.AllowedTypes = []string{}
	fp.CurrentDirectory, _ = os.UserHomeDir()
	fp.DirAllowed = true

	ll := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	ll.SetShowTitle(false)
	ll.SetShowHelp(false)
	ll.SetFilteringEnabled(false)
	ll.SetShowStatusBar(false)

	return SolvePageModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
		},
		style:        MakeStyles(),
		focusedChild: focus.FilePicker,
		children: Children{
			filePicker: fp,
			langList:   ll,
		},
	}
}

func (m *SolvePageModel) ResetSelection() {
	m.selectedLang = ""
	m.focusedChild = focus.FilePicker
}

func (m *SolvePageModel) SetLanguages(snippets []models.CodeSnippet) {
	items := make([]list.Item, len(snippets))
	for i, s := range snippets {
		items[i] = LangItem{lang: s.Lang, langSlug: s.LangSlug}
	}
	m.children.langList.SetItems(items)
}

func (m SolvePageModel) Init() tea.Cmd {
	return tea.Batch(
		m.children.filePicker.Init(),
	)
}

const labelsHeight = 2

func (m SolvePageModel) View() tea.View {
	innerHeight := m.Dims.Height - 2 - labelsHeight // border (2) + labels (2)

	// Directory label
	dirSelection := m.style.selected.Render("N/A")
	if m.selectedDir != "" {
		dirSelection = m.style.selected.Render(m.selectedDir)
	}
	dirLabel := fmt.Sprintf("%s %s", m.style.title.Render("Directory:"), dirSelection)

	// Language label
	langSelection := m.style.selected.Render("N/A")
	if m.selectedLang != "" {
		langSelection = m.style.selected.Render(m.selectedLang)
	}
	langLabel := fmt.Sprintf("%s %s", m.style.title.Render("Language:"), langSelection)

	panelWidth := (m.Dims.Width - 4) / 2

	fpBorder := m.style.border(m.focusedChild == focus.FilePicker).
		Width(panelWidth).
		Height(innerHeight).
		MaxHeight(innerHeight + 2)

	llBorder := m.style.border(m.focusedChild == focus.LangList).
		Width(panelWidth).
		Height(innerHeight).
		MaxHeight(innerHeight + 2)

	picker := fpBorder.Render(m.children.filePicker.View())
	langList := llBorder.Render(m.children.langList.View())

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		dirLabel,
		langLabel,
		lipgloss.JoinHorizontal(lipgloss.Top, picker, langList),
	)

	return tea.NewView(content)
}

func (m SolvePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		innerHeight := msg.Height - 2 - labelsHeight // border (2) + labels (2)
		panelWidth := (msg.Width - 4) / 2
		m.children.filePicker.SetHeight(innerHeight - 1)
		m.children.langList.SetSize(panelWidth, innerHeight)
		return m, nil
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.PassToChildren(msg)
	}
}
