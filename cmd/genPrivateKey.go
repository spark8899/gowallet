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
	Use:   "genPrivateKey",
	Short: "create generate key",
	Long:  "this command is create enerate key.",
	Run: func(cmd *cobra.Command, args []string) {
		commonPrivateKey.GetGenerateKey(number)
	},
}

var getAddressCmd = &cobra.Command{
	Use:   "getAddress",
	Short: "input privateKey and get address",
	Long:  "this command is get address from privateKey.",
	Run: func(cmd *cobra.Command, args []string) {
		if privateKey == "" {
			fmt.Println("Enter private key is empty")
			os.Exit(5)
		}
		address, err := commonPrivateKey.AddressHex(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(address)
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
