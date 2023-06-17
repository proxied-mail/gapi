package papi

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app"
)

var meCacheByAuth map[string]PapiUserStruct = make(map[string]PapiUserStruct)

type PapiUserStruct struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			Username string `json:"username"`
		} `json:"attributes"`
	} `json:"data"`
}

func (userStruct PapiUserStruct) IsAuthenticated() bool {
	return userStruct.Data.Attributes.Username != ""
}

func Me(auth string) (PapiUserStruct, error) {
	Req := PapiRequest{
		Host: app.GetPapiHost(),
	}
	resp, err := Req.GetRequestWithAuth("/api/v1/users/me", auth)
	if err != nil {
		return PapiUserStruct{}, err
	}

	var f = PapiUserStruct{}
	err = json.Unmarshal([]byte(resp), &f)
	if err != nil {
		return PapiUserStruct{}, err
	}

	return f, nil
}

func MeCached(auth string) (PapiUserStruct, error) {
	cacheByAuth, has := meCacheByAuth[auth]
	if has {
		return cacheByAuth, nil
	}
	me, err := Me(auth)
	if err != nil {
		return PapiUserStruct{}, err
	}
	meCacheByAuth[auth] = me
	return me, nil
}
