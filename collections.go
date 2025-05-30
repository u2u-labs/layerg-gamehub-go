package layerggamehub

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) PublicCollection(collectionId string) (*Asset, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/collection/public/%s", c.BaseURL, collectionId)
	req, err := http.NewRequest("POST", url, nil)
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
		return nil, fmt.Errorf("public collection failed: %s", resp.Status)
	}

	var asset Asset
	if err := json.NewDecoder(resp.Body).Decode(&asset); err != nil {
		return nil, err
	}

	return &asset, nil
}
