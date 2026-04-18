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

type ProblemDetails struct {
	ID               string        `graphql:"questionId" json:"questionId"`
	FrontendID       string        `graphql:"questionFrontendId" json:"questionFrontendId"`
	Title            string        `graphql:"title" json:"title"`
	TitleSlug        string        `graphql:"titleSlug" json:"titleSlug"`
	Content          string        `graphql:"content" json:"content"`
	Difficulty       string        `graphql:"difficulty" json:"difficulty"`
	Likes            int           `graphql:"likes" json:"likes"`
	Dislikes         int           `graphql:"dislikes" json:"dislikes"`
	ACRate           float64       `graphql:"acRate" json:"acRate"`
	IsPaidOnly       bool          `graphql:"isPaidOnly" json:"isPaidOnly"`
	CategoryTitle    string        `graphql:"categoryTitle" json:"categoryTitle"`
	TopicTags        []TopicTag    `graphql:"topicTags" json:"topicTags"`
	CodeSnippets     []CodeSnippet `graphql:"codeSnippets" json:"codeSnippets"`
	Hints            []string      `graphql:"hints" json:"hints"`
	SimilarQuestions string        `graphql:"similarQuestions" json:"similarQuestions"`
	Stats            string        `graphql:"stats" json:"stats"`
	SampleTestCase   string        `graphql:"sampleTestCase" json:"sampleTestCase"`
	MetaData         string        `graphql:"metaData" json:"metaData"`
	ExampleTestcases string        `graphql:"exampleTestcases" json:"exampleTestcases"`
}

type TopicTag struct {
	Name string `graphql:"name" json:"name"`
	Slug string `graphql:"slug" json:"slug"`
}

type CodeSnippet struct {
	Lang     string `graphql:"lang" json:"lang"`
	LangSlug string `graphql:"langSlug" json:"langSlug"`
	Code     string `graphql:"code" json:"code"`
}
