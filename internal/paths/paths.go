package paths

import (
	"os"
	"path/filepath"
	"strings"
)

// Converts any path to an absolute path. If the path is already absolute it
// will return the path as is. If it's a relative path to the user's home
// directory ~ will be expanded. Otherwise the path will be appended to the
// current working directory.
//
// Panics if user home directory cannot be inferred.
//
// Examples:
//
//	input -> output
//	path/to/thing -> working/dir/path/to/thing
//	/absolute/path -> /absolute/path
//	~/foo/bar -> /home/username/foo/bar
func AbsFrom(path string) string {
	if strings.HasPrefix(path, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			return strings.Replace(path, "~", home, 1)
		}
		panic("could not get user home directory. is $HOME not set?")
	}
	result, _ := filepath.Abs(path)
	return result
}
