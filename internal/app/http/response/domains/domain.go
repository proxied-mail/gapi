package domains

import (
	"github.com/abrouter/gapi/internal/app/models"
	"time"
)

type DomainResponse struct {
	UserId           int       `json:"user_id"`
	Domain           string    `json:"domain"`
	Status           int       `json:"status"`
	IsShared         bool      `json:"is_shared"`
	IsPremium        bool      `json:"IsPremium"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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

func MapResponse(domain models.CustomDomain) *DomainResponse {
	return &DomainResponse{
		domain.UserId,
		domain.Domain,
		domain.Status,
		domain.IsShared,
		domain.IsPremium,
		domain.CreatedAt,
		domain.UpdatedAt,
		domain,
		"",
		"",
	}
}

func MapResponseList(domains []models.CustomDomain) []*DomainResponse {
	var model models.CustomDomain
	var newMap []*DomainResponse
	for _, model = range domains {
		newMap = append(newMap, MapResponse(model))
	}

	return newMap
}
