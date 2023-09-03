package validator

import (
	"MoodLint/src/utils"
	"fmt"
	"os"

	"github.com/jdkato/prose/v2"
)

// Checks the mood and prints an output based on the result.
func ValidateImperativeMood(prefix string, description string) {
	shouldSkip, isImperative := checkImperativeMood(prefix, description)
	if shouldSkip {
		return
	}

	if !isImperative {
		printImperativeMoodMessage()
		os.Exit(1)
	}
}

// Checks the mood and returns two boolean values.
func checkImperativeMood(prefix string, description string) (bool, bool) {
	skippablePrefixes := []string{"refactor", "fix", "regfix"}
	if utils.Contains(skippablePrefixes, prefix) {
		return true, false
	}
	return false, IsImperativeMood(description)
}

// Contains the logic to check for imperative mood.
func IsImperativeMood(description string) bool {
	doc, _ := prose.NewDocument(description)
	tokens := doc.Tokens()
	const imperativeMoodTag = "VB"

	return len(tokens) > 0 && tokens[0].Tag == imperativeMoodTag
}

// Prints the message guiding the user to use the imperative mood.
func printImperativeMoodMessage() {
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Ensure your PR title is phrased in the imperative mood. Think of the phrase 'This PR will...'")

	fmt.Println("Good examples include:	")
	fmt.Println("\t- feat: Add a new feature")
	fmt.Println("\t- chore: Update the README")
	fmt.Println("\t- chore: Remove deprecated methods")
	fmt.Println("\t- feat: [domain-example] Add a new feature")
	fmt.Println("\t- fix: [domain-example] Fix a bug")
	fmt.Println("\t- regfix: [domain-example] Fix a bug")

	fmt.Println("")

	fmt.Println("Bad examples include:")
	fmt.Println("\t- feat: Added a new feature")
	fmt.Println("\t- chore: Updating the README")
	fmt.Println("\t- chore: Removing deprecated methods")

	fmt.Println("")
	fmt.Println("")
}
