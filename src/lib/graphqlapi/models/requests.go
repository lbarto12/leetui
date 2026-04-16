// Package models: models for leet graph requests
package models

type Problem struct {
	ID         string  `graphql:"questionId" json:"questionId"`
	Title      string  `graphql:"title"`
	TitleSlug  string  `graphql:"titleSlug"`
	Difficulty string  `graphql:"difficulty"`
	IsPaidOnly bool    `graphql:"isPaidOnly"`
	ACRate     float64 `graphql:"acRate"`
}
