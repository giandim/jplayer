package tui

import (
	"jplayer/fs"
	"os/exec"

	tea "charm.land/bubbletea/v2"
)

func checkMpvCmd() tea.Cmd {
	return func() tea.Msg {
		_, err := exec.LookPath("mpv")
		return mpvCheckedMsg{installed: err == nil}
	}
}

func loadDirCmd(dirStack []string) tea.Cmd {
	return func() tea.Msg {
		dir, err := fs.GetDirectoryContents(dirStack)
		if err != nil {
			return errMsg{err: err}
		}

		return dirLoadedMsg{dir}
	}
}
