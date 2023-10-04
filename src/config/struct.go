package config

func Init() Global {
	ui := UI{
		SupportedFiletypes: []string{".csv", ".txt"},
	}

	return Global{
		UI: ui,
	}
}
