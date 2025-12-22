package hdwallet

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrInvalidPathFormat indicates the path doesn't match BIP32 format
	ErrInvalidPathFormat = errors.New("invalid derivation path format")
	// ErrPathTooDeep indicates the derivation path has too many levels
	ErrPathTooDeep = errors.New("derivation path is too deep")
	// ErrInvalidPurpose indicates an unsupported purpose value
	ErrInvalidPurpose = errors.New("invalid purpose in derivation path")
	// ErrEntropyTooShort indicates insufficient entropy bits
	ErrEntropyTooShort = errors.New("entropy is too short")
	// ErrEntropyAllZeros indicates entropy consists of all zeros
	ErrEntropyAllZeros = errors.New("entropy is all zeros (weak randomness)")
	// ErrEntropyPoorQuality indicates entropy fails randomness checks
	ErrEntropyPoorQuality = errors.New("entropy has poor randomness quality")
)

const (
	maxPathDepth   = 10  // Maximum recommended depth for derivation paths
	minEntropyBits = 128 // Minimum entropy bits (BIP39 minimum)
)

// Regex to validate BIP32 path format: m/44'/60'/0'/0/0
var pathRegex = regexp.MustCompile(`^m(/\d+'?)+$`)

// ValidateDerivationPath validates a BIP32/BIP44 derivation path
//
// A valid path must:
//  1. Start with 'm' and have at least one level
//  2. Follow the format m/level1/level2/.../levelN
//  3. Each level is a non-negative integer, optionally followed by ' for hardened
//  4. Not exceed maximum depth (default 10)
//  5. Have a standard purpose (44, 49, or 84) if specified
func ValidateDerivationPath(pathStr string) error {
	if pathStr == "" {
		return ErrInvalidPathFormat
	}

	// Check basic format
	if !pathRegex.MatchString(pathStr) {
		return ErrInvalidPathFormat
	}

	// Parse and count depth
	parts := strings.Split(pathStr, "/")
	if len(parts) > maxPathDepth {
		return fmt.Errorf("%w: depth %d exceeds maximum %d", ErrPathTooDeep, len(parts), maxPathDepth)
	}

	// Validate purpose (first level after m) if present
	if len(parts) >= 2 {
		purposeStr := strings.TrimSuffix(parts[1], "'")
		purpose, err := strconv.Atoi(purposeStr)
		if err != nil {
			return fmt.Errorf("invalid purpose: %w", err)
		}

		// Common BIP purposes:
		// 44 - BIP44 (original HD wallet standard)
		// 49 - BIP49 (SegWit in P2SH)
		// 84 - BIP84 (Native SegWit)
		validPurposes := []int{44, 49, 84}
		valid := false
		for _, vp := range validPurposes {
			if purpose == vp {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("%w: %d (expected 44, 49, or 84)", ErrInvalidPurpose, purpose)
		}
	}

	return nil
}

// ValidateEntropy checks the quality of entropy bytes
//
// Good entropy should:
//  1. Be at least 128 bits (16 bytes) for BIP39
//  2. Not be all zeros or any obvious pattern
//  3. Have roughly 50% ones in binary representation (Hamming weight)
func ValidateEntropy(entropy []byte) error {
	// Check minimum length
	entropyBits := len(entropy) * 8
	if entropyBits < minEntropyBits {
		return fmt.Errorf("%w: got %d bits, need at least %d", ErrEntropyTooShort, entropyBits, minEntropyBits)
	}

	// Check for all zeros (extremely weak entropy)
	allZero := true
	for _, b := range entropy {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return ErrEntropyAllZeros
	}

	// Check Hamming weight (count of 1 bits)
	// Good random data should have approximately 50% ones
	ones := 0
	for _, b := range entropy {
		ones += hammingWeight(b)
	}

	totalBits := len(entropy) * 8
	onesRatio := float64(ones) / float64(totalBits)

	// Allow 40%-60% ones (looser than ideal 50% for practical use)
	// This catches obvious non-random patterns while allowing normal variance
	if onesRatio < 0.40 || onesRatio > 0.60 {
		return fmt.Errorf("%w: hamming weight %.2f%% (expected 40-60%%)",
			ErrEntropyPoorQuality, onesRatio*100)
	}

	return nil
}

// hammingWeight counts the number of 1 bits in a byte
func hammingWeight(b byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		if b&(1<<i) != 0 {
			count++
		}
	}
	return count
}
