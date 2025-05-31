package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateCollection(input UpsertCollectionInput) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", c.BaseURL+"/collection", bytes.NewBuffer(body))
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
		return fmt.Errorf("create collection failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) UpdateCollection(input UpsertCollectionInput, collectionId string) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(input)
	url := fmt.Sprintf("%s/collection/%s", c.BaseURL, collectionId)
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
		return fmt.Errorf("update collection failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) PublicCollection(collectionId string) (bool, error) {
	if err := c.ensureAccessToken(); err != nil {
		return false, err
	}

	url := fmt.Sprintf("%s/collection/public/%s", c.BaseURL, collectionId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("public collection failed: %s", resp.Status)
	}

	return true, nil
}
