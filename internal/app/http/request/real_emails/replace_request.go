package real_emails

type ReplaceRealEmailRequest struct {
	NewEmail string `json:"newEmail" validate:"required"`
	OldEmail string `json:"oldEmail" validate:"required"`
}
