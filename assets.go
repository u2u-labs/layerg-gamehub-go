package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateAsset(asset Asset) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(asset)
	req, err := http.NewRequest("POST", c.BaseURL+"/assets", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("create asset failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) GetAsset(id string) (*Asset, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/assets/%s", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get asset failed: %s", resp.Status)
	}

	var asset Asset
	if err := json.NewDecoder(resp.Body).Decode(&asset); err != nil {
		return nil, err
	}

	return &asset, nil
}
