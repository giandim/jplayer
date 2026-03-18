package tui

import (
	"path/filepath"

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

	s += filepath.Join(a.dirStack...) + "\n\n"

	for index, dir := range a.nextDirectories {
		if a.panels[a.activePanel].cursor == index {
			s += ">"
		} else {
			s += " "
		}
		s += dir + "\n"
	}

	return tea.NewView(s)
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// user actions
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		case "j", "down":
			p := a.panels[a.activePanel]
			if p.cursor < len(a.nextDirectories)-1 {
				p.cursor++
				a.panels[a.activePanel] = p
			}
			return a, nil
		case "k", "up":
			p := a.panels[a.activePanel]
			if p.cursor > 0 {
				p.cursor--
				a.panels[a.activePanel] = p
			}
			return a, nil

		}

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
