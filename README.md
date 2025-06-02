# Layer GameHub Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/u2u-labs/layerg-gamehub-go.svg)](https://pkg.go.dev/github.com/u2u-labs/layerg-gamehub-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/u2u-labs/layerg-gamehub-go)](https://goreportcard.com/report/github.com/u2u-labs/layerg-gamehub-go)
[![License](https://img.shields.io/github/license/yourname/layerggamehub)](LICENSE)

A Go SDK for interacting with **Layer GameHub** — manage game assets (create, update, delete, fetch) and handle authentication seamlessly across development and production environments, now with configurable timeouts and retry logic.

---

## ✨ Features

- 🔑 Authenticate with `apiKey` + `apiKeyId`, auto-fetch and refresh `accessToken`
- 🌍 Multi-environment support (`Dev`, `Prod`)
- 🎮 Asset management: create, update, delete, fetch
- ⏱ Configurable timeout and retry for all requests
- ⚙ Easy integration, no manual `.env` setup required

---

## 📦 Installation

```bash
go get github.com/u2u-labs/layerg-gamehub-go
```

---

## ⚡ Quick Start

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/u2u-labs/layerg-gamehub-go"
)

func main() {
    client, err := layerggamehub.NewClient(
        "your-api-key",
        "your-api-key-id",
        layerggamehub.Prod,
        &layerggamehub.ClientOptions{
            Timeout: 5 * time.Second,
            Retry:   3,
        },
    )
    if err != nil {
        log.Fatalf("Failed to initialize client: %v", err)
    }

    // Example: Create an asset
    assetInput := layerggamehub.CreateAssetInput{
        Name: "My Awesome Asset",
        // add other fields as required
    }

    if err := client.CreateAsset(assetInput); err != nil {
        log.Fatalf("Failed to create asset: %v", err)
    }
    fmt.Println("Asset created successfully!")

    // Example: Fetch an asset
    asset, err := client.GetAsset("asset-id", "collection-id")
    if err != nil {
        log.Fatalf("Failed to get asset: %v", err)
    }
    fmt.Printf("Fetched asset: %+v\n", asset)
}
```

---

## 🌍 Environment Support

The SDK supports:

- `layerggamehub.Dev` → Development environment
- `layerggamehub.Prod` → Production environment

It automatically sets the correct baseURL based on the environment passed to `NewClient()`.

---

## ⚙ Client Options

When initializing the client, you can provide:

- **Timeout** → overall timeout per HTTP request (default: 10 seconds)
- **Retry** → number of times to retry on failure (default: 1)

Example:

```go
client, err := layerggamehub.NewClient(
    "apiKey",
    "apiKeyId",
    layerggamehub.Dev,
    &layerggamehub.ClientOptions{
        Timeout: 5 * time.Second,
        Retry:   3,
    },
)
```

---

## 📚 API Coverage

### Assets

- `CreateAsset`
- `GetAsset`
- `UpdateAsset`
- `DeleteAsset`

### Collections

- `CreateCollection`
- `UpdateCollection`
- `PublicCollection`

Authentication (login/refresh) is handled internally, with automatic retries.
