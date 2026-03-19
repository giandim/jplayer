package tui

func (a app) getActivePanel() panel {
	return a.panels[a.activePanel]
}

func (a app) updateCursor(cursor int) {
	p := a.getActivePanel()
	p.cursor = cursor
	a.panels[a.activePanel] = p
}
