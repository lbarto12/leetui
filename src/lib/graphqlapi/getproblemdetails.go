package graphqlapi

import (
	"context"

	"github.com/lbarto12/leetui/src/lib/graphqlapi/models"

	tea "charm.land/bubbletea/v2"
	"github.com/hasura/go-graphql-client"
)

func GetProblemDetailsRaw(ctx context.Context, problemID string) (*models.ProblemDetails, error) {
	var query struct {
		Details models.ProblemDetails `graphql:"question(titleSlug: $titleSlug)"`
	}

	vars := map[string]any{
		"titleSlug": graphql.String(problemID),
	}

	if err := client.Query(ctx, &query, vars); err != nil {
		return nil, err
	}

	return &query.Details, nil
}

type ProblemDetailsLoadedMsg struct {
	Details *models.ProblemDetails
	Err     error
}

func GetProblemDetails(ctx context.Context, problemID string) tea.Cmd {
	return func() tea.Msg {
		details, err := GetProblemDetailsRaw(ctx, problemID)
		return ProblemDetailsLoadedMsg{
			Details: details,
			Err:     err,
		}
	}
}
