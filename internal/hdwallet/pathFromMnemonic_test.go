package hdwallet

import (
	"strings"
	"testing"
)

func TestPathFromMnemonic(t *testing.T) {
	tests := []struct {
		name        string
		mnemonic    string
		path        string
		expectError bool
		errorMsg    string
	}{
		{
			name:     "Valid 12-word mnemonic with standard ETH path",
			mnemonic: "close same tongue random ice cave aim input whale salute squirrel vivid",
			path:     "m/44'/60'/0'/0/0",
		},
		{
			name:     "Valid 24-word mnemonic with standard ETH path",
			mnemonic: "armed fantasy witness similar prosper poet throw video cannon original video zone talk swear economy bachelor urban crunch mouse trial joy little smart marble",
			path:     "m/44'/60'/0'/0/0",
		},
		{
			name:        "Empty mnemonic",
			mnemonic:    "",
			path:        "m/44'/60'/0'/0/0",
			expectError: true,
			errorMsg:    "mnemonic is required",
		},
		{
			name:        "Invalid mnemonic",
			mnemonic:    "invalid word word word word word word word word word word word word",
			path:        "m/44'/60'/0'/0/0",
			expectError: true,
			errorMsg:    "mnemonic is invalid",
		},
		{
			name:        "Empty path",
			mnemonic:    "close same tongue random ice cave aim input whale salute squirrel vivid",
			path:        "",
			expectError: true,
			errorMsg:    "invalid derivation path",
		},
		{
			name:        "Invalid path format",
			mnemonic:    "close same tongue random ice cave aim input whale salute squirrel vivid",
			path:        "invalid/path",
			expectError: true,
			errorMsg:    "invalid derivation path",
		},
		{
			name:     "Different account index",
			mnemonic: "close same tongue random ice cave aim input whale salute squirrel vivid",
			path:     "m/44'/60'/1'/0/0",
		},
		{
			name:     "Different address index",
			mnemonic: "close same tongue random ice cave aim input whale salute squirrel vivid",
			path:     "m/44'/60'/0'/0/5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PathFromMnemonic(tt.mnemonic, tt.path)

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

func TestPathFromMnemonic_Deterministic(t *testing.T) {
	// Test that the same mnemonic and path always produce the same result
	mnemonic := "close same tongue random ice cave aim input whale salute squirrel vivid"
	path := "m/44'/60'/0'/0/0"

	result1, err1 := PathFromMnemonic(mnemonic, path)
	if err1 != nil {
		t.Fatalf("First call failed: %v", err1)
	}

	result2, err2 := PathFromMnemonic(mnemonic, path)
	if err2 != nil {
		t.Fatalf("Second call failed: %v", err2)
	}

	if result1 != result2 {
		t.Errorf("Results are not deterministic: '%s' != '%s'", result1, result2)
	}
}

func TestPathFromMnemonic_DifferentPaths(t *testing.T) {
	// Test that different paths produce different results
	mnemonic := "close same tongue random ice cave aim input whale salute squirrel vivid"
	path1 := "m/44'/60'/0'/0/0"
	path2 := "m/44'/60'/0'/0/1"

	result1, err1 := PathFromMnemonic(mnemonic, path1)
	if err1 != nil {
		t.Fatalf("Path1 failed: %v", err1)
	}

	result2, err2 := PathFromMnemonic(mnemonic, path2)
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
