package paths

import (
	"os"
	"path/filepath"
	"strings"
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

func FuzzAbsFrom(f *testing.F) {
	tests := []string{
		"/absolute/path",
		"relative/path",
		"",
		"./",
		".",
		"~",
		"~/test",
	}
	for _, test := range tests {
		f.Add(test)
	}
	f.Fuzz(func(t *testing.T, path string) {
		result := AbsFrom(path)
		if !strings.HasPrefix(result, "/") {
			t.Errorf("AbsFrom(%q) = expected absolute path; got %q", path, result)
		}
	})
}
