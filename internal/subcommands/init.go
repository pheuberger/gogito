package subcommands

import (
	"fmt"
	"os"

	"github.com/pheuberger/gogito/internal/file"
	"github.com/pheuberger/gogito/internal/paths"
	"github.com/pheuberger/gogito/internal/repo"
)

const DESCRIPTION_TEXT = "Unnamed repository; edit this file 'description' to name the repository.\n"

// Straight up returns errors wihout polishing them.
// Since this is a learning project, this is fine.
func Init(pathToWorkingDir string) error {
	if repo.IsGitRepo(pathToWorkingDir) {
		fmt.Printf("%s is already a git repository. nothing to do\n", paths.AbsFrom(pathToWorkingDir))
		return nil
	}
	if err := createGitDirectory(pathToWorkingDir); err != nil {
		return err
	}
	if err := repo.WriteDefaultConfig(paths.GitDir(pathToWorkingDir)); err != nil {
		return err
	}

	// It's safe to instantiate a repo object since we just wrote the config.
	// Also, not expecting an error here because we just created the config
	// ourselves and know it to be sound. So ignore.
	repository, _ := repo.From(pathToWorkingDir)
	if err := createDirs(repository); err != nil {
		return err
	}
	if err := createFiles(repository); err != nil {
		return err
	}
	return nil
}

func createGitDirectory(pathToWorkingDir string) error {
	pathToGitDir := paths.GitDir(pathToWorkingDir)
	// mode 0777 before umask
	return os.Mkdir(pathToGitDir, 0777)
}

func createDirs(repository repo.Repo) error {
	dirs := [][]string{
		{"objects"},
		{"objects", "info"},
		{"objects", "pack"},
		{"branches"},
		{"info"},
		{"refs"},
		{"refs", "heads"},
		{"refs", "tags"},
		{"hooks"},
	}

	for _, dir := range dirs {
		if err := repository.EnsureDirs(dir...); err != nil {
			return fmt.Errorf("failed to create internal directory %v: %w", dir, err)
		}
	}
	return nil
}

func createFiles(repo repo.Repo) error {
	files := map[string]string{
		"description":  DESCRIPTION_TEXT,
		"HEAD":         "ref: refs/heads/main\n",
		"info/exclude": "# exclude file\n",
	}

	for path, content := range files {
		if err := file.Write(repo.Path(path), content); err != nil {
			return fmt.Errorf("failed to write internal file %s: %w", path, err)
		}
	}
	return nil
}
