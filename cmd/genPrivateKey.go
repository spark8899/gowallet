package cmd

import (
	"fmt"
	"os"

	"github.com/spark8899/gowallet/internal/commonPrivateKey"
	"github.com/spf13/cobra"
)

var number int
var privateKey string

var genPrivateKeyCmd = &cobra.Command{
	Use:   "genPrivateKey",
	Short: "create generate key",
	Long:  "this command is create enerate key.",
	Run: func(cmd *cobra.Command, args []string) {
		commonPrivateKey.GetGenerateKey(number)
	},
}

var getEthAddressCmd = &cobra.Command{
	Use:   "getEthAddress",
	Short: "input privateKey and get address",
	Long:  "this command is get address from privateKey.",
	Run: func(cmd *cobra.Command, args []string) {
		if privateKey == "" {
			fmt.Println("Enter private key is empty")
			os.Exit(5)
		}
		commonPrivateKey.GetEthAddress(privateKey)
	},
}

var getPublicKeyCmd = &cobra.Command{
	Use:   "getPublicKey",
	Short: "input privateKey and get PublicKey",
	Long:  "this command is get public key from privateKey.",
	Run: func(cmd *cobra.Command, args []string) {
		if privateKey == "" {
			fmt.Println("Enter private key is empty")
			os.Exit(5)
		}
		commonPrivateKey.GetPublicKey(privateKey)
	},
}

func init() {
	genPrivateKeyCmd.Flags().IntVarP(&number, "num", "n", 1, "generated quantity")
	getEthAddressCmd.Flags().StringVarP(&privateKey, "key", "k", "", "private key")
	getPublicKeyCmd.Flags().StringVarP(&privateKey, "key", "k", "", "private key")
}
