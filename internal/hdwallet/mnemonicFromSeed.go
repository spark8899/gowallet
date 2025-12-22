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
		return "", fmt.Errorf("failed to decode seed string: %w", err)
	}

	// Convert seed to master private key
	masterKey01, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", fmt.Errorf("failed to create master key: %w", err)
	}

	// Convert entropy to mnemonic
	mnemonic, err := bip39.NewMnemonic(masterKey01.Key)
	if err != nil {
		return "", fmt.Errorf("failed to create mnemonic: %w", err)
	}

	fmt.Println("Mnemonic:", mnemonic)
	return mnemonic, nil
}
