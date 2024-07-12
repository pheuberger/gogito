package subcommands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pheuberger/gogito/internal/paths"
)

func setupTempDir(t *testing.T) string {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "testrepo")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	return tempDir
}

func teardownTempDir(t *testing.T, tempDir string) {
	t.Helper()

	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Fatalf("Failed to remove temp dir: %v", err)
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		setupFunc  func(string) error
		name       string
		shouldFail bool
	}{
		{
			setupFunc: func(path string) error {
				return nil
			},
			name:       "successful init",
			shouldFail: false,
		},
		{
			setupFunc: func(path string) error {
				return os.Mkdir(filepath.Join(path, ".git"), 0755)
			},
			name:       "pre-existing git repo",
			shouldFail: false,
		},
		{
			setupFunc: func(path string) error {
				return os.Chmod(path, 0444) // read-only permissions
			},
			name:       "failed init due to write error",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := setupTempDir(t)
			defer teardownTempDir(t, tempDir)

			setupErr := tt.setupFunc(tempDir)
			if setupErr != nil {
				t.Fatalf("Test setup failed: %v", setupErr)
			}

			defer os.Chmod(tempDir, 0777) // reset permissions after test

			initErr := Init(tempDir)
			gitDir := paths.GitDir(tempDir)
			if tt.shouldFail {
				if initErr == nil {
					t.Fatalf("Expected Init to fail, but it succeeded")
				}
				if info, _ := os.Stat(gitDir); info != nil {
					t.Fatalf("Expected git directory to not exist, but it does")
				}
			} else {
				if initErr != nil {
					t.Fatalf("Expected Init to succeed, but it failed: %v", initErr)
				}
				if _, statErr := os.Stat(gitDir); os.IsNotExist(statErr) {
					t.Fatalf("Expected git directory to exist, but it does not")
				}

				if tt.name == "pre-existing git repo" {
					// test case for pre-existing git repo. no files were created, so skip
					return
				}
				expectedObjects := []string{
					"description",
					"HEAD",
					"info/exclude",
					"objects",
					"objects/info",
					"objects/pack",
					"branches",
					"info",
					"refs",
					"refs/heads",
					"refs/tags",
					"hooks",
				}
				for _, object := range expectedObjects {
					if _, statErr := os.Stat(filepath.Join(gitDir, object)); os.IsNotExist(statErr) {
						t.Fatalf("Expected filesystem object %s to exist, but it does not", object)
					}
				}
			}
		})
	}
}
