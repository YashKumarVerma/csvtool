package main

import (
	"fmt"
	"github.com/YashKumarVerma/csvtool/src/config"
	"os"

	"github.com/YashKumarVerma/csvtool/src/ui"
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config := config.Init()

	fp := filepicker.New()
	fp.AllowedTypes = config.UI.SupportedFiletypes

	fp.CurrentDirectory, _ = os.UserHomeDir()

	m := ui.Model{
		Filepicker: fp,
	}
	tm, _ := tea.NewProgram(&m, tea.WithOutput(os.Stderr)).Run()
	mm := tm.(ui.Model)
	fmt.Println("\n  You selected: " + m.Filepicker.Styles.Selected.Render(mm.SelectedFile) + "\n")
}
