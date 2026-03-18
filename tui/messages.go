package tui

import "jplayer/fs"

type dirLoadedMsg struct {
	fs.DirectoryContents
}

type mpvEventMsg struct {
	data string
}

type mpvCheckedMsg struct {
	installed bool
}

type errMsg struct {
	err error
}

type tickMsg struct{}
