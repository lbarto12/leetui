package sidebar

import "github.com/lbarto12/leetui/src/panes/sidebar/focus"

func (m *SidebarModel) SetFocused(f bool) { // @override
	m.Focused = f
	if m.focusedChild == focus.SearchBar {
		m.search.Focus()
	}
}

func (m *SidebarModel) ToggleCollapse() {
	m.collapsed = !m.collapsed
}

func (m *SidebarModel) IsCollapsed() bool {
	return m.collapsed
}
