package layerggamehub

import (
	"time"
)

//Asset

type Asset struct {
	ID             string         `json:"id"`
	TokenID        string         `json:"tokenId"`
	CollectionID   string         `json:"collectionId"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	MediaStorageID string         `json:"mediaStorageId"`
	MetaDataID     string         `json:"metaDataId"`
	NameSlug       string         `json:"nameSlug"`
	Slug           string         `json:"slug"`
	Quantity       int            `json:"quantity"`
	ApiKeyID       string         `json:"apiKeyID"`
	Media          Media          `json:"media"`
	Metadata       Metadata       `json:"metadata"`
	Image          string         `json:"image"`
	Attributes     []Attribute    `json:"attributes"`
	ExternalURL    string         `json:"external_url"`
	Collection     CollectionInfo `json:"collection"`
}

type CollectionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

type CreateAssetInput struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	TokenID      *string  `json:"tokenId,omitempty"`
	CollectionID string   `json:"collectionId"`
	Quantity     string   `json:"quantity"`
	Media        Media    `json:"media"`
	Metadata     Metadata `json:"metadata"`
}

type UpdateAssetInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TokenID     *string  `json:"tokenId,omitempty"`
	Quantity    string   `json:"quantity"`
	Media       Media    `json:"media"`
	Metadata    Metadata `json:"metadata"`
}

type Media struct {
	S3Url string `json:"S3Url"`
}

type Metadata struct {
	Metadata InnerMetadata `json:"metadata"`
}

type InnerMetadata struct {
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

//Collection

type UpsertCollectionInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatarUrl"`
	ProjectID   string `json:"projectId"`
	SMC         SMC    `json:"smc"`
}

type SMC struct {
	ContractAddress string `json:"contractAddress"`
	ContractType    string `json:"contractType"`
	NetworkID       int    `json:"networkID"`
	TokenSymbol     string `json:"tokenSymbol"`
	TotalSupply     int    `json:"totalSupply"`
}
