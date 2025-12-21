package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "v1.0.0"
	GitCommit = "HEAD"
	BuildTime = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gowallet",
	Long:  "Print the build version and git commit hash of gowallet.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nGit Commit: %s\nBuild Time: %s\n", Version, GitCommit, BuildTime)
	},
}
