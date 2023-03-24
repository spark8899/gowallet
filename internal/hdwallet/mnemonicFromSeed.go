package hdwallet

import (
	"encoding/hex"
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func MnemonicFromSeed(seedStr string) (string, error) {
	// to seed bytes
	seed, err := hex.DecodeString(seedStr)
	if err != nil {
		fmt.Println("seed str to bytes error")
		return "seed str to bytes", err
	}

	// create BIO32 master key
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		fmt.Println("seed to masterKey error")
		return "seed to masterKey", err
	}

	// from master key to mnemonic
	mnemonic, err := bip39.NewMnemonic(masterKey.Key)
	if err != nil {
		fmt.Println("masterKey to mnemonic error")
		return "masterKey to mnemonic", err
	}

	//fmt.Println("Mnemonic:", mnemonic)
	return mnemonic, nil
}
