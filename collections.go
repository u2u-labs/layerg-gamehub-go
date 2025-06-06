package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *CollectionModule) GetById(collectionId string) (*Collection, error) {
	url := fmt.Sprintf("%s/collection/%s", c.BaseURL, collectionId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken())

	resp, err := c.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get collection failed: %s", resp.Status)
	}

	var collection Collection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}

	return &collection, nil
}

func (c *CollectionModule) Create(input CreateCollectionInput) (*Collection, error) {
	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", c.BaseURL+"/collection", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var collection Collection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create collection failed: %s", resp.Status)
	}

	return &collection, nil
}

func (c *CollectionModule) Update(input UpdateCollectionInput) (*Collection, error) {
	body, _ := json.Marshal(input.Data)
	url := fmt.Sprintf("%s/collection/%s", c.BaseURL, input.Where.CollectionID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var collection Collection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update collection failed: %s", resp.Status)
	}

	return &collection, nil
}

func (c *CollectionModule) Public(collectionId string) (*Collection, error) {
	url := fmt.Sprintf("%s/collection/public/%s", c.BaseURL, collectionId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken())

	resp, err := c.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var collection Collection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("public collection failed: %s", resp.Status)
	}

	return &collection, nil
}
