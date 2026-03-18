package fs

import (
	"jplayer/model"
	"os"
	"path/filepath"
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

func GoTo(nextDir string, dirStack []string) []string {
	if nextDir == ".." {
		if len(dirStack) > 1 {
			return dirStack[:len(dirStack)-1]
		}
	}

	return append(dirStack, nextDir)
}

func GetDirectoryContents(dirStack []string) (DirectoryContents, error) {
	var directoryNames []string
	var tracks []model.Track

	path := filepath.Join(dirStack...)
	files, err := os.ReadDir(path)
	if err != nil {
		return DirectoryContents{}, err
	}

	for _, file := range files {
		if file.Type().IsDir() {
			directoryNames = append(directoryNames, file.Name())
		}

		if audioExtensions[filepath.Ext(file.Name())] {
			tracks = append(tracks, model.Track{
				Title: file.Name(),
				Path:  path + file.Name(),
			})
		}
	}

	return DirectoryContents{
		Tracks:      tracks,
		Directories: directoryNames,
	}, nil
}
