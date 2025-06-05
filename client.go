package layerggamehub

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type AssetModule struct {
	BaseURL        string
	GetAccessToken func() string
	DoWithRetry    func(*http.Request) (*http.Response, error)
}

type CollectionModule struct {
	BaseURL        string
	GetAccessToken func() string
	DoWithRetry    func(*http.Request) (*http.Response, error)
}

type Client struct {
	APIKey             string
	APIKeyID           string
	AccessToken        string
	RefreshToken       string
	AccessTokenExpire  int64
	RefreshTokenExpire int64
	BaseURL            string
	HTTPClient         *http.Client
	Asset              *AssetModule
	Collection         *CollectionModule
	Retry              int

	mu sync.Mutex
}

type ClientOptions struct {
	Timeout time.Duration // e.g., 10 * time.Second
	Retry   int           // e.g., 3 retries
}

// NewClient creates a new Client instance with optional configurations.
func NewClient(apiKey, apiKeyID string, env Environment, opts *ClientOptions) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API key cannot be empty")
	}
	if apiKeyID == "" {
		return nil, errors.New("API key ID cannot be empty")
	}

	timeout := 10 * time.Second
	retry := 1

	if opts != nil {
		if opts.Timeout > 0 {
			timeout = opts.Timeout
		}
		if opts.Retry > 0 {
			retry = opts.Retry
		}
	}

	client := &Client{
		APIKey:     apiKey,
		APIKeyID:   apiKeyID,
		BaseURL:    GetBaseURL(env),
		HTTPClient: &http.Client{Timeout: timeout},
		Retry:      retry,
	}

	client.Asset = &AssetModule{
		BaseURL:        client.BaseURL,
		GetAccessToken: func() string { return client.AccessToken },
		DoWithRetry:    client.DoWithRetry,
	}

	client.Collection = &CollectionModule{
		BaseURL:        client.BaseURL,
		GetAccessToken: func() string { return client.AccessToken },
		DoWithRetry:    client.DoWithRetry,
	}

	return client, nil
}

// DoWithRetry performs an HTTP request with retry logic.
func (c *Client) DoWithRetry(req *http.Request) (*http.Response, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, fmt.Errorf("failed to ensure access token: %w", err)
	}

	var resp *http.Response
	var err error

	for attempt := 1; attempt <= c.Retry; attempt++ {
		resp, err = c.HTTPClient.Do(req)
		if err == nil {
			return resp, nil
		}
		if !shouldRetry(resp, err) {
			return nil, err
		}
	}

	return nil, fmt.Errorf("request failed after %d retries: %w", c.Retry, err)
}
