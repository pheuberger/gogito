package cmd

import (
	"github.com/pheuberger/gogito/internal/subcommands"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new git repository",
	Long: `
description:
    This command creates an empty Git repository - basically a .git
    directory with subdirectories for objects, refs/heads, refs/tags, and
    template files. An initial branch without any commits will be created. 

    Running gogito init in an existing repository is safe. It will not overwrite
    things that are already there.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			subcommands.Init(".")
		case 1:
			subcommands.Init(args[0])
		default:
			cmd.Usage()
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.SetUsageTemplate("usage: gogito init [<directory>]\n")
}
