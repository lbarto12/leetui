package graphqlapi

import (
	"context"
	"log"
	"sort"
	"strconv"

	"github.com/lbarto12/leetui/src/lib/graphqlapi/models"

	tea "charm.land/bubbletea/v2"
	"github.com/hasura/go-graphql-client"
)

// It's a limited number at the moment, pagination can come later

type QuestionGetFilter struct {
	SearchKeywords string   `json:"searchKeywords,omitempty"`
	Difficulty     string   `json:"difficulty,omitempty"` // "EASY", "MEDIUM", "HARD"
	Status         string   `json:"status,omitempty"`     // "AC", "NOT_STARTED", "TRIED"
	Tags           []string `json:"tags,omitempty"`       // e.g. ["array", "hash-table"]
	ListID         string   `json:"listId,omitempty"`
	PremiumOnly    bool     `json:"premiumOnly,omitempty"`
}

func (QuestionGetFilter) GetGraphQLType() string { return "QuestionListFilterInput" }

func GetProblemsRaw(ctx context.Context, page int, pagesize int, filter QuestionGetFilter) ([]models.Problem, error) {
	var result []models.Problem

	var query struct {
		QuestionList struct {
			Total     int              `graphql:"totalNum"`
			Questions []models.Problem `graphql:"data"`
		} `graphql:"questionList(categorySlug: $categorySlug, limit: $limit, skip: $skip, filters: $filters)"`
	}

	vars := map[string]any{
		"categorySlug": graphql.String(""),
		"limit":        graphql.Int(pagesize),
		"skip":         graphql.Int(page),
		"filters":      filter,
	}

	if err := client.Query(ctx, &query, vars); err != nil {
		return nil, err
	}

	result = append(result, query.QuestionList.Questions...)

	sort.Slice(result, func(i, j int) bool {
		iAsInt, err := strconv.ParseInt(result[i].ID, 10, 32)
		if err != nil {
			log.Fatal("unexpected integer type")
		}

		jAsInt, err := strconv.ParseInt(result[j].ID, 10, 32)
		if err != nil {
			log.Fatal("unexpected integer type")
		}

		return iAsInt < jAsInt
	})

	return result, nil
}

// Tea cmd for context

type ProblemsLoadedMsg struct {
	Problems []models.Problem
	Err      error
}

func GetProblems(ctx context.Context, page int, pagesize int, filter QuestionGetFilter) tea.Cmd {
	return func() tea.Msg {
		problems, err := GetProblemsRaw(ctx, page, pagesize, filter)
		return ProblemsLoadedMsg{Problems: problems, Err: err}
	}
}
