package commonPrivateKey

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
)

var (
	// ErrPrivateKeyNil indicates the private key is nil
	ErrPrivateKeyNil = errors.New("private key is nil")
	// ErrPrivateKeyZero indicates the private key is zero
	ErrPrivateKeyZero = errors.New("private key is zero")
	// ErrPrivateKeyTooSmall indicates a potentially weak private key
	ErrPrivateKeyTooSmall = errors.New("private key is too small (potentially weak)")
	// ErrPrivateKeyTooLarge indicates the private key exceeds the curve order
	ErrPrivateKeyTooLarge = errors.New("private key exceeds curve order")
)

// secp256k1 curve order (n)
// This is the order of the base point G on the secp256k1 curve
var curveOrder = new(big.Int).SetBytes([]byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE,
	0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B,
	0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x41,
})

// Minimum acceptable private key value
// We avoid very small values as they could be considered weak
// In practice, the probability of generating such a key randomly is negligible
var minPrivateKey = big.NewInt(1000)

// ValidatePrivateKey checks if an ECDSA private key meets security requirements
//
// A valid private key must:
//  1. Not be nil
//  2. Have a value D in the range [1, n-1] where n is the curve order
//  3. Not be too small (avoid potential weak keys)
//
// Note: The probability of randomly generating a weak key is astronomically small
// (less than 2^-128), but we check anyway for defense in depth.
func ValidatePrivateKey(privateKey *ecdsa.PrivateKey) error {
	if privateKey == nil {
		return ErrPrivateKeyNil
	}

	if privateKey.D == nil {
		return ErrPrivateKeyNil
	}

	// Check if key is zero
	if privateKey.D.Sign() == 0 {
		return ErrPrivateKeyZero
	}

	// Check if key is too small (potential weak key)
	// Note: This is extremely paranoid - crypto/rand should never generate such a key
	if privateKey.D.Cmp(minPrivateKey) < 0 {
		return ErrPrivateKeyTooSmall
	}

	// Check if key exceeds curve order
	// This should also never happen with properly generated keys
	if privateKey.D.Cmp(curveOrder) >= 0 {
		return ErrPrivateKeyTooLarge
	}

	return nil
}
