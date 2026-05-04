// Package keymap: shared interface for components to expose their active
// key bindings, plus helpers for combining bindings from a focus chain.
package keymap

import "charm.land/bubbles/v2/key"

// Provider is anything that can describe its currently-active bindings.
// Satisfies bubbles/v2/help.KeyMap, so any Provider can be passed straight
// to help.Model.View.
type Provider interface {
	ShortHelp() []key.Binding
	FullHelp() [][]key.Binding
}

// Merged combines several Providers into one. ShortHelp concatenates;
// FullHelp appends each Provider's columns in order. Nil entries are
// skipped so callers can build the chain without branching.
type Merged []Provider

func (m Merged) ShortHelp() []key.Binding {
	var out []key.Binding
	for _, p := range m {
		if p == nil {
			continue
		}
		out = append(out, p.ShortHelp()...)
	}
	return out
}

func (m Merged) FullHelp() [][]key.Binding {
	var out [][]key.Binding
	for _, p := range m {
		if p == nil {
			continue
		}
		out = append(out, p.FullHelp()...)
	}
	return out
}
