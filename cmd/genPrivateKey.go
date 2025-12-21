package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spark8899/gowallet/internal/commonPrivateKey"
	"github.com/spf13/cobra"
)

var number int
var privateKey string

var genPrivateKeyCmd = &cobra.Command{
	Use:     "genPrivateKey [count]",
	Short:   "Generate a new random private key",
	Long:    "Generate a secure, random private key using crypto/rand.",
	Example: `  gowallet genPrivateKey 5`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "help" {
				cmd.Help()
				os.Exit(0)
			}
			n, err := strconv.Atoi(args[0])
			if err == nil && n > 0 {
				number = n
			}
		}
		commonPrivateKey.GetGenerateKey(number)
	},
}

var getAddressCmd = &cobra.Command{
	Use:     "getAddress [private_key]",
	Short:   "Derive a wallet address from a private key",
	Long:    "Derive a wallet address from a given private key.",
	Example: `  gowallet getAddress <private_key_hex>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "help" {
				cmd.Help()
				os.Exit(0)
			}
			privateKey = args[0]
		}
		if privateKey == "" {
			fmt.Println("Error: Private key is required. Provide it as an argument or use -k flag.")
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
	Use:     "getPublicKey [private_key]",
	Short:   "Derive a public key from a private key",
	Long:    "Derive a public key from a given private key.",
	Example: `  gowallet getPublicKey <private_key_hex>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "help" {
				cmd.Help()
				os.Exit(0)
			}
			privateKey = args[0]
		}
		if privateKey == "" {
			fmt.Println("Error: Private key is required. Provide it as an argument or use -k flag.")
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
