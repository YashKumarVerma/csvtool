package main

import "github.com/YashKumarVerma/csvtool/src/ui"

func main() {
	ui.Init()

	//config := config.Init()
	//fp := filepicker.New()
	//fp.AllowedTypes = config.UI.SupportedFiletypes
	//
	//fp.CurrentDirectory, _ = os.UserHomeDir()
	//
	//m := filepicker2.Model{
	//	Filepicker: fp,
	//}
	//tm, _ := tea.NewProgram(&m, tea.WithOutput(os.Stderr)).Run()
	//mm := tm.(filepicker2.Model)
	//fmt.Println("\n  You selected: " + m.Filepicker.Styles.Selected.Render(mm.SelectedFile) + "\n")
}
