package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	// utils and validators are in the same package
)

// Fetches the title argument from the command line
func GetTitle() string {
	return os.Args[1]
}

// Removes square brackets from the title
func RemoveBrackets(title string) string {
	re := regexp.MustCompile(`\[[^\]]*\]`)
	return re.ReplaceAllString(title, "")
}

// Splits the title into a prefix and a description
func SplitTitle(cleanedTitle string) (string, string) {
	split := strings.SplitN(cleanedTitle, ": ", 2)
	if len(split) < 2 {
		fmt.Println("Incomplete PR title. Ensure you have both a prefix and a description.")
		os.Exit(1)
	}
	return split[0], split[1]
}

// Checks if a slice contains a specific string
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
