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
		fmt.Println("mnemonic is required")
		return "mnemonic is required", errors.New("mnemonic is required")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		fmt.Println("mnemonic is invalid")
		return "mnemonic is invalid", errors.New("mnemonic is invalid")
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		fmt.Println("mnemonic to seed err")
		return "get seed err", err
	}

	path, err := accounts.ParseDerivationPath(pathStr)
	if err != nil {
		fmt.Println("path str to path error")
		return "path str to path", err
	}

	// crate master private key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Get masterKey error")
		return "Get masterKey error", err
	}
	//fmt.Println("masterKey:", masterKey)

	// derivation path
	fixIssue172 := true
	for _, n := range path {
		if fixIssue172 && masterKey.IsAffectedByIssue172() {
			masterKey, err = masterKey.Derive(n)
		} else {
			masterKey, err = masterKey.DeriveNonStandard(n)
		}
		if err != nil {
			return "", err
		}
	}

	//fmt.Println("masterKey:", masterKey)
	privateKey, err := masterKey.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return "", err
	}

	// get eth address
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Unable to get public key")
		return "", err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	info := fmt.Sprintf("%v:%v", address, hexutil.Encode(crypto.FromECDSA(privateKeyECDSA)))
	return info, nil
}
