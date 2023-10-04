package ui

type uiState struct {
	currentView string
}

var state = uiState{
	currentView: "welcome",
}
