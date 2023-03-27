package commonPrivateKey

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestPrivateKey(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	if privateKeyStr != privateKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestPrivateKeyBytes(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	privateKeyBytes, err := PrivateKeyBytes(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	privateKeyHex := hexutil.Encode(privateKeyBytes)
	if privateKeyStr != privateKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestPrivateKeyHex(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	privateKeyHex, err := PrivateKeyHex(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	if privateKeyStr != privateKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestPublicKey(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	publicKeyStr := "0x046005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18"
	publicKey, err := PublicKey(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	publicKeyHex := hexutil.Encode(publicKeyBytes)
	if publicKeyStr != publicKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestPublicKeyBytes(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	publicKeyStr := "0x046005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18"
	publicKeyBytes, err := PublicKeyBytes(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	publicKeyHex := hexutil.Encode(publicKeyBytes)
	if publicKeyStr != publicKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestPublicKeyHex(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	publicKeyStr := "0x046005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18"
	publicKeyHex, err := PublicKeyHex(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	if publicKeyStr != publicKeyHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestAddress(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	addressStr := "0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"
	address, err := Address(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	addressHex := address.Hex()
	if addressStr != addressHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestAddressBytes(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	addressStr := "0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"
	addressBytes, err := AddressBytes(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	addressHex := hexutil.Encode(addressBytes)
	if strings.ToLower(addressStr) != addressHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestAddressHex(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	addressStr := "0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"
	addressHex, err := AddressHex(privateKeyStr)
	if err != nil {
		t.Error(err)
	}

	if addressStr != addressHex {
		t.Errorf("expected match")
	} else {
		t.Log("ok")
	}
}

func TestSignHash(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	sig, err := SignHash(privateKeyStr, hash.Bytes())
	if err != nil {
		t.Error(err)
	}
	if len(sig) == 0 {
		t.Error("expected signature")
	} else {
		t.Log("ok")
	}
}

func TestSignTxEIP155(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	data := []byte{}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := SignTxEIP155(privateKeyStr, tx, big.NewInt(42))
	if err != nil {
		t.Error(err)
	}

	v, r, s := signedTx.RawSignatureValues()
	if v.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected v value")
	}
	if r.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected r value")
	}
	if s.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected s value")
	} else {
		t.Log("ok")
	}
}

func TestSignTx(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	data := []byte{}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := SignTx(privateKeyStr, tx, big.NewInt(42))
	if err != nil {
		t.Error(err)
	}

	v, r, s := signedTx.RawSignatureValues()
	if v.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected v value")
	}
	if r.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected r value")
	}
	if s.Cmp(big.NewInt(0)) != 1 {
		t.Error("expected s value")
	} else {
		t.Log("ok")
	}
}

func TestSignData(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	mimeType := "text/plain"
	data := []byte("hello world")

	signedData, err := SignData(privateKeyStr, mimeType, data)
	if err != nil {
		t.Error(err)
	}

	if len(signedData) == 0 {
		t.Error("Expected signature")
	} else {
		t.Log("ok")
	}
}

func TestSignText(t *testing.T) {
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	data := []byte("hello world")

	signedTextData, err := SignText(privateKeyStr, data)
	if err != nil {
		t.Error(err)
	}

	if len(signedTextData) == 0 {
		t.Error("Expected signature")
	} else {
		t.Log("ok")
	}
}
