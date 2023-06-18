package domains

import (
	"database/sql"
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/pkg/mxapi"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"math/rand"
	"strings"
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

	mxapiReponseEntity, err := mxapi.CreateNewUserCatchAllRequest(request.Domain, "1")
	if err != nil || !mxapiReponseEntity.IsCreated {
		return emptyModel, errors.New("Error creating domain on MX")
	}
	dkim, err2 := mxapi.RequestDkim(request.Domain)
	if err2 != nil {
		return emptyModel, err2
	}

	model := models.CustomDomain{
		UserId:    userid,
		Domain:    request.Domain,
		Status:    models.DomainStatusNew,
		IsShared:  false,
		IsPremium: false,
		SmtpPassword: sql.NullString{
			String: cds.generateRandomPass(),
			Valid:  true,
		},
		DkimKey:   dkim.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cds.Db.Create(&model)

	return model, nil
}

func (cds CreateDomainService) generateRandomPass() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 20
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
