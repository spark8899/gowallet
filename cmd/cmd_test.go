package cmd

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestCLIBuild tests that the binary builds successfully
func TestCLIBuild(t *testing.T) {
	// Build the binary
	cmd := exec.Command("go", "build", "-o", "/tmp/gowallet_test", "../main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build: %v\nOutput: %s", err, output)
	}

	// Clean up
	defer os.Remove("/tmp/gowallet_test")

	// Verify the binary exists
	if _, err := os.Stat("/tmp/gowallet_test"); os.IsNotExist(err) {
		t.Fatal("Binary was not created")
	}
}

// TestVersionCommand tests the version command
func TestVersionCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("version command failed: %v\nOutput: %s", err, output)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Version:") {
		t.Errorf("Expected version output to contain 'Version:', got: %s", outputStr)
	}
}

// TestGenPrivateKeyCommand tests the genPrivateKey command
func TestGenPrivateKeyCommand(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantLines int
	}{
		{
			name:      "Generate 1 key (default)",
			args:      []string{"genPrivateKey"},
			wantLines: 1,
		},
		{
			name:      "Generate 3 keys",
			args:      []string{"genPrivateKey", "3"},
			wantLines: 3,
		},
		{
			name:      "Generate with -n flag",
			args:      []string{"genPrivateKey", "-n", "2"},
			wantLines: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "../main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, output)
			}

			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) != tt.wantLines {
				t.Errorf("Expected %d lines, got %d", tt.wantLines, len(lines))
			}

			// Verify format: address:privatekey
			for i, line := range lines {
				parts := strings.Split(line, ":")
				if len(parts) != 2 {
					t.Errorf("Line %d has invalid format: %s", i, line)
					continue
				}
				// Check address format (0x + 40 hex)
				if !strings.HasPrefix(parts[0], "0x") || len(parts[0]) != 42 {
					t.Errorf("Line %d has invalid address: %s", i, parts[0])
				}
				// Check private key format (0x + 64 hex)
				if !strings.HasPrefix(parts[1], "0x") || len(parts[1]) != 66 {
					t.Errorf("Line %d has invalid private key length: %s", i, parts[1])
				}
			}
		})
	}
}

// TestGenMnemonicCommand tests the genMnemonic command
func TestGenMnemonicCommand(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantWords int
	}{
		{
			name:      "Generate 12-word mnemonic",
			args:      []string{"genMnemonic", "12"},
			wantWords: 12,
		},
		{
			name:      "Generate 24-word mnemonic",
			args:      []string{"genMnemonic", "24"},
			wantWords: 24,
		},
		{
			name:      "Default (12 words)",
			args:      []string{"genMnemonic"},
			wantWords: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "../main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, output)
			}

			words := strings.Fields(strings.TrimSpace(string(output)))
			if len(words) != tt.wantWords {
				t.Errorf("Expected %d words, got %d", tt.wantWords, len(words))
			}
		})
	}
}

// TestHelpCommand tests the help command
func TestHelpCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		contains []string
	}{
		{
			name:     "Root help",
			args:     []string{"--help"},
			contains: []string{"Available Commands", "genPrivateKey", "genMnemonic"},
		},
		{
			name:     "genPrivateKey help",
			args:     []string{"genPrivateKey", "help"},
			contains: []string{"Generate", "private key", "Examples"},
		},
		{
			name:     "genMnemonic help",
			args:     []string{"genMnemonic", "--help"},
			contains: []string{"BIP39", "mnemonic", "12"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "../main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()
			// Help commands exit with 0
			if err != nil {
				// Some help commands may exit with non-zero, that's OK
				t.Logf("Help command exited with: %v", err)
			}

			outputStr := string(output)
			for _, want := range tt.contains {
				if !strings.Contains(outputStr, want) {
					t.Errorf("Expected output to contain '%s', got: %s", want, outputStr)
				}
			}
		})
	}
}

// TestMnToSeedCommand tests mnToSeed with mnemonic
func TestMnToSeedCommand(t *testing.T) {
	mnemonic := "close same tongue random ice cave aim input whale salute squirrel vivid"

	cmd := exec.Command("go", "run", "../main.go", "mnToSeed", "-m", mnemonic)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("mnToSeed command failed: %v\nOutput: %s", err, output)
	}

	outputStr := strings.TrimSpace(string(output))
	// Should be a hex string
	if len(outputStr) != 128 { // 64 bytes = 128 hex chars
		t.Errorf("Expected 128 hex characters, got %d: %s", len(outputStr), outputStr)
	}
}

// TestExamplesInHelp verifies examples work
func TestExamplesInHelp(t *testing.T) {
	// Test that the example from genPrivateKey help actually works
	cmd := exec.Command("go", "run", "../main.go", "genPrivateKey", "5")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Example command failed: %v\nOutput: %s", err, output)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) != 5 {
		t.Errorf("Example should generate 5 keys, got %d", len(lines))
	}
}

// TestInputValidation tests error handling
func TestInputValidation(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "getAddress without key",
			args:        []string{"getAddress"},
			expectError: true,
		},
		{
			name:        "mnToSeed without mnemonic",
			args:        []string{"mnToSeed"},
			expectError: true,
		},
		{
			name:        "getPublicKey without key",
			args:        []string{"getPublicKey"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "../main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got success. Output: %s", output)
				}
			} else {
				if err != nil {
					t.Errorf("Expected success, got error: %v\nOutput: %s", err, output)
				}
			}
		})
	}
}
