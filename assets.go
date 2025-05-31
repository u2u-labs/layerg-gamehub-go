package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAsset(assetId string, collectionId string) (*Asset, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/assets/%s/%s", c.BaseURL, collectionId, assetId)
	req, err := http.NewRequest("GET", url, nil)
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

func (c *Client) CreateAsset(input CreateAssetInput) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", c.BaseURL+"/assets/create", bytes.NewBuffer(body))
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create asset failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) UpdateAsset(input UpdateAssetInput, collectionId string, assetId string) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(input)
	url := fmt.Sprintf("%s/assets/%s/%s", c.BaseURL, collectionId, assetId)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
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
		return fmt.Errorf("update asset failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) DeleteAsset(collectionId string, assetId string) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	url := fmt.Sprintf("%s/assets/%s/%s", c.BaseURL, collectionId, assetId)
	req, err := http.NewRequest("DELETE", url, nil)
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
		return fmt.Errorf("delete asset failed: %s", resp.Status)
	}

	return nil
}
