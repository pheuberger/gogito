package repo

import (
	"errors"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	formatVersion    int
	fileMode         bool
	bare             bool
	logAllRefUpdates bool
}

const ConfigName = "config"

func WriteDefaultConfig(pathToGitDir string) error {
	return writeConfig(defaultConfig(), pathToGitDir)
}

// Read git config file in git directory specified by pathToGitDir.
//
// Returns error "not a git repository" if the config file is not found or a
// more specific error if the config could not be read.
func readConfig(pathToGitDir string) (Config, error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("ini")
	viper.AddConfigPath(pathToGitDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, errors.New("not a git repository")
		} else {
			// Config file was found but another error was produced
			return Config{}, err
		}
	}
	return Config{
		formatVersion:    viper.GetInt("core.repositoryformatversion"),
		fileMode:         viper.GetBool("core.filemode"),
		bare:             viper.GetBool("core.bare"),
		logAllRefUpdates: viper.GetBool("core.logallrefupdates"),
	}, nil
}

func writeConfig(config Config, pathToGitDir string) error {
	viper.SetConfigType("ini")
	viper.Set("core.repositoryformatversion", config.formatVersion)
	viper.Set("core.filemode", config.fileMode)
	viper.Set("core.bare", config.bare)
	viper.Set("core.logallrefupdates", config.logAllRefUpdates)
	return viper.WriteConfigAs(filepath.Join(pathToGitDir, ConfigName))
}

func defaultConfig() Config {
	return Config{
		formatVersion:    0,
		fileMode:         true,
		bare:             false,
		logAllRefUpdates: true,
	}
}
