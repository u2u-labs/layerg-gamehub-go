package layerggamehub

import (
	"time"
)

// Asset

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

type UpdateAssetData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TokenID     *string  `json:"tokenId,omitempty"`
	Quantity    string   `json:"quantity"`
	Media       Media    `json:"media"`
	Metadata    Metadata `json:"metadata"`
}

type UpdateAssetWhereInput struct {
	AssetId      string `json:"assetId"`
	CollectionId string `json:"collectionId"`
}
type UpdateAssetInput struct {
	Data  UpdateAssetData       `json:"data"`
	Where UpdateAssetWhereInput `json:"where"`
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

// Collection

type Collection struct {
	ID            string          `json:"id"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	TotalAssets   int             `json:"totalAssets"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	AvatarURL     string          `json:"avatarUrl"`
	ProjectID     string          `json:"projectId"`
	NameSlug      string          `json:"nameSlug"`
	Slug          string          `json:"slug"`
	APIKeyID      string          `json:"apiKeyID"`
	IsPublic      bool            `json:"isPublic"`
	Project       Project         `json:"project"`
	SmartContract []SmartContract `json:"SmartContract"`
}

type CreateCollectionInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatarUrl"`
	ProjectID   string `json:"projectId"`
	SMC         SMC    `json:"smc"`
}

type UpdateCollectionData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatarUrl"`
	ProjectID   string `json:"projectId"`
	SMC         SMC    `json:"smc"`
}

type UpdateCollectionWhereInput struct {
	CollectionID string `json:"collectionId"`
}

type UpdateCollectionInput struct {
	Data  UpdateCollectionData       `json:"data"`
	Where UpdateCollectionWhereInput `json:"where"`
}

type SMC struct {
	ContractAddress string `json:"contractAddress"`
	ContractType    string `json:"contractType"`
	NetworkID       int    `json:"networkID"`
	TokenSymbol     string `json:"tokenSymbol"`
	TotalSupply     int    `json:"totalSupply"`
}

// Project
type Project struct {
	ID           string       `json:"id"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	IsEnabled    bool         `json:"isEnabled"`
	CountFav     int          `json:"countFav"`
	Platform     []string     `json:"platform"`
	TotalCls     int          `json:"totalCls"`
	Name         string       `json:"name"`
	GameIcon     string       `json:"gameIcon"`
	Banner       string       `json:"banner"`
	APIKeyID     string       `json:"apiKeyID"`
	Telegram     string       `json:"telegram"`
	Facebook     string       `json:"facebook"`
	Instagram    string       `json:"instagram"`
	Discord      string       `json:"discord"`
	Twitter      string       `json:"twitter"`
	NameSlug     string       `json:"nameSlug"`
	Avatar       string       `json:"avatar"`
	Description  string       `json:"description"`
	Information  string       `json:"information"`
	Policy       string       `json:"policy"`
	Version      string       `json:"version"`
	SlideShow    []string     `json:"slideShow"`
	TotalReview  int          `json:"totalReview"`
	TotalRating  int          `json:"totalRating"`
	Slug         string       `json:"slug"`
	IsRcm        bool         `json:"isRcm"`
	UserID       *string      `json:"userId"` // nullable
	Mode         *string      `json:"mode"`   // nullable
	Index        *int         `json:"index"`  // nullable
	PlatformLink PlatformLink `json:"platformLink"`
}

type PlatformLink struct {
	IOS     string `json:"iOS"`
	MacOS   string `json:"macOS"`
	Android string `json:"android"`
	Windows string `json:"windows"`
}

type SmartContract struct {
	ID              string     `json:"id"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	ContractAddress string     `json:"contractAddress"`
	ContractType    string     `json:"contractType"`
	NetworkID       int        `json:"networkID"`
	ContractName    string     `json:"contractName"`
	TokenSymbol     string     `json:"tokenSymbol"`
	TotalSupply     *int       `json:"totalSupply"` // nullable
	CollectionID    string     `json:"collectionId"`
	DeployedAt      *time.Time `json:"deployedAt"` // nullable
	NameSlug        string     `json:"nameSlug"`
}
