package security

import "math/big"

// ZeroBytes securely overwrites a byte slice with zeros
//
// This function clears sensitive data from memory after use.
// Note: Go's garbage collector doesn't guarantee immediate memory clearing,
// so we explicitly zero out sensitive data.
func ZeroBytes(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

// ZeroBigInt securely overwrites a big.Int with zero
//
// This is useful for clearing large numbers like private keys.
func ZeroBigInt(n *big.Int) {
	if n != nil {
		n.SetInt64(0)
	}
}
