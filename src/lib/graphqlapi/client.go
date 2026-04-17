// Package graphqlapi: wrapper around graphQL client
package graphqlapi

import (
	"net/http"

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
