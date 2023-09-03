package utils

import (
	"os"
	"testing"
)

func TestGetTitle(t *testing.T) {
	os.Args = []string{"cmd", "Test: This is a test"}
	if got := GetTitle(); got != "Test: This is a test" {
		t.Errorf("GetTitle() = %v, want 'Test: This is a test'", got)
	}
}

func TestRemoveBrackets(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[WIP] Add: New feature", " Add: New feature"},
		{"Add: New feature", "Add: New feature"},
		{"[FIX][WIP] Bug fix", " Bug fix"},
	}

	for _, tt := range tests {
		if got := RemoveBrackets(tt.input); got != tt.want {
			t.Errorf("RemoveBrackets(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestSplitTitle(t *testing.T) {
	tests := []struct {
		input          string
		wantPrefix     string
		wantDescriptor string
	}{
		{"Add: New feature", "Add", "New feature"},
		{"Fix: Bug fix", "Fix", "Bug fix"},
		{"Remove: Old feature", "Remove", "Old feature"},
	}

	for _, tt := range tests {
		prefix, descriptor := SplitTitle(tt.input)
		if prefix != tt.wantPrefix || descriptor != tt.wantDescriptor {
			t.Errorf("SplitTitle(%v) = (%v, %v), want (%v, %v)", tt.input, prefix, descriptor, tt.wantPrefix, tt.wantDescriptor)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice []string
		val   string
		want  bool
	}{
		{[]string{"apple", "banana", "cherry"}, "banana", true},
		{[]string{"apple", "banana", "cherry"}, "grape", false},
		{[]string{}, "apple", false},
	}

	for _, tt := range tests {
		if got := Contains(tt.slice, tt.val); got != tt.want {
			t.Errorf("Contains(%v, %v) = %v, want %v", tt.slice, tt.val, got, tt.want)
		}
	}
}
