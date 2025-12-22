package commonPrivateKey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"testing"
)

func TestValidatePrivateKey(t *testing.T) {
	tests := []struct {
		name    string
		key     *ecdsa.PrivateKey
		wantErr bool
		errType error
	}{
		{
			name:    "nil private key",
			key:     nil,
			wantErr: true,
			errType: ErrPrivateKeyNil,
		},
		{
			name:    "private key with nil D",
			key:     &ecdsa.PrivateKey{D: nil},
			wantErr: true,
			errType: ErrPrivateKeyNil,
		},
		{
			name: "zero private key",
			key: &ecdsa.PrivateKey{
				D: big.NewInt(0),
			},
			wantErr: true,
			errType: ErrPrivateKeyZero,
		},
		{
			name: "too small private key",
			key: &ecdsa.PrivateKey{
				D: big.NewInt(999), // Less than minPrivateKey (1000)
			},
			wantErr: true,
			errType: ErrPrivateKeyTooSmall,
		},
		{
			name: "valid private key (minimum acceptable)",
			key: &ecdsa.PrivateKey{
				D:         big.NewInt(1000),
				PublicKey: ecdsa.PublicKey{Curve: elliptic.P256()},
			},
			wantErr: false,
		},
		{
			name: "valid private key (typical value)",
			key: &ecdsa.PrivateKey{
				D:         new(big.Int).SetBytes([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}),
				PublicKey: ecdsa.PublicKey{Curve: elliptic.P256()},
			},
			wantErr: false,
		},
		{
			name: "private key exceeding curve order",
			key: &ecdsa.PrivateKey{
				D: new(big.Int).Add(curveOrder, big.NewInt(1)),
			},
			wantErr: true,
			errType: ErrPrivateKeyTooLarge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePrivateKey(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != tt.errType {
				t.Errorf("ValidatePrivateKey() error = %v, want %v", err, tt.errType)
			}
		})
	}
}

func TestValidatePrivateKey_RealKey(t *testing.T) {
	// Test with a real generated key
	privateKeyStr := "0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"
	privateKey, err := PrivateKey(privateKeyStr)
	if err != nil {
		t.Fatalf("Failed to create private key: %v", err)
	}

	err = ValidatePrivateKey(privateKey)
	if err != nil {
		t.Errorf("Valid real private key failed validation: %v", err)
	}
}
