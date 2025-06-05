package layerggamehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *AssetModule) GetByTokenId(tokenId string, collectionId string) (*Asset, error) {
	url := fmt.Sprintf("%s/assets/%s/%s", a.BaseURL, collectionId, tokenId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.GetAccessToken())

	resp, err := a.DoWithRetry(req)
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

func (a *AssetModule) Create(input CreateAssetInput) (*Asset, error) {
	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", a.BaseURL+"/assets/create", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var asset Asset
	if err := json.NewDecoder(resp.Body).Decode(&asset); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create asset failed: %s", resp.Status)
	}

	return &asset, nil
}

func (a *AssetModule) Update(input UpdateAssetInput) (*Asset, error) {
	body, _ := json.Marshal(input.Data)
	url := fmt.Sprintf("%s/assets/%s/%s", a.BaseURL, input.Where.CollectionId, input.Where.AssetId)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.DoWithRetry(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var asset Asset
	if err := json.NewDecoder(resp.Body).Decode(&asset); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update asset failed: %s", resp.Status)
	}

	return &asset, nil
}

func (a *AssetModule) Delete(collectionId string, tokenId string) (bool, error) {
	url := fmt.Sprintf("%s/assets/%s/%s", a.BaseURL, collectionId, tokenId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+a.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.DoWithRetry(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("delete asset failed: %s", resp.Status)
	}

	return true, nil
}
