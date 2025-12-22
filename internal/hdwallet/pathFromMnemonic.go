package hdwallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func PathFromMnemonic(mnemonic string, pathStr string) (string, error) {
	var err error
	if mnemonic == "" {
		return "", errors.New("mnemonic is required")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return "", errors.New("mnemonic is invalid")
	}

	// Validate derivation path format and security
	if err := ValidateDerivationPath(pathStr); err != nil {
		return "", fmt.Errorf("invalid derivation path: %w", err)
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", fmt.Errorf("failed to generate seed from mnemonic: %w", err)
	}

	path, err := accounts.ParseDerivationPath(pathStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse derivation path: %w", err)
	}

	// Create master private key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("failed to create master key: %w", err)
	}
	//fmt.Println("masterKey:", masterKey)

	// Derivation path
	fixIssue172 := true
	for _, n := range path {
		if fixIssue172 && masterKey.IsAffectedByIssue172() {
			masterKey, err = masterKey.Derive(n)
		} else {
			masterKey, err = masterKey.DeriveNonStandard(n)
		}
		if err != nil {
			return "", fmt.Errorf("failed to derive key at path %d: %w", n, err)
		}
	}

	//fmt.Println("masterKey:", masterKey)
	privateKey, err := masterKey.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return "", fmt.Errorf("failed to get EC private key: %w", err)
	}

	// Get Ethereum address from public key
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("unable to get public key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	info := fmt.Sprintf("%v:%v", address, hexutil.Encode(crypto.FromECDSA(privateKeyECDSA)))
	return info, nil
}
