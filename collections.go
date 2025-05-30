package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateCollection(coll Collection) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	body, _ := json.Marshal(coll)
	req, err := http.NewRequest("POST", c.BaseURL+"/collections", bytes.NewBuffer(body))
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
		return fmt.Errorf("create collection failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) GetCollection(id string) (*Collection, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/collections/%s", c.BaseURL, id), nil)
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
		return nil, fmt.Errorf("get collection failed: %s", resp.Status)
	}

	var coll Collection
	if err := json.NewDecoder(resp.Body).Decode(&coll); err != nil {
		return nil, err
	}

	return &coll, nil
}
