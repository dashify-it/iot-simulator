package simulator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var ApiClient *http.Client

func InitApiClient(config Config) {
	ApiClient = &http.Client{Timeout: 10 * time.Second}
}

func SendApiRequest(config Config, body interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", config.API.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(config.API.ApiKeyHeaderName, config.API.Key)

	resp, err := ApiClient.Do(req)
	if err != nil {
		return err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(responseBody))
	resp.Body.Close()
	return nil
}
