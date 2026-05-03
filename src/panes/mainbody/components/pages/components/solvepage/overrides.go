package solvepage

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"leetui/src/lib/common/layers/alerts"
	"leetui/src/lib/graphqlapi/models"

	tea "charm.land/bubbletea/v2"
)

var langExtensions = map[string]string{
	"cpp":        "cpp",
	"java":       "java",
	"python":     "py",
	"python3":    "py",
	"c":          "c",
	"csharp":     "cs",
	"javascript": "js",
	"typescript": "ts",
	"php":        "php",
	"swift":      "swift",
	"kotlin":     "kt",
	"dart":       "dart",
	"golang":     "go",
	"ruby":       "rb",
	"scala":      "scala",
	"rust":       "rs",
	"racket":     "rkt",
	"erlang":     "erl",
	"elixir":     "ex",
	"bash":       "sh",
	"mysql":      "sql",
	"mssql":      "sql",
	"oraclesql":  "sql",
}

func (m SolvePageModel) CloneProblemAndOpenEditor(dir string, language string, problem models.ProblemDetails) (tea.Model, tea.Cmd) {
	if dir == "" || language == "" {
		return m, alerts.AlertErrorCmd(alerts.AlertErrorMsg{
			Duration: time.Second * 5,
			Message:  "please specify directory and lanugage",
		})
	}

	problemDir := filepath.Join(dir, problem.TitleSlug)
	os.MkdirAll(problemDir, 0o755)

	// Find the code snippet for the selected language
	var snippet string
	var langSlug string
	for _, cs := range problem.CodeSnippets {
		if cs.Lang == language {
			snippet = cs.Code
			langSlug = cs.LangSlug
			break
		}
	}

	// Create the source file with the snippet
	ext := langExtensions[langSlug]
	if ext == "" {
		ext = langSlug
	}
	mainFile := filepath.Join(problemDir, "main."+ext)
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		os.WriteFile(mainFile, []byte(snippet), 0o644)
	}

	terminal := os.Getenv("TERMINAL")
	if terminal == "" {
		terminal = "x-terminal-emulator"
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(terminal, "-e", editor, ".")
	cmd.Dir = problemDir
	cmd.Start()

	return m, nil
}
