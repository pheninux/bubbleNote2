package main

import (
	"bubbleNote2/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	sm := ui.StateManager{}
	if err := tea.NewProgram(sm, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}
}
