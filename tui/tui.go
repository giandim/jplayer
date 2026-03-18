package tui

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) Init() tea.Cmd {
	//tea.Batch(
	//	checkMpvCmd(),
	//	loadDirCmd(m.dirStack),
	//)
	return nil
}

func (m model) View() tea.View {
	return tea.NewView("hello")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// user actions
		return m, nil
	case dirLoadedMsg:
		m.tracks = msg.tracks
		return m, nil
	case mpvEventMsg:
		return m, nil
	case mpvCheckedMsg:
		return m, nil
	case tickMsg:
		return m, nil
	}

	return m, nil
}
