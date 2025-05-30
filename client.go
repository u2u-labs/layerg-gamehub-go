package layerggamehub

import (
	"net/http"
	"sync"
	"time"
)

type Client struct {
	ApiKey       string
	ApiKeyId     string
	AccessToken  string
	RefreshToken string
	BaseURL      string
	HTTPClient   *http.Client

	mu sync.Mutex
}

func NewClient(apiKey, apiKeyId string, env Environment) (*Client, error) {
	c := &Client{
		ApiKey:     apiKey,
		ApiKeyId:   apiKeyId,
		BaseURL:    GetBaseURL(env),
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}
