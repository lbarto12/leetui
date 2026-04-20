// Package components: components for sidebar pane
package components

import "charm.land/bubbles/v2/textinput"

func ProblemSearchBar() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Find problems..."
	ti.Prompt = " "
	ti.CharLimit = 128
	ti.Focus() // as initial

	return ti
}
