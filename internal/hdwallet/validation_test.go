package hdwallet

import (
	"strings"
	"testing"
)

func TestValidateDerivationPath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
		errType error
	}{
		{
			name:    "empty path",
			path:    "",
			wantErr: true,
			errType: ErrInvalidPathFormat,
		},
		{
			name:    "valid BIP44 Ethereum path",
			path:    "m/44'/60'/0'/0/0",
			wantErr: false,
		},
		{
			name:    "valid BIP49 path",
			path:    "m/49'/0'/0'/0/0",
			wantErr: false,
		},
		{
			name:    "valid BIP84 path",
			path:    "m/84'/0'/0'/0/0",
			wantErr: false,
		},
		{
			name:    "path without m prefix",
			path:    "44'/60'/0'/0/0",
			wantErr: true,
			errType: ErrInvalidPathFormat,
		},
		{
			name:    "path with invalid characters",
			path:    "m/44'/60'/a/0/0",
			wantErr: true,
			errType: ErrInvalidPathFormat,
		},
		{
			name:    "path too deep",
			path:    "m/44'/60'/0'/0/0/1/2/3/4/5/6",
			wantErr: true,
			errType: ErrPathTooDeep,
		},
		{
			name:    "valid short path",
			path:    "m/44'",
			wantErr: false,
		},
		{
			name:    "path with invalid purpose",
			path:    "m/99'/60'/0'/0/0",
			wantErr: true,
			errType: ErrInvalidPurpose,
		},
		{
			name:    "path with mixed hardened notation",
			path:    "m/44'/60/0'/0/0",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDerivationPath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDerivationPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errType != nil {
				if !strings.Contains(err.Error(), tt.errType.Error()) {
					t.Errorf("ValidateDerivationPath() error = %v, want error containing %v", err, tt.errType)
				}
			}
		})
	}
}

func TestValidateEntropy(t *testing.T) {
	tests := []struct {
		name    string
		entropy []byte
		wantErr bool
		errType error
	}{
		{
			name:    "too short entropy",
			entropy: []byte{0x01, 0x02, 0x03},
			wantErr: true,
			errType: ErrEntropyTooShort,
		},
		{
			name:    "all zeros (weak)",
			entropy: make([]byte, 16), // 128 bits of zeros
			wantErr: true,
			errType: ErrEntropyAllZeros,
		},
		{
			name:    "all ones (poor quality)",
			entropy: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			wantErr: true,
			errType: ErrEntropyPoorQuality,
		},
		{
			name: "good quality entropy (50% ones)",
			entropy: []byte{
				0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, // 10101010 repeated
				0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, // 50% ones
			},
			wantErr: false,
		},
		{
			name: "acceptable entropy (45% ones)",
			entropy: []byte{
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, // 01010101 repeated
				0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, // Mix to get ~45%
			},
			wantErr: false,
		},
		{
			name: "256-bit good entropy",
			entropy: []byte{
				0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
				0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEntropy(tt.entropy)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEntropy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errType != nil {
				if !strings.Contains(err.Error(), tt.errType.Error()) {
					t.Errorf("ValidateEntropy() error = %v, want error containing %v", err, tt.errType)
				}
			}
		})
	}
}

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want int
	}{
		{"all zeros", 0x00, 0},
		{"all ones", 0xff, 8},
		{"alternating 1", 0xaa, 4}, // 10101010
		{"alternating 2", 0x55, 4}, // 01010101
		{"single bit", 0x01, 1},
		{"single bit", 0x80, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.b)
			if got != tt.want {
				t.Errorf("hammingWeight(%08b) = %d, want %d", tt.b, got, tt.want)
			}
		})
	}
}
