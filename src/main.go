package main

import (
	"MoodLint/src/utils"
	"MoodLint/src/validator"
	"fmt"
	"os"
)

func main() {
	title := utils.GetTitle()
	cleanedTitle := utils.RemoveBrackets(title)
	prefix, description := utils.SplitTitle(cleanedTitle)

	_, err := validator.ValidatePrefix(prefix)
	if err != "" {
		fmt.Println(err)
		os.Exit(1)
	}
	validator.ValidateImperativeMood(prefix, description)
}
