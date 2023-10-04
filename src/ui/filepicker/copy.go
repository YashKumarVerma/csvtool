package filepicker

type filePicker struct {
	title             string
	fileNotValidError string
}

type contentHolder struct {
	filePicker filePicker
}

var content = contentHolder{
	filePicker{
		title:             "Pick a csv file to operate : ",
		fileNotValidError: " is of invalid type. Please select a valid file.",
	},
}
