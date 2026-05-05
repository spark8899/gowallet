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
	"github.com/spark8899/gowallet/internal/security"
)

func generateKey() {
	// Generate ECDSA private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	defer security.ZeroBigInt(privateKey.D)

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
	keyBts := crypto.FromECDSA(privateKey)
	defer security.ZeroBytes(keyBts)
	privateKeyHex := hexutil.Encode(keyBts)

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

func PrivateKey(privateKeyHex []byte) (*ecdsa.PrivateKey, error) {
	privateKeyStr := string(privateKeyHex)
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKeyStr, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func PrivateKeyBytes(privateKeyHex []byte) ([]byte, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}
	defer security.ZeroBigInt(privateKey.D)

	return crypto.FromECDSA(privateKey), nil
}

func PrivateKeyHex(privateKeyHex []byte) (string, error) {
	privateKeyBytes, err := PrivateKeyBytes(privateKeyHex)
	if err != nil {
		return "", err
	}
	defer security.ZeroBytes(privateKeyBytes)

	return hexutil.Encode(privateKeyBytes), nil
}

func PublicKey(privateKeyHex []byte) (*ecdsa.PublicKey, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}
	defer security.ZeroBigInt(privateKey.D)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}

	return publicKeyECDSA, nil
}

func PublicKeyBytes(privateKeyHex []byte) ([]byte, error) {
	publicKey, err := PublicKey(privateKeyHex)
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSAPub(publicKey), nil
}

func PublicKeyHex(privateKeyHex []byte) (string, error) {
	publicKeyBytes, err := PublicKeyBytes(privateKeyHex)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(publicKeyBytes), nil
}

func Address(privateKeyHex []byte) (common.Address, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return common.Address{}, err
	}
	defer security.ZeroBigInt(privateKey.D)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("failed to cast public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func AddressBytes(privateKeyHex []byte) ([]byte, error) {
	address, err := Address(privateKeyHex)
	if err != nil {
		return nil, err
	}
	return address.Bytes(), nil
}

func AddressHex(privateKeyHex []byte) (string, error) {
	address, err := Address(privateKeyHex)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

func SignHash(privateKeyHex []byte, hash []byte) ([]byte, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}
	defer security.ZeroBigInt(privateKey.D)

	return crypto.Sign(hash, privateKey)
}

func SignTxEIP155(privateKeyHex []byte, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}
	defer security.ZeroBigInt(privateKey.D)

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

	address, err := Address(privateKeyHex)
	if err != nil {
		return nil, err
	}

	if sender != address {
		return nil, fmt.Errorf("signer mismatch: expected %s, got %s", address.Hex(), sender.Hex())
	}

	return signedTx, nil
}

func SignTx(privateKeyHex []byte, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	privateKey, err := PrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}
	defer security.ZeroBigInt(privateKey.D)

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

	address, err := Address(privateKeyHex)
	if err != nil {
		return nil, err
	}

	if sender != address {
		return nil, fmt.Errorf("signer mismatch: expected %s, got %s", address.Hex(), sender.Hex())
	}

	return signedTx, nil
}

func SignData(privateKeyHex []byte, mimeType string, data []byte) ([]byte, error) {
	return SignHash(privateKeyHex, crypto.Keccak256(data))
}

func SignText(privateKeyHex []byte, text []byte) ([]byte, error) {
	return SignHash(privateKeyHex, accounts.TextHash(text))
}
