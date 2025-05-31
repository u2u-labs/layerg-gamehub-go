# LayerG GameHub Go SDK

A Go SDK for interacting with **LayerG GameHub**.

---

## Features

- Authenticate with `apiKey` + `apiKeyId`, auto-fetch and refresh `accessToken`
- Multi-environment support (`Dev`, `Prod`)
- Asset management: create, update, delete, fetch
- Collection management: create, update, public collection

---

## Installation

```bash
go get github.com/u2u-labs/layerg-gamehub-go
```

---

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/u2u-labs/layerg-gamehub-go"
)

func main() {
    // Initialize client in dev
    client, err := layerggamehub.NewClient("your-api-key", "your-api-key-id", layerggamehub.Dev)
    if err != nil {
        log.Fatalf("Failed to initialize client: %v", err)
    }

    // Example: Create an asset
    assetInput := layerggamehub.CreateAssetInput{
        Name: "My Asset",
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

## Environment Support

The SDK supports:

- `layerggamehub.Dev` → Development environment
- `layerggamehub.Prod` → Production environment

---

## API Coverage

### Asset

- `CreateAsset`
- `GetAsset`
- `UpdateAsset`
- `DeleteAsset`

### Collection

- `CreateCollection`
- `UpdateCollection`
- `PublicCollection`

Authentication (login/refresh) is handled internally.
