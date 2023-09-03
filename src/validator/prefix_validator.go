package validator

import (
	"MoodLint/src/utils"
	"fmt"
)

// Checks if the prefix is one of the allowed prefixes
func ValidatePrefix(prefix string) (bool, string) {
	allowedPrefixes := []string{"refactor", "chore", "feat", "fix", "regfix"}
	if !utils.Contains(allowedPrefixes, prefix) {
		return false, fmt.Sprintf("Invalid prefix. Allowed prefixes are: %v\n", allowedPrefixes)
	}
	return true, ""
}
