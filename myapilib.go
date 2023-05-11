package shopify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyAPI struct {
	baseURL string
	apiKey  string
}

func NewMyAPI(baseURL, apiKey string) *MyAPI {
	return &MyAPI{baseURL: baseURL, apiKey: apiKey}
}

func (a *MyAPI) GetResource(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/resources/%s", a.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", a.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
