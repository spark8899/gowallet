package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gowallet",
	Short: "a wallet tools",
	Long:  "gowallet is golang wallet tools.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(genPrivateKeyCmd)
	rootCmd.AddCommand(getEthAddressCmd)
	rootCmd.AddCommand(getPublicKeyCmd)
	rootCmd.AddCommand(genMnemonicCmd)
	rootCmd.AddCommand(getSeedCmd)
	rootCmd.AddCommand(getPathCmd)
	rootCmd.AddCommand(mnemonicToSeedCmd)
}
