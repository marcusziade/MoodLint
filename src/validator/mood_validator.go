package validator

import (
	"MoodLint/src/utils"
	"fmt"
	"os"

	"github.com/jdkato/prose/v2"
)

// Checks the mood of the PR description based on the prefix
func ValidateImperativeMood(prefix string, description string) {
	skippablePrefixes := []string{"refactor", "fix", "regfix"}
	if utils.Contains(skippablePrefixes, prefix) {
		fmt.Printf("Prefix '%s' detected. Skipping imperative mood check.\n", prefix)
		os.Exit(0)
	}

	checkImperativeMood(description)
}

// Checks if the description is phrased in the imperative mood and gives examples if not.
func checkImperativeMood(description string) {
	doc, _ := prose.NewDocument(description)
	tokens := doc.Tokens()
	const imperativeMoodTag = "VB"

	if len(tokens) > 0 && tokens[0].Tag != imperativeMoodTag {
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
		os.Exit(1)
	}
}
