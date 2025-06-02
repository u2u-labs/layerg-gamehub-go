package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type authResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *Client) authenticate() error {
	payload := map[string]string{
		"apiKey":   c.ApiKey,
		"apiKeyID": c.ApiKeyId,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", c.BaseURL+"/auth/login", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("auth failed with status: %s", resp.Status)
	}

	var authResp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.AccessToken = authResp.AccessToken
	c.RefreshToken = authResp.RefreshToken
	return nil
}

func (c *Client) ensureAccessToken() error {
	return c.refreshAccessToken()
}

func (c *Client) refreshAccessToken() error {
	payload := map[string]string{
		"refreshToken": c.RefreshToken,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", c.BaseURL+"/auth/refresh", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("refresh failed with status: %s", resp.Status)
	}

	var authResp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.AccessToken = authResp.AccessToken
	c.RefreshToken = authResp.RefreshToken
	return nil
}
