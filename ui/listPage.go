package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ListPage struct {
}

func (p ListPage) Init() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (p ListPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return &StateManager{}, tea.Quit
		case "ctrl+r":
			pageCursor = 1
		}
	}
	return &StateManager{}, nil
}

func (p ListPage) View() string {
	return "List notes"
}
