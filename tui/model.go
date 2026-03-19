package tui

import (
	"fmt"
	"jplayer/model"
	"os"
	"time"
)

// FAKE ENUMS
type PlaybackState int

const (
	StateStopped PlaybackState = iota
	StatePlaying
	StatePaused
)

type PanelType int

const (
	FoldersPanel PanelType = iota
	TracksPanel
	QueuePanel
)

// MODEL STRUCTS
type panel struct {
	panelType PanelType
	cursor    int
	visible   bool
}

type position struct {
	row, col int
}

// MODEL
type app struct {
	dirStack        []string
	tracks          []model.Track
	nextDirectories []string
	queue           []model.Track

	panels       map[position]panel
	activePanel  position
	windowWidth  int
	windowHeight int

	playerState     PlaybackState
	selectedSong    *int
	currentPlayback time.Duration

	error        error
	mpvInstalled bool
}

func InitialModel() app {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Cannot access user dir: %v", err)
	}

	p := make(map[position]panel)
	p[position{row: 1, col: 1}] = panel{cursor: 0}

	return app{
		dirStack:    []string{userHomeDir},
		activePanel: position{row: 1, col: 1},
		panels:      p,
	}
}
