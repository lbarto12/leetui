package main

import (
	"fmt"
	"os"

	"leetui/src/app"
	"leetui/src/lib/graphqlapi"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	graphqlapi.InitLeetcodeGraphQLClient()

	p := tea.NewProgram(app.MakeAppState())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
