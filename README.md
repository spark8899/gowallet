# gowallet

A comprehensive cryptocurrency wallet CLI tool for generating keys, addresses, and mnemonics for Ethereum and other cryptocurrencies.

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

- 🔑 **Private Key Generation**: Generate secure random private keys using crypto/rand
- 📝 **BIP39 Mnemonic**: Create and manage mnemonic phrases (12/15/18/21/24 words)
- 🌳 **HD Wallet**: Support for BIP32/BIP44 hierarchical deterministic wallets
- 🔄 **Key Derivation**: Derive keys and addresses from derivation paths
- 🛡️ **Security Validation**: Built-in key strength, entropy quality, and path validation
- 🧹 **Memory Safety**: Automatic zeroing of sensitive data after use
- ⚡ **Fast & Lightweight**: Zero external runtime dependencies
- 🎯 **Cross-Platform**: Supports Linux, macOS, and Windows (including ARM)

## Architecture

`gowallet` follows a clean architecture with three main layers:

- **Entry Point** (`main.go`): Bootstraps the application
- **Interface Layer** (`cmd/`): CLI commands built with [Cobra](https://github.com/spf13/cobra)
- **Domain Layer** (`internal/`): Core cryptographic and wallet logic

For detailed architecture information, see [ARCHITECTURE.md](ARCHITECTURE.md).

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/spark8899/gowallet.git
cd gowallet

# Build using Make (recommended)
make build

# Or build directly with go
go build -o gowallet main.go
```

### Cross-Platform Builds

```bash
# Linux AMD64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gowallet_linux_amd64 main.go

# Linux ARM64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o gowallet_linux_arm64 main.go

# macOS AMD64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gowallet_darwin_amd64 main.go

# macOS ARM64 (Apple Silicon)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o gowallet_darwin_arm64 main.go

# Windows AMD64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gowallet_windows_amd64.exe main.go

# Windows ARM64
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o gowallet_windows_arm64.exe main.go
```

## Usage

### Private Key Operations

#### Generate Private Keys

```bash
# Generate a single private key
./gowallet genPrivateKey
# Output: 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947:0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9

# Generate multiple private keys
./gowallet genPrivateKey 10
# Or using flag
./gowallet genPrivateKey -n 10
```

#### Get Public Key from Private Key

```bash
./gowallet getPublicKey 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
# Or using flag
./gowallet getPublicKey -k 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
```

#### Get Address from Private Key

```bash
./gowallet getAddress 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
# Output: 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
```

### HD Wallet Operations

#### Generate Mnemonic

```bash
# Generate 12-word mnemonic (default)
./gowallet genMnemonic
# Output: tag volcano eight thank tide danger coast health above argue embrace heavy

# Generate 24-word mnemonic
./gowallet genMnemonic 24
# Or using flag
./gowallet genMnemonic -s 24
```

#### Get Seed from Mnemonic

```bash
./gowallet mnToSeed "tag volcano eight thank tide danger coast health above argue embrace heavy"
# Or using flag
./gowallet mnToSeed -m "tag volcano eight thank tide danger coast health above argue embrace heavy"
# Output: efea201152e37883bdabf10b28fdac9c146f80d2e161a544a7079d2ecc4e65948a0d74e47e924f26bf35aaee72b24eb210386bcb1deda70ded202a2b7d1a8c2e
```

#### Derive Keys from Derivation Path

```bash
# From mnemonic
./gowallet getPath -m "tag volcano eight thank tide danger coast health above argue embrace heavy" -p "m/44'/60'/0'/0/0"
# Output: 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947:0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9

# From seed
./gowallet getPath -s "efea201152e37883bdabf10b28fdac9c146f80d2e161a544a7079d2ecc4e65948a0d74e47e924f26bf35aaee72b24eb210386bcb1deda70ded202a2b7d1a8c2e" -p "m/44'/60'/0'/0/0"
```

### Help Commands

```bash
# General help
./gowallet --help

# Command-specific help
./gowallet genPrivateKey help
./gowallet genMnemonic --help
```

### Version Information

```bash
./gowallet version
# Output includes version, git commit, and build time
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run specific package tests
go test ./internal/hdwallet -v
go test ./internal/commonPrivateKey -v

# Run benchmarks
go test ./internal/hdwallet -bench=.
```

### Code Quality

```bash
# Run linter
go vet ./...

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy
```

### Build Commands

```bash
# Development build
go build -o gowallet main.go

# Production build with version info
make build

# Clean build artifacts
make clean
```

## Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [btcsuite/btcd](https://github.com/btcsuite/btcd) - Bitcoin utilities for HD wallets
- [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) - Ethereum cryptography
- [tyler-smith/go-bip39](https://github.com/tyler-smith/go-bip39) - BIP39 implementation
- [tyler-smith/go-bip32](https://github.com/tyler-smith/go-bip32) - BIP32 implementation

## Security Notice

⚠️ **WARNING**: Private keys and mnemonics generated by this tool should be kept secure. Never share them with anyone or commit them to version control.

- This tool is intended for development and testing purposes
- For production use, ensure you're running on a secure, offline machine
- Always verify addresses before sending real funds
- Back up your mnemonics securely

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by various cryptocurrency wallet implementations
- Built with Go's excellent cryptography libraries
- Thanks to the Bitcoin and Ethereum communities for standards like BIP39/BIP32/BIP44

## Package Update (for maintainers)

```bash
# Update Go package proxy
curl -X POST https://proxy.golang.org/github.com/spark8899/gowallet/@v/v1.0.0.info
curl -X POST https://pkg.go.dev/fetch/github.com/spark8899/gowallet@v1.0.0
```
