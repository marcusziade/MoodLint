package main

import (
	"MoodLint/src/utils"
	"MoodLint/src/validator"
)

func main() {
	title := utils.GetTitle()
	cleanedTitle := utils.RemoveBrackets(title)
	prefix, description := utils.SplitTitle(cleanedTitle)

	validator.ValidatePrefix(prefix)
	validator.ValidateImperativeMood(prefix, description)
}
