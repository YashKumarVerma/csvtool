package ui

import "github.com/charmbracelet/bubbles/filepicker"

type Model struct {
	Filepicker   filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
}

type ClearErrorMsg struct{}
