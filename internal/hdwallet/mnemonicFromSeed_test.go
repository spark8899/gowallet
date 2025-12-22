package hdwallet

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestMnemonicFromSeed(t *testing.T) {
	tests := []struct {
		name        string
		seedHex     string
		expectError bool
		errorMsg    string
	}{
		{
			name:    "Valid 128-bit seed",
			seedHex: "00112233445566778899aabbccddeeff",
		},
		{
			name:    "Valid 256-bit seed",
			seedHex: "00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff",
		},
		{
			name:        "Invalid hex characters",
			seedHex:     "gghhiijj",
			expectError: true,
			errorMsg:    "failed to decode seed string",
		},
		{
			name:        "Odd length hex string",
			seedHex:     "00112233445",
			expectError: true,
			errorMsg:    "failed to decode seed string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MnemonicFromSeed(tt.seedHex)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorMsg)
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Verify result is a non-empty string
			if result == "" {
				t.Error("Expected non-empty mnemonic result")
			}

			// Count words in the mnemonic
			words := strings.Fields(result)
			if len(words) < 12 {
				t.Errorf("Expected at least 12 words, got %d", len(words))
			}
		})
	}
}

func TestMnemonicFromSeed_Deterministic(t *testing.T) {
	// Test that the same seed always produces the same mnemonic
	seedHex := "00112233445566778899aabbccddeeff"

	result1, err1 := MnemonicFromSeed(seedHex)
	if err1 != nil {
		t.Fatalf("First call failed: %v", err1)
	}

	result2, err2 := MnemonicFromSeed(seedHex)
	if err2 != nil {
		t.Fatalf("Second call failed: %v", err2)
	}

	if result1 != result2 {
		t.Errorf("Results are not deterministic: '%s' != '%s'", result1, result2)
	}
}

func TestMnemonicFromSeed_DifferentSeeds(t *testing.T) {
	// Test that different seeds produce different mnemonics
	seed1 := "00112233445566778899aabbccddeeff"
	seed2 := "ffeeddccbbaa99887766554433221100"

	result1, err1 := MnemonicFromSeed(seed1)
	if err1 != nil {
		t.Fatalf("Seed1 failed: %v", err1)
	}

	result2, err2 := MnemonicFromSeed(seed2)
	if err2 != nil {
		t.Fatalf("Seed2 failed: %v", err2)
	}

	if result1 == result2 {
		t.Error("Different seeds should produce different mnemonics")
	}
}

func TestMnemonicFromSeed_ValidBIP39(t *testing.T) {
	// Generate a mnemonic from a seed and verify it's a valid BIP39 mnemonic
	seedHex := "0123456789abcdef0123456789abcdef"

	mnemonic, err := MnemonicFromSeed(seedHex)
	if err != nil {
		t.Fatalf("Failed to generate mnemonic: %v", err)
	}

	// The generated mnemonic should be convertible back to a seed
	// (though not necessarily the original seed due to the way the function works)
	seed, err := Bip39MnemonicToSeed(mnemonic, "")
	if err != nil {
		t.Errorf("Generated mnemonic is not valid BIP39: %v", err)
	}

	if len(seed) != 64 {
		t.Errorf("Expected 64-byte seed, got %d bytes", len(seed))
	}
}

func TestMnemonicFromSeed_HexDecoding(t *testing.T) {
	// Test with uppercase and lowercase hex
	seedLower := "aabbccdd"
	seedUpper := "AABBCCDD"

	resultLower, errLower := MnemonicFromSeed(seedLower)
	if errLower != nil {
		t.Fatalf("Lowercase hex failed: %v", errLower)
	}

	resultUpper, errUpper := MnemonicFromSeed(seedUpper)
	if errUpper != nil {
		t.Fatalf("Uppercase hex failed: %v", errUpper)
	}

	if resultLower != resultUpper {
		t.Error("Hex case should not affect result")
	}
}

// Benchmark for performance testing
func BenchmarkMnemonicFromSeed(b *testing.B) {
	seedHex := "00112233445566778899aabbccddeeff"
	for i := 0; i < b.N; i++ {
		_, _ = MnemonicFromSeed(seedHex)
	}
}

// Test helper function to verify hex encoding/decoding
func TestHexEncodeDecode(t *testing.T) {
	original := "0123456789abcdef"
	decoded, err := hex.DecodeString(original)
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	encoded := hex.EncodeToString(decoded)
	if original != encoded {
		t.Errorf("Hex encode/decode roundtrip failed: %s != %s", original, encoded)
	}
}
