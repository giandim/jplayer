package tui

import (
	"fmt"
	"jplayer/fs"

	tea "charm.land/bubbletea/v2"
)

func (a app) Init() tea.Cmd {
	return tea.Batch(
		checkMpvCmd(),
		loadDirCmd(a.dirStack),
	)
}

func (a app) View() tea.View {
	s := ""
	if !a.mpvInstalled {
		s = "Mpv is not installed but required to run jplayer"
		return tea.NewView(s)
	}

	s += renderHeader(a.dirStack, a.windowWidth) + "\n"

	for index, dir := range a.nextDirectories {
		if a.panels[a.activePanel].cursor == index {
			s += fmt.Sprintf(">  %v \n", dir)
		} else {
			s += fmt.Sprintf("  %v \n", dir)
		}
	}

	s += "\n\n\n"

	for _, track := range a.tracks {
		s += track.Title + "\n"
	}

	v := tea.NewView(s)
	v.AltScreen = true
	v.WindowTitle = "jplayer"
	return v
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// user actions
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		case "j", "down":
			p := a.getActivePanel()
			if p.cursor < len(a.nextDirectories)-1 {
				p.cursor++
				a.panels[a.activePanel] = p
			}
			return a, nil

		case "k", "up":
			p := a.getActivePanel()
			if p.cursor > 0 {
				p.cursor--
				a.panels[a.activePanel] = p
			}
			return a, nil

		case "enter":
			p := a.getActivePanel()
			nextDir := a.nextDirectories[p.cursor]
			a.dirStack = fs.GoTo(nextDir, a.dirStack)
			a.updateCursor(0)
			return a, loadDirCmd(a.dirStack)

		case "backspace":
			a.dirStack = fs.GoTo("..", a.dirStack)
			a.updateCursor(0)
			return a, loadDirCmd(a.dirStack)

		}

	case tea.WindowSizeMsg:
		a.windowWidth = msg.Width
		a.windowHeight = msg.Height
		return a, nil

	case dirLoadedMsg:
		a.tracks = msg.Tracks
		a.nextDirectories = msg.Directories
		return a, nil
	case mpvEventMsg:
		return a, nil
	case mpvCheckedMsg:
		a.mpvInstalled = msg.installed
		return a, nil
	case tickMsg:
		return a, nil
	case errMsg:
		a.error = msg.err
		return a, nil
	}

	return a, nil
}
