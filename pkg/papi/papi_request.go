package papi

import (
	"io"
	"net/http"
)

type PapiRequest struct {
	Host string
}

func (pr PapiRequest) GetRequestWithAuth(url string, auth string) (string, error) {
	req, err := http.NewRequest("GET", pr.Host+url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
