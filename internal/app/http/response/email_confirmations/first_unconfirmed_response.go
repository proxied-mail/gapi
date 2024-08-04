package email_confirmations

type FirstUnconfirmedResponse struct {
	HasUnconfirmedNotShown bool   `json:"has_unconfirmed_not_shown"`
	Id                     string `json:"id"`
	ContinueChecking       bool   `json:"continue_checking"`
}
