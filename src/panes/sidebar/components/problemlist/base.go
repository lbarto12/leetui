// Package problemlist: Recyclerview for problems
package problemlist

import (
	"context"
	"fmt"
	"log"
	"strings"

	"leetui/src/lib/graphqlapi"
	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ProblemlistViewModel struct {
	viewmodel.ViewModel
	problems       []models.Problem
	searchCancel   context.CancelFunc
	loading        bool
	loadingSpinner spinner.Model
}

func NewProblemListViewModel() *ProblemlistViewModel {
	problems, err := graphqlapi.GetProblemsRaw(context.Background(), 0, 25, graphqlapi.QuestionGetFilter{})
	if err != nil {
		log.Fatal(err)
	}

	problemLoadingSpinner := spinner.New()
	problemLoadingSpinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	problemLoadingSpinner.Spinner = spinner.Dot

	return &ProblemlistViewModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
			Dims: viewmodel.ViewModelDims{
				Width:  30,
				Height: 30,
			},
		},
		problems:       problems,
		loadingSpinner: problemLoadingSpinner,
	}
}

func (m ProblemlistViewModel) Init() tea.Cmd {
	return m.loadingSpinner.Tick
}

func (m ProblemlistViewModel) View() tea.View {
	innerWidth := max(m.Dims.Width, 10)

	var content string
	if m.loading {
		content = m.loadingSpinner.View() + " searching..."
	} else if len(m.problems) == 0 {
		content = "no results."
	} else {
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
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)

	case SearchQueryMsg:
		if m.searchCancel != nil {
			return m, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		m.loading = true
		m.searchCancel = cancel
		return m, graphqlapi.GetProblems(ctx, 0, 25, graphqlapi.QuestionGetFilter{
			SearchKeywords: msg.Query,
		})

	case graphqlapi.ProblemsLoadedMsg:
		m.loading = false
		m.searchCancel = nil
		if msg.Err == nil {
			m.problems = msg.Problems
		}
		return m, nil

	case spinner.TickMsg:
		if msg.ID == m.loadingSpinner.ID() {
			var cmd tea.Cmd
			m.loadingSpinner, cmd = m.loadingSpinner.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}
