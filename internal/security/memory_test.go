package security

import (
	"bytes"
	"math/big"
	"testing"
)

func TestZeroBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{
			name:  "small slice",
			input: []byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			name:  "larger slice",
			input: bytes.Repeat([]byte{0xff}, 32),
		},
		{
			name:  "empty slice",
			input: []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to preserve original for verification
			original := make([]byte, len(tt.input))
			copy(original, tt.input)

			ZeroBytes(tt.input)

			// Verify all bytes are zero
			for i, b := range tt.input {
				if b != 0 {
					t.Errorf("ZeroBytes() byte at index %d is %d, want 0", i, b)
				}
			}
		})
	}
}

func TestZeroBigInt(t *testing.T) {
	tests := []struct {
		name  string
		input *big.Int
	}{
		{
			name:  "small positive number",
			input: big.NewInt(12345),
		},
		{
			name:  "large number",
			input: new(big.Int).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
		},
		{
			name:  "negative number",
			input: big.NewInt(-12345),
		},
		{
			name:  "nil big int",
			input: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ZeroBigInt(tt.input)

			if tt.input != nil {
				if tt.input.Cmp(big.NewInt(0)) != 0 {
					t.Errorf("ZeroBigInt() result is %v, want 0", tt.input)
				}
			}
		})
	}
}

func BenchmarkZeroBytes(b *testing.B) {
	data := make([]byte, 32) // 256 bits like a private key
	for i := 0; i < b.N; i++ {
		// Reset data
		for j := range data {
			data[j] = 0xff
		}
		ZeroBytes(data)
	}
}
