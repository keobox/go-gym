package main

var AGlobalVariableFromAnotherFile = "from another file in the main package"

func AFunctionFromAnotherFile() bool {
	// It's exportable since starts with a capital letter
	return true
}

func stillVisibleEvenIfIsNotCapitalized() bool {
	// because is in package main
	return true
}
