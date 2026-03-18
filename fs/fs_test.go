package fs_test

import (
	"jplayer/fs"
	"jplayer/model"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGoTo(t *testing.T) {
	tests := []struct {
		name     string
		nextDir  string
		dirStack []string
		expected []string
	}{
		{
			name:     "navigate forward",
			nextDir:  "music",
			dirStack: []string{"/home", "user"},
			expected: []string{"/home", "user", "music"},
		},
		{
			name:     "navigate back",
			nextDir:  "..",
			dirStack: []string{"/home", "user"},
			expected: []string{"/home"},
		},
		{
			name:     "at root",
			nextDir:  "..",
			dirStack: []string{"/home"},
			expected: []string{"/home"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fs.GoTo(tt.nextDir, tt.dirStack)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func trackNames(tracks []model.Track) []string {
	names := make([]string, len(tracks))
	for i, t := range tracks {
		names[i] = filepath.Base(t.Path)
	}
	return names
}

func TestGetDirectoryContents(t *testing.T) {
	tests := []struct {
		name          string
		dirStack      []string
		setup         func(tmpDir string)
		expectedFiles []string
		expectedDirs  []string
		wantErr       bool
	}{
		{
			name: "mixed folders and audio files",
			setup: func(tmpDir string) {
				os.Mkdir(filepath.Join(tmpDir, "albums"), 0755)
				os.WriteFile(filepath.Join(tmpDir, "track.mp3"), []byte{}, 0644)
				os.WriteFile(filepath.Join(tmpDir, "song.flac"), []byte{}, 0644)
			},
			expectedFiles: []string{"song.flac", "track.mp3"},
			expectedDirs:  []string{"albums"},
		},
		{
			name:          "empty directory",
			setup:         func(tmpDir string) {},
			expectedFiles: nil,
			expectedDirs:  nil,
		},
		{
			name: "only folders",
			setup: func(tmpDir string) {
				os.Mkdir(filepath.Join(tmpDir, "jazz"), 0755)
				os.Mkdir(filepath.Join(tmpDir, "rock"), 0755)
			},
			expectedFiles: nil,
			expectedDirs:  []string{"jazz", "rock"},
		},
		{
			name: "only audio files",
			setup: func(tmpDir string) {
				os.WriteFile(filepath.Join(tmpDir, "track.mp3"), []byte{}, 0644)
				os.WriteFile(filepath.Join(tmpDir, "track.ogg"), []byte{}, 0644)
			},
			expectedFiles: []string{"track.mp3", "track.ogg"},
			expectedDirs:  nil,
		},
		{
			name:     "invalid path",
			setup:    func(tmpDir string) {},
			dirStack: []string{"/nonexistent/path"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tmpDir, err := os.MkdirTemp("", "jplayer-test-*")

			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}

			t.Cleanup(func() { os.RemoveAll(tmpDir) })

			tt.setup(tmpDir)

			dirStack := []string{tmpDir}
			if tt.dirStack != nil {
				dirStack = tt.dirStack
			}

			result, err := fs.GetDirectoryContents(dirStack)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			expectedTracks := make([]string, len(tt.expectedFiles))
			for i, f := range tt.expectedFiles {
				expectedTracks[i] = filepath.Join(tmpDir, f)
			}

			actualPaths := make([]string, len(result.Tracks))
			for i, track := range result.Tracks {
				actualPaths[i] = track.Path
			}

			if !reflect.DeepEqual(actualPaths, expectedTracks) {
				t.Errorf("expected tracks %v, got %v", expectedTracks, actualPaths)
			}

			if !reflect.DeepEqual(result.Directories, tt.expectedDirs) {
				t.Errorf("expected dirs %v, got %v", tt.expectedDirs, result.Directories)
			}
		})
	}
}
