package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strings"
)

type mainPage struct {
}

var (
	with           int
	height         int
	selection      = []string{"New Note \n", "List of Notes"}
	selected       = 1
	selectionStyle = termenv.String("").Foreground(termenv.EnvColorProfile().Color("#36AFF2"))
	titleStyle     = lipgloss.NewStyle().Align(lipgloss.Center).Background(lipgloss.AdaptiveColor{
		Light: "#F4E063",
		Dark:  "#F4E063",
	}).Foreground(lipgloss.AdaptiveColor{
		Light: "#5229F2",
		Dark:  "#5229F2",
	})
	helpStyle     = titleStyle.Copy()
	wrapMenuStyle = lipgloss.NewStyle().Align(lipgloss.Left)
)

func (p mainPage) Init() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (p mainPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return &StateManager{}, tea.Quit
		case "tab":
			if selected > 1 {
				selected--
			} else {
				selected++
			}
		case "enter":
			switch selected {
			case 1:
				pageCursor = 2
			case 2:
				pageCursor = 3
			}
		}
	case tea.WindowSizeMsg:
		with, height = msg.Width, msg.Height
	}
	return &StateManager{}, nil
}

func (p mainPage) View() string {

	layout := lipgloss.JoinVertical(lipgloss.Center, titleView(), wrapMenuStyle.Render(contentView()), helpView())
	return layout
}

// data for main page

func helpView() string {
	return helpStyle.Width(with).Render(
		"tab/change selection . ctrl+q/quit")
}

func titleView() string {
	return titleStyle.Render(
		"Acceuil")
}
func contentView() string {
	menu := strings.Builder{}
	choice := ""
	// block note choices : choice tpl
	for i, v := range selection {
		cursor := " "
		if selected == i+1 {
			cursor = ">"
			//choice = termenv.String(v).Foreground(termenv.EnvColorProfile().Color("121")).String()
			choice = selectionStyle.Styled(v)
		} else {
			choice = v
		}
		menu.WriteString(fmt.Sprintf("%s %s", cursor, choice))
	}
	return lipgloss.Place(with, height-(h(titleView())+h(helpView())), lipgloss.Center, lipgloss.Center, wrapMenuStyle.Render(menu.String()))

}
