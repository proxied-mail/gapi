package mxapi

import (
	"encoding/json"
	"errors"
)

type DkimResponse struct {
	Content string
	Message string
	Error   int
}

func RequestDkim(domain string) (DkimResponse, error) {
	mxapi := MxApiRequest{
		Host: "http://mx.proxiedmail.com:8080",
	}
	response, err := mxapi.GetRequest("/dkim/?domain=" + domain)
	if err != nil {
		return DkimResponse{}, err
	}
	dkimResponse := DkimResponse{}
	json.Unmarshal([]byte(response), &dkimResponse)

	if dkimResponse.Error > 0 {
		return DkimResponse{}, errors.New(dkimResponse.Message)
	}

	return dkimResponse, nil
}
