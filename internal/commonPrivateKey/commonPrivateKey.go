package commonPrivateKey

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateKey() {
	// from privateKey create publicKey
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("failed to cast public key to ECDSA")
	}

	// get privateKey
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	//fmt.Printf("Private_key: %v\n", privateKeyHex)

	// get PublicKey and address
	//publicKeyBytes := crypto.FromECDSAPub(ecdsaPublicKey)
	address := crypto.PubkeyToAddress(*ecdsaPublicKey)
	//fmt.Printf("Address: %v\nPublic key: %v\n", address.Hex(), hexutil.Encode(publicKeyBytes))

	fmt.Printf("%v:%v\n", address.Hex(), privateKeyHex)
}

func GetGenerateKey(num int) {
	for i := 0; i < num; i++ {
		generateKey()
	}
}

func GetEthAddress(privateKey string) {
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKey, "0x")
	privKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		fmt.Println("invalid private key")
		return
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Unable to get public key")
		return
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
}

func GetPublicKey(privateKey string) {
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKey, "0x")
	privKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		fmt.Println("invalid private key")
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes))
}
