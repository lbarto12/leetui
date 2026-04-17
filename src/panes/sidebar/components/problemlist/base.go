// Package problemlist: Recyclerview for problems
package problemlist

import (
	"fmt"
	"log"
	"strings"

	"leetui/src/lib/graphqlapi"
	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ProblemlistViewModel struct {
	viewmodel.ViewModel
	problems []models.Problem
}

func NewProblemListViewModel() *ProblemlistViewModel {
	return &ProblemlistViewModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
			Dims: viewmodel.ViewModelDims{
				Width:  30,
				Height: 30,
			},
		},
	}
}

func (m ProblemlistViewModel) Init() tea.Cmd {
	// problems, err := graphqlapi.GetProblems(0, 100, graphqlapi.QuestionGetFilter{
	// 	SearchKeywords: "generate",
	// })
	// if err != nil {
	// 	return nil
	// }
	// m.problems = problems
	return nil
}

func (m ProblemlistViewModel) View() tea.View {
	innerWidth := m.Dims.Width
	if innerWidth < 10 {
		innerWidth = 10
	}

	content := "(no results)"
	if len(m.problems) > 0 {
		var lines []string
		for _, p := range m.problems {
			line := fmt.Sprintf("%4s [%-6s] %s", p.ID, p.Difficulty, p.Title)
			lines = append(lines, truncate(line, innerWidth))
		}
		content = strings.Join(lines, "\n")
	}

	style := lipgloss.NewStyle().
		Width(m.Dims.Width).
		Height(m.Dims.Height)

	return tea.NewView(style.Render(content))
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 1 {
		return s[:max]
	}
	return s[:max-1] + "…"
}

type SearchQueryMsg struct{ Query string }

func (m ProblemlistViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SearchQueryMsg:

		problems, err := graphqlapi.GetProblems(0, 25, graphqlapi.QuestionGetFilter{
			SearchKeywords: msg.Query,
		})
		if err != nil {
			log.Fatal(err)
		}
		m.problems = problems

		log.Printf("found new problems: %v\n", m.problems)

		// m.search.Update(msg
		//
		// m.search, cmd = m.search.Update(msg)
		// return m, cmd
	}

	return m, nil
}
