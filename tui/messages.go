package tui

type dirLoadedMsg struct {
	tracks []Track
}

type mpvEventMsg struct {
	data string
}

type mpvCheckedMsg struct {
	installed bool
}

type tickMsg struct{}
