# LayerG Gamehub Go SDK

---

## Installation

```bash
go get github.com/u2u-labs/layerg-gamehub-go
```

---

## Mode

```go
layerggamehub.Sandbox 
layerggamehub.Production
```

---

## Client Initialization & Authentication

**Example:**

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/u2u-labs/layerg-gamehub-go"
)

func main() {
	opts := &layerggamehub.ClientOptions{
		Retry:   3,
		Timeout: 10 * time.Second,
	}

	client, err := layerggamehub.NewClient("apiKey", "apiKeyId", Mode.Sandbox, opts)
	if err != nil {
		log.Fatal("Failed to initialize client:", err)
	}

	authResp, err := client.Authenticate()
	if err != nil {
		log.Fatal("Authentication failed:", err)
	}

	fmt.Println("Authenticated successfully. AccessToken:", authResp.AccessToken)

	// continue to call asset/collection methods here
}
```

---

## Asset

### Methods

#### GetByTokenId

```go
GetByTokenId(tokenId string, collectionId string) (*Asset, error)
```

**Example:**

```go
asset, err := client.Asset.GetByTokenId("TOKEN_ID", "COLLECTION_ID")
if err != nil {
	log.Fatal("Failed to fetch asset:", err)
}
fmt.Println("Asset:", asset)
```

#### Create

```go
Create(input CreateAssetInput) (*Asset, error)
```

**Example:**

```go
input := layerggamehub.CreateAssetInput{
	Name:         "test",
	Description:  "test",
	TokenID:      nil,
	CollectionID: "COLLECTION_ID",
	Quantity:     "1",
	Media:        layerggamehub.Media{S3Url: ""},
	Metadata:     layerggamehub.Metadata{Metadata: layerggamehub.InnerMetadata{Attributes: []layerggamehub.Attribute{}}},
}

asset, err := client.Asset.Create(input)
if err != nil {
	log.Fatal("Failed to create asset:", err)
}
fmt.Println("Created asset:", asset)
```

#### Update

```go
Update(input UpdateAssetInput) (*Asset, error)
```

**Example:**

```go
input := layerggamehub.UpdateAssetInput{
	Data: layerggamehub.UpdateAssetData{
		Name:        "updated name",
		Description: "updated description",
		TokenID:     nil,
		Quantity:    "1",
		Media:       layerggamehub.Media{S3Url: ""},
		Metadata:    layerggamehub.Metadata{Metadata: layerggamehub.InnerMetadata{Attributes: []layerggamehub.Attribute{}}},
	},
	Where: layerggamehub.UpdateAssetWhereInput{
		CollectionId: "COLLECTION_ID",
		AssetId:      "ASSET_ID",
	},
}

updatedAsset, err := client.Asset.Update(input)
if err != nil {
	log.Fatal("Failed to update asset:", err)
}
fmt.Println("Updated asset:", updatedAsset)
```

#### Delete

```go
Delete(collectionId string, tokenId string) (bool, error)
```

**Example:**

```go
ok, err := client.Asset.Delete("COLLECTION_ID", "TOKEN_ID")
if err != nil {
	log.Fatal("Failed to delete asset:", err)
}
if ok {
	fmt.Println("Asset deleted successfully.")
}
```

---

## Collection

### Methods

#### GetById

```go
GetById(collectionId string) (*Collection, error)
```

**Example:**

```go
collection, err := client.Collection.GetById("COLLECTION_ID")
if err != nil {
	log.Fatal("Failed to fetch collection:", err)
}
fmt.Println("Collection:", collection)
```

#### Create

```go
Create(input CreateCollectionInput) (*Collection, error)
```

**Example:**

```go
input := layerggamehub.CreateCollectionInput{
	Name:        "Test Collection",
	Description: "Description here",
	AvatarURL:   "https://example.com/avatar.png",
	ProjectID:   "PROJECT_ID",
	SMC: layerggamehub.SMC{
		ContractAddress: "0x123...",
		ContractType:    "ERC721",
		NetworkID:       1,
		TokenSymbol:     "TEST",
		TotalSupply:     10000,
	},
}

collection, err := client.Collection.Create(input)
if err != nil {
	log.Fatal("Failed to create collection:", err)
}
fmt.Println("Created collection:", collection)
```

#### Update

```go
Update(input UpdateCollectionInput) (*Collection, error)
```

**Example:**

```go
input := layerggamehub.UpdateCollectionInput{
	Data: layerggamehub.UpdateCollectionData{
		Name:        "Updated Name",
		Description: "Updated Description",
		AvatarURL:   "https://example.com/avatar.png",
		ProjectID:   "PROJECT_ID",
		SMC: layerggamehub.SMC{
			ContractAddress: "0x123...",
			ContractType:    "ERC721",
			NetworkID:       1,
			TokenSymbol:     "TEST",
			TotalSupply:     10000,
		},
	},
	Where: layerggamehub.UpdateCollectionWhereInput{
		CollectionID: "COLLECTION_ID",
	},
}

updatedCollection, err := client.Collection.Update(input)
if err != nil {
	log.Fatal("Failed to update collection:", err)
}
fmt.Println("Updated collection:", updatedCollection)
```

#### Public

```go
Public(collectionId string) (*Collection, error)
```

**Example:**

```go
published, err := client.Collection.Public("COLLECTION_ID")
if err != nil {
	log.Fatal("Failed to public a collection:", err)
}
fmt.Println("Public collection success:", published)
```

---

## Error Handling

All SDK methods return `(..., error)`. If `error != nil`, the operation failed.

**Example:**

```go
asset, err := client.Asset.GetByTokenId("TOKEN_ID", "COLLECTION_ID")
if err != nil {
	log.Fatal("Error fetching asset:", err)
}
fmt.Println("Asset ID:", asset.ID)
```

---
