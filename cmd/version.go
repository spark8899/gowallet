package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gowallet",
	Long:  `All software has versions. This is gowallet's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gowallet v0.9 -- HEAD")
	},
}
