package layerggamehub

import (
	"fmt"
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
	Retry        int

	mu sync.Mutex
}

type ClientOptions struct {
	Timeout time.Duration // e.g., 10 * time.Second
	Retry   int           // e.g., 3 retries
}

func NewClient(apiKey, apiKeyId string, env Environment, opts *ClientOptions) (*Client, error) {
	timeout := 10 * time.Second
	if opts != nil && opts.Timeout > 0 {
		timeout = opts.Timeout
	}

	c := &Client{
		ApiKey:     apiKey,
		ApiKeyId:   apiKeyId,
		BaseURL:    GetBaseURL(env),
		HTTPClient: &http.Client{Timeout: timeout},
	}

	if opts != nil {
		c.Retry = opts.Retry
	}

	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) DoWithRetry(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	retries := 1
	if c.Retry > 0 {
		retries = c.Retry
	}

	for i := 0; i < retries; i++ {
		if err := c.ensureAccessToken(); err != nil {
			return nil, err
		}
	}

	for i := 0; i < retries; i++ {
		resp, err = c.HTTPClient.Do(req)

		if err == nil {
			return resp, nil
		}

		// If error is not connection error, return without retry
		if !isConnectionError(err) {
			return nil, err
		}
	}

	return nil, fmt.Errorf("request failed after %d retries: %w", retries, err)
}
