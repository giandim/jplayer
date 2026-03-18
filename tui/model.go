package tui

import (
	"fmt"
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
type Track struct {
	title  string
	path   string
	album  string
	artist string
	lenght time.Duration
}

type panel struct {
	panelType PanelType
	cursor    int
	visible   bool
}

type Position struct {
	x, y int
}

// MODEL
type model struct {
	dirStack        []string
	tracks          []Track
	nextDirectories []string
	queue           []Track
	panels          map[Position]panel
	activePanel     Position
	windowWidth     int
	windowHeight    int
	playerState     PlaybackState
	selectedSong    *int
	currentPlayback time.Duration
}

func InitialModel() model {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("Cannot access user dir", err)
	}

	return model{
		dirStack: []string{userHomeDir},
	}
}
