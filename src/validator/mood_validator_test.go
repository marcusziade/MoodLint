package validator

import (
	"testing"
)

func TestCheckImperativeMood(t *testing.T) {
	tests := []struct {
		prefix       string
		description  string
		shouldSkip   bool
		isImperative bool
	}{
		{"refactor", "Refactor code", true, false},
		{"fix", "Fix bug", true, false},
		{"feat", "Add a feature", false, true},
		{"chore", "Adds chore", false, false},
	}

	for _, tt := range tests {
		shouldSkip, isImperative := checkImperativeMood(tt.prefix, tt.description)
		if shouldSkip != tt.shouldSkip {
			t.Errorf("Expected skip for prefix %s but got otherwise", tt.prefix)
		}

		if isImperative != tt.isImperative {
			t.Errorf("For description %s, got imperative as %v but want %v", tt.description, isImperative, tt.isImperative)
		}
	}
}
