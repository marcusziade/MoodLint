package validator

import (
	"MoodLint/src/utils"
	"fmt"
	"os"
)

// Checks if the prefix is one of the allowed prefixes
func ValidatePrefix(prefix string) {
	allowedPrefixes := []string{"refactor", "chore", "feat", "fix", "regfix"}
	if !utils.Contains(allowedPrefixes, prefix) {
		fmt.Printf("Invalid prefix. Allowed prefixes are: %v\n", allowedPrefixes)
		os.Exit(1)
	}
}
