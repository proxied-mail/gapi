package mxapi

import "encoding/json"

type CreateNewUserRequest struct {
	Email    string `json:"email"`
	Type     string `json:"type"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
}

type MailboxEntityResponse struct {
	IsCreated bool
	Type      string
	IsMocked  bool
}

func CreateNewUserCatchAllRequest(domain string, pass string) (MailboxEntityResponse, error) {
	reqPayload := CreateNewUserRequest{
		Email:    "catchall@" + domain,
		Type:     "catchall",
		Domain:   domain,
		Password: pass,
	}
	mxAR := MxApiRequest{
		Host: "http://mx.proxiedmail.com:8080",
	}
	jsonReq, _ := json.Marshal(reqPayload)

	resp, err := mxAR.PostJsonRequest("/add-user", jsonReq)

	respEntity := MailboxEntityResponse{}
	if err != nil {
		return respEntity, err
	}

	json.Unmarshal(resp, &respEntity)

	return respEntity, nil
}
