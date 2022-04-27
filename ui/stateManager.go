package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	pageCursor int = 1
	w              = lipgloss.Width
	h              = lipgloss.Height
)

const (
	MainP = iota + 1
	NoteP
	ListP
)

type StateManager struct {
	mainPage
	NotePage
	ListPage
}

func (sm StateManager) Init() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (sm StateManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch pageCursor {
	case 1:
		return sm.mainPage.Update(msg)
	case 2:
		return sm.NotePage.Update(msg)
	case 3:
		return sm.ListPage.Update(msg)
	default:
		return &StateManager{}, nil
	}

}

func (sm StateManager) View() string {
	switch pageCursor {
	case MainP:
		return sm.mainPage.View()
	case NoteP:
		return sm.NotePage.View()
	case ListP:
		return sm.ListPage.View()
	default:
		return ""
	}
}
