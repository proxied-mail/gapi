package domains

import (
	"github.com/abrouter/gapi/internal/app/models"
	"time"
)

type DomainResponse struct {
	UserId           int       `json:"user_id"`
	Domain           string    `json:"domain"`
	Status           int       `json:"status"`
	IsShared         bool      `json:"isShared"`
	IsPremium        bool      `json:"isPremium"`
	DkimKey          string    `json:"dkimKey"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	model            models.CustomDomain
	VerificationHash string `json:"verification_hash"`
	Spf              string `json:"spf"`
}

func (dr DomainResponse) GetModel() *models.CustomDomain {
	return &dr.model
}

func (dr DomainResponse) SetVerificationHash(hash string) {
	dr.VerificationHash = hash
}

func MapResponse(userModel models.UserModel, domain models.CustomDomain) *DomainResponse {
	dkimKey := domain.DkimKey
	if userModel.Id != domain.UserId {
		dkimKey = ""
	}

	return &DomainResponse{
		0,
		domain.Domain,
		domain.Status,
		domain.IsShared,
		domain.IsPremium,
		dkimKey,
		domain.CreatedAt,
		domain.UpdatedAt,
		domain,
		"",
		"",
	}
}

func MapResponseList(userModel models.UserModel, domains []models.CustomDomain) []*DomainResponse {
	var model models.CustomDomain
	var newMap []*DomainResponse
	for _, model = range domains {
		newMap = append(newMap, MapResponse(userModel, model))
	}

	return newMap
}
