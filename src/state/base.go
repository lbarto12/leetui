// Package state: Defines general app state. Sobmodules may include state for stateful components
package state

import (
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"
	"leetui/src/state/focus"
)

type AppState struct {
	focused  string
	mainbody mainbody.MainBodyModel
	sidebar  sidebar.SidebarModel

	width  int
	height int
}

func MakeAppState() AppState {
	as := AppState{
		mainbody: mainbody.MakeMainBodyModel(),
		sidebar:  sidebar.MakeSidebarModel(),
		focused:  focus.Main,
	}

	return as
}
