package sidebar

func (m *SidebarModel) SetFocused(f bool) { // @override
	m.Focused = f
	m.search.Focus()
}

func (m *SidebarModel) ToggleCollapse() {
	m.collapsed = !m.collapsed
}

func (m *SidebarModel) IsCollapsed() bool {
	return m.collapsed
}
