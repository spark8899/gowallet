package hdwallet

import (
	"fmt"
	"strings"
	"testing"
)

func TestPathFromSeed(t *testing.T) {
	// Use a valid 64-byte hex seed (128 hex characters)
	validSeed := "126b7f8653ce2b1f05dd78d33c57737df4edf889ee2729338202d164831e2ab43d40d2a26d73739570cf816cb96d766b8d3850258d58c89f7e9901edf13e80a8"

	tests := []struct {
		name        string
		seedHex     string
		path        string
		expectError bool
		errorMsg    string
	}{
		{
			name:    "Valid seed with standard ETH path",
			seedHex: validSeed,
			path:    "m/44'/60'/0'/0/0",
		},
		{
			name:        "Empty seed",
			seedHex:     "",
			path:        "m/44'/60'/0'/0/0",
			expectError: true,
			errorMsg:    "masterKey",
		},
		{
			name:        "Invalid hex seed",
			seedHex:     "gghhiijj",
			path:        "m/44'/60'/0'/0/0",
			expectError: true,
			errorMsg:    "seed str to bytes",
		},
		{
			name:        "Empty path",
			seedHex:     validSeed,
			path:        "",
			expectError: true,
			errorMsg:    "path str to path",
		},
		{
			name:        "Invalid path format",
			seedHex:     validSeed,
			path:        "invalid/path",
			expectError: true,
			errorMsg:    "path str to path",
		},
		{
			name:    "Different account index",
			seedHex: validSeed,
			path:    "m/44'/60'/1'/0/0",
		},
		{
			name:    "Different address index",
			seedHex: validSeed,
			path:    "m/44'/60'/0'/0/5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PathFromSeed(tt.seedHex, tt.path)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorMsg)
				} else if !strings.Contains(result, tt.errorMsg) && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Logf("Expected error containing '%s', got result='%s', err='%v'", tt.errorMsg, result, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check result format: should be address:privateKey
			parts := strings.Split(result, ":")
			if len(parts) != 2 {
				t.Errorf("Expected result format 'address:privateKey', got: %s", result)
				return
			}

			// Verify address format (0x + 40 hex chars)
			address := parts[0]
			if !strings.HasPrefix(address, "0x") || len(address) != 42 {
				t.Errorf("Invalid address format: %s", address)
			}

			// Verify private key format (0x + 64 hex chars)
			privateKey := parts[1]
			if !strings.HasPrefix(privateKey, "0x") || len(privateKey) != 66 {
				t.Errorf("Invalid private key format: %s", privateKey)
			}
		})
	}
}

func TestPathFromSeed_Deterministic(t *testing.T) {
	// Test that the same seed and path always produce the same result
	seedHex := "126b7f8653ce2b1f05dd78d33c57737df4edf889ee2729338202d164831e2ab43d40d2a26d73739570cf816cb96d766b8d3850258d58c89f7e9901edf13e80a8"
	path := "m/44'/60'/0'/0/0"

	result1, err1 := PathFromSeed(seedHex, path)
	if err1 != nil {
		t.Fatalf("First call failed: %v", err1)
	}

	result2, err2 := PathFromSeed(seedHex, path)
	if err2 != nil {
		t.Fatalf("Second call failed: %v", err2)
	}

	if result1 != result2 {
		t.Errorf("Results are not deterministic: '%s' != '%s'", result1, result2)
	}
}

func TestPathFromSeed_DifferentPaths(t *testing.T) {
	// Test that different paths produce different results
	seedHex := "126b7f8653ce2b1f05dd78d33c57737df4edf889ee2729338202d164831e2ab43d40d2a26d73739570cf816cb96d766b8d3850258d58c89f7e9901edf13e80a8"
	path1 := "m/44'/60'/0'/0/0"
	path2 := "m/44'/60'/0'/0/1"

	result1, err1 := PathFromSeed(seedHex, path1)
	if err1 != nil {
		t.Fatalf("Path1 failed: %v", err1)
	}

	result2, err2 := PathFromSeed(seedHex, path2)
	if err2 != nil {
		t.Fatalf("Path2 failed: %v", err2)
	}

	if result1 == result2 {
		t.Errorf("Different paths should produce different results")
	}

	// Extract addresses
	addr1 := strings.Split(result1, ":")[0]
	addr2 := strings.Split(result2, ":")[0]

	if addr1 == addr2 {
		t.Errorf("Different paths should produce different addresses")
	}
}

func TestPathFromSeed_MnemonicConsistency(t *testing.T) {
	// Test that mnemonic->seed->path and direct path from mnemonic give same result
	mnemonic := "close same tongue random ice cave aim input whale salute squirrel vivid"
	path := "m/44'/60'/0'/0/0"

	// Get result from mnemonic directly
	resultFromMnemonic, err1 := PathFromMnemonic(mnemonic, path)
	if err1 != nil {
		t.Fatalf("PathFromMnemonic failed: %v", err1)
	}

	// Convert mnemonic to seed then derive path
	seed, err2 := Bip39MnemonicToSeed(mnemonic, "")
	if err2 != nil {
		t.Fatalf("Bip39MnemonicToSeed failed: %v", err2)
	}

	seedHex := fmt.Sprintf("%x", seed)
	resultFromSeed, err3 := PathFromSeed(seedHex, path)
	if err3 != nil {
		t.Fatalf("PathFromSeed failed: %v", err3)
	}

	// Both methods should produce the same result
	if resultFromMnemonic != resultFromSeed {
		t.Errorf("Mismatch between PathFromMnemonic and PathFromSeed:\nMnemonic: %s\nSeed: %s", resultFromMnemonic, resultFromSeed)
	}
}

// Benchmark for performance testing
func BenchmarkPathFromSeed(b *testing.B) {
	seedHex := "126b7f8653ce2b1f05dd78d33c57737df4edf889ee2729338202d164831e2ab43d40d2a26d73739570cf816cb96d766b8d3850258d58c89f7e9901edf13e80a8"
	path := "m/44'/60'/0'/0/0"

	for i := 0; i < b.N; i++ {
		_, _ = PathFromSeed(seedHex, path)
	}
}
