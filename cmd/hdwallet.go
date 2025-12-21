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
	Use:     "genMnemonic",
	Short:   "Generate a BIP39 mnemonic phrase",
	Long:    "Generate a BIP39 mnemonic phrase. Supported sizes: 12, 15, 18, 21, 24 words.",
	Example: `  gowallet genMnemonic -s 12`,
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
	Use:     "getSeed",
	Short:   "Convert a mnemonic phrase to a deterministic seed",
	Long:    "Convert a BIP39 mnemonic phrase to a deterministic seed (hex encoded).",
	Example: `  gowallet getSeed -m "apple banana ... "`,
	Run: func(cmd *cobra.Command, args []string) {
		if mnemonicStr == "" {
			fmt.Println("Error: Mnemonic is required. Use -m flag.")
			os.Exit(1)
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
	Short: "Derive keys/addresses from a derivation path",
	Long:  "Derive private key or address from a mnemonic or seed using a derivation path (e.g., m/44'/60'/0'/0/0).",
	Example: `  gowallet getPath -m "apple banana ..." -p "m/44'/60'/0'/0/0"
  gowallet getPath -s <seed_hex> -p "m/44'/60'/0'/0/0"`,
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			fmt.Println("Error: Path is required. Use -p flag.")
			os.Exit(1)
		}
		if seedStr == "" && mnemonicStr == "" {
			fmt.Println("Error: Either seed (-s) or mnemonic (-m) is required.")
			os.Exit(1)
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
	Use:     "mnToSeed",
	Short:   "Generate a mnemonic from a seed (entropy) hex string",
	Long:    "Generate a BIP39 mnemonic phrase from a provided seed/entropy hex string.",
	Example: `  gowallet mnToSeed -s <seed_hex>`,
	Run: func(cmd *cobra.Command, args []string) {
		if seedStr == "" {
			fmt.Println("Error: Seed is required. Use -s flag")
			os.Exit(1)
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
