package ui

import (
	"github.com/YashKumarVerma/csvtool/src/ui/fullscreen"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func Init() {
	newWelcomeScreen()
}

func newWelcomeScreen() {
	handler := tea.NewProgram(fullscreen.State{}, tea.WithAltScreen())
	if _, err := handler.Run(); err != nil {
		log.Fatal(err)
	}
}
