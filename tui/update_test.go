package tui

import (
	"reflect"
	"testing"
)

func TestGetActivePanel(t *testing.T) {
	myApp := app{
		activePanel: position{col: 1, row: 1},
		panels: map[position]panel{
			{col: 1, row: 1}: {cursor: 0},
			{col: 2, row: 1}: {cursor: 5},
		},
	}

	expected := panel{cursor: 0}
	got := myApp.getActivePanel()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("getActivePanel() = %v, want %v", got, expected)
	}
}

func TestUpdateCursor(t *testing.T) {
	tests := []struct {
		name          string
		initialApp    app
		newCursor     int
		expectedPanel panel
	}{
		{
			name: "update cursor in first panel",
			initialApp: app{
				activePanel: position{col: 1, row: 1},
				panels: map[position]panel{
					{col: 1, row: 1}: {cursor: 0},
					{col: 2, row: 1}: {cursor: 5},
				},
			},
			newCursor:     10,
			expectedPanel: panel{cursor: 10},
		},
		{
			name: "update cursor in secondary panel",
			initialApp: app{
				activePanel: position{col: 2, row: 1},
				panels: map[position]panel{
					{col: 1, row: 1}: {cursor: 0},
					{col: 2, row: 1}: {cursor: 5},
				},
			},
			newCursor:     42,
			expectedPanel: panel{cursor: 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Work on a copy of the initial app
			a := tt.initialApp

			a.updateCursor(tt.newCursor)

			// 1. Check the active panel specifically
			got := a.getActivePanel()
			if got.cursor != tt.newCursor {
				t.Errorf("After updateCursor(%d), got cursor %d, want %d",
					tt.newCursor, got.cursor, tt.newCursor)
			}

			// 2. Double check the internal slice was actually updated
			if a.panels[a.activePanel].cursor != tt.newCursor {
				t.Errorf("Internal slice not updated: got %d, want %d",
					a.panels[a.activePanel].cursor, tt.newCursor)
			}
		})
	}
}
