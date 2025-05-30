# LayerG GameHub Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/yourname/layerggamehub.svg)](https://pkg.go.dev/github.com/yourname/layerggamehub)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourname/layerggamehub)](https://goreportcard.com/report/github.com/yourname/layerggamehub)
[![License](https://img.shields.io/github/license/yourname/layerggamehub)](LICENSE)

A Go SDK for interacting with **LayerG GameHub** — manage game assets, collections, and publish content easily via Go.

---

## ✨ Features

- 🔑 Authenticate with `apiKey` + `apiKeyId`, auto-fetch and refresh `accessToken`
- 🎮 Manage **Assets**: create, get, update, delete
- 📦 Manage **Collections**: create, get, update, delete
- 🚀 Publish / remove content to GameHub
- ⚙ Configurable backend `baseURL` via `.env`

---

## 📦 Installation

```bash
go get github.com/yourname/layerggamehub
```

---

## ⚡ Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourname/layerggamehub"
)

func main() {
    client, err := layerggamehub.NewClient("your-api-key", "your-api-key-id")
    if err != nil {
        log.Fatalf("Failed to initialize client: %v", err)
    }

    // Create an asset
    asset := layerggamehub.Asset{
        ID:   "asset-123",
        Description: "Test Asset",
        ...rest
    }
    if err := client.CreateAsset(asset); err != nil {
        log.Fatalf("Failed to create asset: %v", err)
    }
    fmt.Println("Asset created successfully!")
}
```
