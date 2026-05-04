package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lbarto12/leetui/src/app"
	"github.com/lbarto12/leetui/src/lib/graphqlapi"

	tea "charm.land/bubbletea/v2"
)

// Version is overridden at build time via -ldflags "-X main.Version=...".
var Version = "dev"

func main() {
	if os.Getenv("LEETUI_DEBUG") != "" {
		path, err := debugLogPath()
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not resolve debug log path:", err)
			os.Exit(1)
		}
		f, err := tea.LogToFile(path, "debug")
		if err != nil {
			fmt.Fprintln(os.Stderr, "fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	graphqlapi.InitLeetcodeGraphQLClient()

	p := tea.NewProgram(app.MakeAppState())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func debugLogPath() (string, error) {
	stateDir := os.Getenv("XDG_STATE_HOME")
	if stateDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		stateDir = filepath.Join(home, ".local", "state")
	}
	dir := filepath.Join(stateDir, "leetui")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "debug.log"), nil
}
