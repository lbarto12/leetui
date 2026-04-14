package main

import (
	"fmt"
	"os"

	"leetui/src/state"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(state.MakeAppState())

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
