package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spark8899/gowallet/internal/commonPrivateKey"
	"github.com/spf13/cobra"
)

var number int
var privateKey string

var genPrivateKeyCmd = &cobra.Command{
	Use:     "genPrivateKey",
	Short:   "Generate a new random private key",
	Long:    "Generate a secure, random private key using crypto/rand.",
	Example: `  gowallet genPrivateKey -n 5`,
	Run: func(cmd *cobra.Command, args []string) {
		commonPrivateKey.GetGenerateKey(number)
	},
}

var getAddressCmd = &cobra.Command{
	Use:     "getAddress",
	Short:   "Derive a wallet address from a private key",
	Long:    "Derive a wallet address from a given private key.",
	Example: `  gowallet getAddress -k <private_key_hex>`,
	Run: func(cmd *cobra.Command, args []string) {
		if privateKey == "" {
			fmt.Println("Error: Private key is required. Use -k flag.")
			os.Exit(1)
		}
		address, err := commonPrivateKey.AddressHex(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(address)
	},
}

var getPublicKeyCmd = &cobra.Command{
	Use:     "getPublicKey",
	Short:   "Derive a public key from a private key",
	Long:    "Derive a public key from a given private key.",
	Example: `  gowallet getPublicKey -k <private_key_hex>`,
	Run: func(cmd *cobra.Command, args []string) {
		if privateKey == "" {
			fmt.Println("Error: Private key is required. Use -k flag.")
			os.Exit(1)
		}
		publicKey, err := commonPrivateKey.PublicKeyHex(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(publicKey)
	},
}

func init() {
	genPrivateKeyCmd.Flags().IntVarP(&number, "num", "n", 1, "generated quantity")
	getAddressCmd.Flags().StringVarP(&privateKey, "key", "k", "", "private key")
	getPublicKeyCmd.Flags().StringVarP(&privateKey, "key", "k", "", "private key")
}
