package fs

import (
	"jplayer/model"
	"os"
	"path/filepath"
	"strings"
)

type DirectoryContents struct {
	Tracks      []model.Track
	Directories []string
}

var audioExtensions = map[string]bool{
	".mp3":  true,
	".flac": true,
	".wav":  true,
	".ogg":  true,
	".aac":  true,
	".m4a":  true,
}

// GoTo navigate to the next or previous directory and returns
// the new directory stack
func GoTo(nextDir string, dirStack []string) []string {
	if nextDir == ".." {
		if len(dirStack) > 1 {
			return dirStack[:len(dirStack)-1]
		}

		return dirStack
	}

	return append(dirStack, nextDir)
}

// GetDirectoryContents returns all subdirectories and audio tracks
// found at the path represented by dirStack.
func GetDirectoryContents(dirStack []string) (DirectoryContents, error) {
	var directoryNames []string
	var tracks []model.Track

	path := filepath.Join(dirStack...)
	files, err := os.ReadDir(path)
	if err != nil {
		return DirectoryContents{}, err
	}

	for _, file := range files {
		// Check for dirs and hide the hidden ones
		if file.Type().IsDir() && !strings.HasPrefix(file.Name(), ".") {
			directoryNames = append(directoryNames, file.Name())
		}

		if audioExtensions[filepath.Ext(file.Name())] {
			tracks = append(tracks, model.Track{
				Title: file.Name(),
				Path:  filepath.Join(path, file.Name()),
			})
		}
	}

	return DirectoryContents{
		Tracks:      tracks,
		Directories: directoryNames,
	}, nil
}
