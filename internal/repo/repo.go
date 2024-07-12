package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pheuberger/gogito/internal/paths"
)

type Repo struct {
	PathToWorkingDir string
	PathToGitDir     string
	Config           Config
}

// Initializes a new Repo struct from a given path.
//
// It will load .git/config and return the initialized repo. If the working
// directory is not a git repository, the function will return an error "not a
// git repository".
// If repositoryformatversion is not 0, the function will return an error as
// well since we're not equipped to deal with that repo type.
// See https://git-scm.com/docs/repository-version
func From(pathToWorkingDir string) (Repo, error) {
	config, err := readConfig(paths.GitDir(pathToWorkingDir))
	if err != nil {
		return Repo{}, err
	}
	if config.formatVersion != 0 {
		return Repo{}, fmt.Errorf("unsupported repositoryformatversion: %d", config.formatVersion)
	}
	return Repo{
		Config:           config,
		PathToWorkingDir: pathToWorkingDir,
		PathToGitDir:     paths.GitDir(pathToWorkingDir),
	}, nil
}

func IsGitRepo(pathToWorkingDir string) bool {
	_, err := os.Stat(paths.GitDir(pathToWorkingDir))
	return err == nil
}

func (repo Repo) Path(pathElements ...string) string {
	combined := append([]string{repo.PathToGitDir}, pathElements...)
	return filepath.Join(combined...)
}

func (repo Repo) EnsureDirs(pathElements ...string) error {
	path := repo.Path(pathElements...)
	return os.MkdirAll(path, 0777)
}
