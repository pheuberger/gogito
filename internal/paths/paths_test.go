package paths

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAbsFrom(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get current working directory: %v", err)
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("could not get user home directory: %v", err)
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"/absolute/path", "/absolute/path"},
		{"relative/path", filepath.Join(workingDir, "relative/path")},
		{"", workingDir},
		{"./", workingDir},
		{".", workingDir},
		{"~", userHomeDir},
		{"~/test", filepath.Join(userHomeDir, "test")},
	}

	for _, test := range tests {
		result := AbsFrom(test.input)
		if result != test.expected {
			t.Errorf("AbsFrom(%q) = expected %q; got %q", test.input, test.expected, result)
		}
	}
}
