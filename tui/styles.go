package tui

import (
	"path/filepath"

	"charm.land/lipgloss/v2"
)

func renderHeader(dirStack []string, width int) string {
	borderColor := lipgloss.Color("60")

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(width).
		Padding(0, 1)

	return style.Render(filepath.Join(dirStack...))
}
