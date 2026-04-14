// Package state: Defines general app state. Sobmodules may include state for stateful components
package state

import (
	"leetui/src/panes/mainbody"
	"leetui/src/panes/sidebar"
)

type AppState struct {
	focused  string
	mainbody mainbody.MainBodyModel
	sidebar  sidebar.SidebarModel

	width  int
	height int
}

func MakeAppState() AppState {
	return AppState{
		mainbody: mainbody.MakeMainBodyModel(),
		sidebar:  sidebar.MakeSidebarModel(),
		focused:  "main",
	}
}
