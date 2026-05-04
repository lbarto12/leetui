package problemlist

import (
	"context"

	"github.com/lbarto12/leetui/src/lib/common/cmds"
	"github.com/lbarto12/leetui/src/lib/graphqlapi"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

func (m ProblemlistViewModel) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height)
		return m, nil
	case cmds.SearchQueryMsg:
		// Cancel any in-flight search before starting a new one
		if m.searchCancel != nil {
			m.searchCancel()
		}

		ctx, cancel := context.WithCancel(context.Background())
		m.loading = true
		m.searchCancel = cancel
		return m, graphqlapi.GetProblems(ctx, 0, 300, graphqlapi.QuestionGetFilter{
			SearchKeywords: msg.Query,
		})

	case graphqlapi.ProblemsLoadedMsg:
		// Ignore cancelled requests - a newer search is in flight
		if msg.Err != nil {
			return m, nil
		}

		m.loading = false
		m.searchCancel = nil
		m.problems = msg.Problems

		ptabledata := ConvertProblemsToTableRows(m.problems)
		m.table.SetRows(ptabledata)
		m.table.GotoTop()

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
