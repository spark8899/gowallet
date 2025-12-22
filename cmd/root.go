package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gowallet",
	Short: "A comprehensive cryptocurrency wallet tool",
	Long:  "gowallet is a CLI tool for generating keys, addresses, and mnemonics for various cryptocurrencies.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(genPrivateKeyCmd)
	rootCmd.AddCommand(getAddressCmd)
	rootCmd.AddCommand(getPublicKeyCmd)
	rootCmd.AddCommand(genMnemonicCmd)
	rootCmd.AddCommand(mnToSeedCmd)
	rootCmd.AddCommand(getPathCmd)
	rootCmd.AddCommand(seedToMnCmd)
}
