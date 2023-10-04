package fullscreen

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func (state State) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (state State) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return state, tea.Quit
		}
	}

	return state, nil
}

func (state State) View() string {
	return fmt.Sprintf("\n\n   " + content.welcomeMessage)
}
