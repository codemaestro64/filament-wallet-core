# wallet-core

**wallet-core** is a lightweight Go library providing core wallet primitives for Filament Labs projects — including address parsing, cryptographic helpers, persistent key storage, RPC utilities, and a high-level manager that ties everything together.

---

## Features

- Filecoin & hex address parsing/validation
- Cryptographic helpers (key management, signing, verifying)
- Simple on-disk key/value DB wrapper
- JSON-RPC client abstraction
- `Manager` layer that orchestrates storage, crypto, and RPC calls

---

## Repository Layout

```
address.go      # address parsing, validation, formatting
crypto.go       # crypto helpers: key handling, signing, verifying
db.go           # simple local database wrapper
manager.go      # high-level manager for wallet operations
rpcclient.go    # JSON-RPC client utilities
types.go        # core shared types
go.mod
LICENSE-APACHE
LICENSE-MIT
README.md
```

---

## Installation

```bash
go get github.com/filament-labs/wallet-core@latest
```

---

## Quick Start

```go
package main

import (
    "fmt"
    walletcore "github.com/filament-labs/wallet-core"
)

func main() {
    // Example: initialize manager (adjust constructor name to match actual source)
    m, err := walletcore.NewManager("./data")
    if err != nil {
        panic(err)
    }
    defer m.Close()

    // Parse an address
    addr, err := walletcore.ParseAddress("f1...")
    if err != nil {
        fmt.Println("invalid address:", err)
        return
    }

    fmt.Println("parsed address:", addr.String())

    // Sign a message
    sig, err := m.SignMessage(addr, []byte("hello wallet-core"))
    if err != nil {
        panic(err)
    }
    fmt.Printf("signature: %x
", sig)
}
```

> Replace function names with the exact ones from the library if they differ.

---

## High-Level API Overview

### Address Utilities (`address.go`)
- `ParseAddress(string) (Address, error)`
- `Address.String()` — canonical formatting

### Crypto Helpers (`crypto.go`)
- Key generation & management
- `Sign(message)` / `Verify(signature)`

### DB Layer (`db.go`)
- File-backed key/value storage for keys and wallet metadata
- Typically used internally by `Manager`

### RPC Client (`rpcclient.go`)
- Lightweight JSON-RPC wrapper
- `Call(method string, params ...any) (any, error)`

### Manager (`manager.go`)
- High-level orchestration
- Manages:
  - Address + key creation
  - Signing operations
  - RPC interactions
  - Persistence via the DB

---

## Example: JSON-RPC Usage

```go
client := walletcore.NewRPCClient("https://node.example/rpc")

resp, err := client.Call("Wallet.GetBalance", addr.String())
if err != nil {
    panic(err)
}

fmt.Println("balance:", resp)
```

---

## Contributing

1. Fork the repo
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Write tests where appropriate
4. Run linters: `go fmt`, `go vet`
5. Submit a pull request

---

## Security Notes

- Private keys should be handled and stored carefully.
- Avoid using raw private keys in production systems — consider hardware wallets or OS-backed keystores.
- Crypto-related contributions should include tests and clear documentation.

---

## License

This project is dual-licensed under **MIT** and **Apache-2.0**.

See the included license files for full details.

