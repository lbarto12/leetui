// Package problemlist: Recyclerview for problems
package problemlist

import (
	"context"
	"log"

	"leetui/src/lib/graphqlapi"
	"leetui/src/lib/graphqlapi/models"
	"leetui/src/lib/viewmodel"

	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type ProblemlistViewModel struct {
	viewmodel.ViewModel
	styles         Styles
	problems       []models.Problem
	searchCancel   context.CancelFunc
	loading        bool
	loadingSpinner spinner.Model
	table          table.Model
}

func MakeProblemListViewModel() ProblemlistViewModel {
	styles := MakeStyles()

	problems, err := graphqlapi.GetProblemsRaw(context.Background(), 0, 300, graphqlapi.QuestionGetFilter{})
	if err != nil {
		log.Fatal(err)
	}

	// Problem Table

	ptabledata := ConvertProblemsToTableRows(problems)

	ptable := table.New(
		table.WithColumns([]table.Column{
			{Title: "#", Width: 4},
			{Title: "Diff", Width: 6},
			{Title: "Title", Width: 20},
		}),
		table.WithRows(ptabledata),
		table.WithFocused(false),
		table.WithStyles(styles.tableStyles),
	)

	// Spinner
	problemLoadingSpinner := spinner.New()
	problemLoadingSpinner.Style = styles.spinnerStyle
	problemLoadingSpinner.Spinner = spinner.Dot

	return ProblemlistViewModel{
		ViewModel: viewmodel.ViewModel{
			Focused: false,
			Dims: viewmodel.ViewModelDims{
				Width:  30,
				Height: 30,
			},
		},
		problems:       problems,
		loadingSpinner: problemLoadingSpinner,
		table:          ptable,
		styles:         styles,
	}
}

func (m ProblemlistViewModel) Init() tea.Cmd {
	return m.loadingSpinner.Tick
}

func (m ProblemlistViewModel) View() tea.View {
	var content string
	if m.loading {
		content = m.loadingSpinner.View() + " searching..."
	} else if len(m.problems) == 0 {
		content = "no results."
	} else {
		idWidth := 4
		diffWidth := 6
		cellPadding := 2 * 3 // 1 char padding on each side per column, 3 columns
		titleWidth := max(m.Dims.Width-idWidth-diffWidth-cellPadding, 4)
		m.table.SetColumns([]table.Column{
			{Title: "#", Width: idWidth},
			{Title: "Diff", Width: diffWidth},
			{Title: "Title", Width: titleWidth},
		})
		m.table.SetWidth(m.Dims.Width)
		m.table.SetHeight(m.Dims.Height)
		content = m.table.View()
	}

	return tea.NewView(content)
}

func (m ProblemlistViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.HandleKeypress(msg)
	default:
		return m.HandleUpdate(msg)
	}
}

func (m *ProblemlistViewModel) GotoTop() {
	m.table.GotoTop()
}
