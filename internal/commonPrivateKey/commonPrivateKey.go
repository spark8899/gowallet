package commonPrivateKey

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateKey() {
	// Generate ECDSA private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Validate private key strength
	// Note: This is a defense-in-depth measure. crypto.GenerateKey() should
	// always produce valid keys, but we verify anyway for security.
	if err := ValidatePrivateKey(privateKey); err != nil {
		log.Fatalf("Generated key failed validation: %v", err)
	}

	publicKey := privateKey.Public()
	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("failed to cast public key to ECDSA")
	}

	// Convert private key to hex
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))

	// Derive address from public key
	address := crypto.PubkeyToAddress(*ecdsaPublicKey)

	fmt.Printf("%v:%v\n", address.Hex(), privateKeyHex)
}

func GetGenerateKey(num int) {
	if num <= 0 {
		log.Println("Warning: number of keys must be positive, defaulting to 1")
		num = 1
	}
	if num > 1000 {
		log.Println("Warning: limiting key generation to 1000 keys")
		num = 1000
	}
	for i := 0; i < num; i++ {
		generateKey()
	}
}

func PrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKeyStr, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func PrivateKeyBytes(privateKeyStr string) ([]byte, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSA(privateKey), nil
}

func PrivateKeyHex(privateKeyStr string) (string, error) {
	privateKeyBytes, err := PrivateKeyBytes(privateKeyStr)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(privateKeyBytes), nil
}

func PublicKey(privateKeyStr string) (*ecdsa.PublicKey, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}

	return publicKeyECDSA, nil
}

func PublicKeyBytes(privateKeyStr string) ([]byte, error) {
	publicKey, err := PublicKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSAPub(publicKey), nil
}

func PublicKeyHex(privateKeyStr string) (string, error) {
	publicKeyBytes, err := PublicKeyBytes(privateKeyStr)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(publicKeyBytes), nil
}

func Address(privateKeyStr string) (common.Address, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("failed to cast public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func AddressBytes(privateKeyStr string) ([]byte, error) {
	address, err := Address(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return address.Bytes(), nil
}

func AddressHex(privateKeyStr string) (string, error) {
	address, err := Address(privateKeyStr)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

func SignHash(privateKeyStr string, hash []byte) ([]byte, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(hash, privateKey)
}

func SignTxEIP155(privateKeyStr string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID)
	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return nil, err
	}

	address, err := Address(privateKeyStr)
	if err != nil {
		return nil, err
	}

	if sender != address {
		return nil, fmt.Errorf("signer mismatch: expected %s, got %s", address.Hex(), sender.Hex())
	}

	return signedTx, nil
}

func SignTx(privateKeyStr string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signer := types.LatestSignerForChainID(chainID)

	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return nil, err
	}

	address, err := Address(privateKeyStr)
	if err != nil {
		return nil, err
	}

	if sender != address {
		return nil, fmt.Errorf("signer mismatch: expected %s, got %s", address.Hex(), sender.Hex())
	}

	return signedTx, nil
}

func SignData(privateKeyStr string, mimeType string, data []byte) ([]byte, error) {
	return SignHash(privateKeyStr, crypto.Keccak256(data))
}

func SignText(privateKeyStr string, text []byte) ([]byte, error) {
	return SignHash(privateKeyStr, accounts.TextHash(text))
}
