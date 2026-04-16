// Package graphqlapi: wrapper around graphQL client
package graphqlapi

import (
	"context"
	"net/http"
	"sort"

	"leetui/src/lib/graphqlapi/models"

	"github.com/hasura/go-graphql-client"
)

var client *graphql.Client

func InitLeetcodeGraphQLClient() *graphql.Client {
	if client == nil {
		client = graphql.NewClient("https://leetcode.com/graphql", nil).
			WithRequestModifier(func(r *http.Request) {
				r.Header.Set("Referer", "https://leetcode.com/")
				r.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
				r.Header.Set("Content-Type", "application/json")
			})
	}
	return client
}

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

func GetProblems(page int, pagesize int, filter QuestionGetFilter) ([]models.Problem, error) {
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

	if err := client.Query(context.Background(), &query, vars); err != nil {
		return nil, err
	}

	result = append(result, query.QuestionList.Questions...)

	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })

	return result, nil
}
