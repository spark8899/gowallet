package hdwallet

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

var wordList = strings.Split(alternatingWords, "\n")

var wordIndexes = make(map[string]uint16, len(wordList))

func init() {
	for i, word := range wordList {
		wordIndexes[strings.ToLower(word)] = uint16(i)
	}
}

func Bip39GenMnemonic(size int) (string, error) {
	entropyBytes, err := bip39.NewEntropy(size)
	if err != nil {
		return "", err
	}
	//生成助记词
	mnemonic, err := bip39.NewMnemonic(entropyBytes)
	return mnemonic, err
}

func Bip39MnemonicToSeed(mnemonic string, password string) ([]byte, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("mnemonic not valid")
	}
	return bip39.NewSeed(mnemonic, password), nil
}

// dcr 的助记词跟种子的转换可以互逆，没有遵守 bip39 规范
func DcrSeedToMnemonic(seed []byte) string {
	var buf bytes.Buffer
	for i, b := range seed {
		if i != 0 {
			buf.WriteRune(' ')
		}
		buf.WriteString(byteToMnemonic(b, i))
	}
	checksum := checksumByte(seed)
	buf.WriteRune(' ')
	buf.WriteString(byteToMnemonic(checksum, len(seed)))
	return buf.String()
}

// DecodeMnemonics returns the decoded value that is encoded by words.  Any
// words that are whitespace are empty are skipped.
func DcrMnemonicToSeed(words []string) ([]byte, error) {
	decoded := make([]byte, len(words))
	idx := 0
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w == "" {
			continue
		}
		b, ok := wordIndexes[strings.ToLower(w)]
		if !ok {
			return nil, errors.New("word %v is not in the PGP word list")
		}
		if int(b%2) != idx%2 {
			return nil, errors.New("word %v is not valid at position %v, ")
		}
		decoded[idx] = byte(b / 2)
		idx++
	}
	return decoded[:idx], nil
}

func checksumByte(data []byte) byte {
	intermediateHash := sha256.Sum256(data)
	return sha256.Sum256(intermediateHash[:])[0]
}

// byteToMnemonic returns the PGP word list encoding of b when found at index.
func byteToMnemonic(b byte, index int) string {
	bb := uint16(b) * 2
	if index%2 != 0 {
		bb++
	}
	return wordList[bb]
}
