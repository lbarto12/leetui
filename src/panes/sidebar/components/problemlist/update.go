package problemlist

import (
	"context"

	"leetui/src/lib/graphqlapi"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

type SearchQueryMsg struct{ Query string }

func (m ProblemlistViewModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(m.GetSize().Width, msg.Height-4)
		m.table.SetWidth(m.GetSize().Width)
		m.table.SetHeight(m.GetSize().Height)
		return m, nil
	case SearchQueryMsg:
		if m.searchCancel != nil {
			return m, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		m.loading = true
		m.searchCancel = cancel
		return m, graphqlapi.GetProblems(ctx, 0, 300, graphqlapi.QuestionGetFilter{
			SearchKeywords: msg.Query,
		})

	case graphqlapi.ProblemsLoadedMsg:
		m.loading = false
		m.searchCancel = nil
		if msg.Err == nil {
			m.problems = msg.Problems
		}

		ptabledata := ConvertProblemsToTableRows(m.problems)
		m.table.SetRows(ptabledata)

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
