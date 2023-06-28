package common

import "encoding/json"

type Success struct {
	Status bool `json:"status"`
}

func GetSuccess() string {
	resp, _ := json.Marshal(Success{true})
	return string(resp)
}
