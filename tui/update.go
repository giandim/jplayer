package tui

// The following functions are mothods on a value receiver and not on a pointer receiver
// because the mutation needs to be performed in the main update func
func (a app) getActivePanel() panel {
	return a.panels[a.activePanel]
}

func (a app) updateCursor(cursor int) app {
	p := a.getActivePanel()
	p.cursor = cursor
	a.panels[a.activePanel] = p
	return a
}

func (a app) moveCursorUp() app {
	p := a.getActivePanel()
	if p.cursor > 0 {
		p.cursor--
		a.panels[a.activePanel] = p
	}
	return a
}

func (a app) moveCursorDown() app {
	p := a.getActivePanel()
	if p.cursor < a.activeItemsLen()-1 {
		p.cursor++
		a.panels[a.activePanel] = p
	}
	return a
}

func (a app) activeItemsLen() int {
	switch a.getActivePanel().panelType {
	case FoldersPanel:
		return len(a.nextDirectories)
	case TracksPanel:
		return len(a.tracks)
	}
	return 0
}

// Direction represents a navigation direction.
// Valid values are Up, Down, Left and Right.
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// moveCursor moves the cursor or switches the active panel based on the given direction.
// d must be one of the defined Direction constants: Up, Down, Left, Right.
func (a app) changePanel(d Direction) app {
	activePanel := a.activePanel
	switch d {
	case Up:
		activePanel.row--
	case Down:
		activePanel.row++
	case Left:
		activePanel.col--
	case Right:
		activePanel.col++
	}

	if _, ok := a.panels[activePanel]; ok {
		a.activePanel = activePanel
	}
	return a
}
