# Project Context: gowallet

## Project Overview

`gowallet` is a comprehensive CLI tool for generating and managing cryptocurrency wallet components. It focuses on Ethereum compatibility but adheres to broader standards like BIP39 (mnemonics), BIP32 (HD wallets), and BIP44 (derivation paths).

**Key Features:**
- **Private Key Generation:** Secure random generation (CSPRNG).
- **Mnemonic Management:** BIP39 support (12-24 words).
- **HD Wallet:** Hierarchical Deterministic wallet support (BIP32/44).
- **Derivation:** Key and address derivation from paths.
- **Security:** Built-in validation for key strength and entropy; memory zeroing for sensitive data.
- **Cross-Platform:** Builds for Linux, macOS, and Windows (AMD64/ARM64).

## Architecture

The project follows a clean architecture pattern:

*   **Entry Point:** `main.go` - Bootstraps the application.
*   **Interface Layer (`cmd/`):** Handles CLI interactions using [Cobra](https://github.com/spf13/cobra).
    *   `root.go`: Base command configuration.
    *   `genPrivateKey.go`: Single key operations (ETH style).
    *   `hdwallet.go`: HD wallet operations (mnemonics, seeds, paths).
    *   `version.go`: Build version info.
*   **Domain Layer (`internal/`):** Core logic, isolated from the CLI.
    *   `commonPrivateKey/`: Private key generation, ETH address derivation, signing.
    *   `hdwallet/`: BIP39/32/44 implementation (mnemonics, seeds, paths).
    *   `security/`: Memory safety utilities (zeroing sensitive data).

## Building and Running

**Prerequisites:**
*   Go 1.24+
*   `make` (optional, for build scripts)
*   `upx` (optional, for binary compression)

**Key Commands:**

| Action | Command | Description |
| :--- | :--- | :--- |
| **Build (Dev)** | `go build -o gowallet main.go` | Quick build for development. |
| **Build (Prod)** | `make build` | Production build with version/commit/time injection. |
| **Build (Release)**| `make build-release` | Production build + UPX compression. |
| **Run** | `./gowallet` | Run the binary (after building). |
| **Clean** | `make clean` | Remove build artifacts. |

**Example Usage:**

```bash
# Generate a private key
./gowallet genPrivateKey

# Generate a 24-word mnemonic
./gowallet genMnemonic -s 24

# Derive address from mnemonic and path
./gowallet getPath -m "your mnemonic..." -p "m/44'/60'/0'/0/0"
```

## Development Conventions

*   **Language:** Go (Golang).
*   **Style:** Follows standard Go formatting (`go fmt`) and vetting (`go vet`).
*   **Project Layout:** Standard Go project layout (`cmd/`, `internal/`).
*   **Testing:**
    *   Unit tests colocated with source (`*_test.go`).
    *   Integration tests in `cmd/cmd_test.go`.
    *   Run all tests: `go test ./...`
    *   Run with verbose output: `go test -v ./...`
*   **Security:**
    *   Sensitive data (mnemonics, keys) must be zeroed out in memory after use (`security/memory.go`).
    *   No external runtime dependencies (compiled statically).
*   **Validation:** Input validation (key strength, path format, entropy quality) is mandatory for core logic.

## Key Files

*   `main.go`: Application entry point.
*   `cmd/root.go`: Root CLI command definition.
*   `internal/hdwallet/hdwallet.go`: Core HD wallet implementation.
*   `internal/commonPrivateKey/commonPrivateKey.go`: Private key logic.
*   `Makefile`: Build automation scripts.
*   `ARCHITECTURE.md`: Detailed architectural documentation.
