package cmd

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/spark8899/gowallet/internal/hdwallet"
	"github.com/spf13/cobra"
)

var size int
var mnemonicStr string
var seedStr string
var path string

var genMnemonicCmd = &cobra.Command{
	Use:   "genMnemonic",
	Short: "generate mnemonic, protocol support: bip39",
	Long:  "generate mnemonic, protocol support: bip39",
	Run: func(cmd *cobra.Command, args []string) {
		bitSize := size*11 - size/3
		mnemonic, err := hdwallet.Bip39GenMnemonic(bitSize)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mnemonic)
	},
}

var getSeedCmd = &cobra.Command{
	Use:   "getSeed",
	Short: "get seed from mnemonic, protocol support: bip39",
	Long:  "get seed from mnemonic, protocol support: bip39",
	Run: func(cmd *cobra.Command, args []string) {
		if mnemonicStr == "" {
			fmt.Println("Enter mnemonic is empty")
			os.Exit(5)
		}
		seedBts, err := hdwallet.Bip39MnemonicToSeed(mnemonicStr, "")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(seedBts))
	},
}

var getPathCmd = &cobra.Command{
	Use:   "getPath",
	Short: "get address from hdwallet path",
	Long:  "get address from hdwallet path",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			fmt.Println("Enter path is empty")
			os.Exit(5)
		}
		if seedStr == "" && mnemonicStr == "" {
			fmt.Println("Enter seed and mnemonicS are empty")
			os.Exit(5)
		}

		if mnemonicStr != "" {
			KeyInfo, err := hdwallet.PathFromMnemonic(mnemonicStr, path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(KeyInfo)
		}

		if seedStr != "" {
			privateKeyInfo, err := hdwallet.PathFromSeed(seedStr, path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(privateKeyInfo)
		}
	},
}

var mnemonicToSeedCmd = &cobra.Command{
	Use:   "mnToSeed",
	Short: "seed convert to mnemonic",
	Long:  "seed convert to mnemonic",
	Run: func(cmd *cobra.Command, args []string) {
		if seedStr == "" {
			fmt.Println("Enter seed are empty")
			os.Exit(5)
		}

		mnemonicInfo, err := hdwallet.MnemonicFromSeed(seedStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mnemonicInfo)
	},
}

func init() {
	genMnemonicCmd.Flags().IntVarP(&size, "size", "s", 12, "size is the word number of mnemonic, support: 12, 15, 18, 21, 24")
	getSeedCmd.Flags().StringVarP(&mnemonicStr, "mnemonic", "m", "", "mnemonic is mnemonic string")
	getPathCmd.Flags().StringVarP(&seedStr, "seed", "s", "", "seed is string")
	getPathCmd.Flags().StringVarP(&path, "path", "p", "", "path is string, For example \"m/44'/60'/0'/0/0\"")
	getPathCmd.Flags().StringVarP(&mnemonicStr, "mnemonic", "m", "", "mnemonic is mnemonic string")
	mnemonicToSeedCmd.Flags().StringVarP(&seedStr, "seed", "s", "", "seed is string")
}
