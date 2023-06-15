package mxapi

import (
	"bytes"
	"io"
	"net/http"
)

type MxApiRequest struct {
	Host string
}

func (mar MxApiRequest) PostJsonRequest(url string, json []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", mar.Host+url, bytes.NewBuffer(json))
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func (mar MxApiRequest) GetRequest(url string) (string, error) {
	req, err := http.NewRequest("GET", mar.Host+url, nil)
	if err != nil {
		return "", err
	}

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
