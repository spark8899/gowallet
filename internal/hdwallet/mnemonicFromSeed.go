package hdwallet

import (
	"crypto/sha256"
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

	//将种子转换为主私钥
	masterKey01, _ := bip32.NewMasterKey(seed)

	//对seed进行SHA-256哈希
	hashedSeed := sha256.Sum256(seed)

	//将熵转换为助记词
	mnemonic, _ := bip39.NewMnemonic(masterKey01.Key)

	fmt.Println("hashedSeed: ", hashedSeed)
	fmt.Println("Mnemonic: ", mnemonic)
	// fmt.Println("Mnemonic:", mnemonic)
	return mnemonic, nil
}
