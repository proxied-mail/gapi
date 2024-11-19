package domains

import (
	"database/sql"
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
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

	restrictedDomains := []string{
		"icloud",
		"gmail",
		"outlook",
	}
	domainParts := strings.Split(".", request.Domain)
	domainFirstPart := domainParts[0]

	for i, _ := range restrictedDomains {
		if domainFirstPart == restrictedDomains[i] {
			return emptyModel, errors.New("Domain is restricted")
		}
	}

	condition := models.CustomDomain{
		Domain: request.Domain,
	}
	var count int64
	cds.Db.Model(condition).Where(condition).Count(&count)

	if count > 0 {
		return emptyModel, errors.New("Domain is already exists")
	}

	pass := cds.generateRandomPass()

	model := models.CustomDomain{
		UserId:    userid,
		Domain:    request.Domain,
		Status:    models.DomainStatusNew,
		IsShared:  false,
		IsPremium: false,
		SmtpPassword: sql.NullString{
			String: pass,
			Valid:  true,
		},
		DkimKey:   "",
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
