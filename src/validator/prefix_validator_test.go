package validator

import (
	"testing"
)

func TestValidatePrefix(t *testing.T) {
	tests := []struct {
		prefix  string
		isValid bool
	}{
		{"refactor", true},
		{"chore", true},
		{"feat", true},
		{"fix", true},
		{"regfix", true},
		{"invalid", false},
		{"test", false},
	}

	for _, tt := range tests {
		isValid, _ := ValidatePrefix(tt.prefix)
		if isValid != tt.isValid {
			t.Errorf("For prefix %s, got validity as %v, want %v", tt.prefix, isValid, tt.isValid)
		}
	}
}
