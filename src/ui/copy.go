package ui

type filePicker struct {
	title             string
	fileNotValidError string
}

type copy struct {
	filePicker filePicker
}

var content = copy{
	filePicker{
		title:             "Pick a csv file to operate : ",
		fileNotValidError: " is of invalid type. Please select a valid file.",
	},
}
