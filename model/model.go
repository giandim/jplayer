package model

import "time"

type Track struct {
	Title  string
	Path   string
	Album  string
	Artist string
	Lenght time.Duration
}
