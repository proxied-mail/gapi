package domains

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type CreateDomainRequest struct {
	Domain string `json:"domain"`
}

type CreateDomainService struct {
	fx.In
	Db *gorm.DB
}

func (cds CreateDomainService) CreateDomain(
	userid int,
	request CreateDomainRequest,
) (models.CustomDomain, error) {

	emptyModel := models.CustomDomain{}
	if request.Domain == "" {
		return emptyModel, errors.New("Empty domain")
	}

	condition := models.CustomDomain{
		Domain: request.Domain,
	}
	var count int64
	cds.Db.Model(condition).Where(condition).Count(&count)

	if count > 0 {
		return emptyModel, errors.New("Domain is already exists")
	}

	model := models.CustomDomain{
		UserId:    userid,
		Domain:    request.Domain,
		Status:    models.DomainStatusNew,
		IsShared:  false,
		IsPremium: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cds.Db.Create(&model)

	return model, nil
}
