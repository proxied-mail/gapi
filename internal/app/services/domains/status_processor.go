package domains

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/abrouter/gapi/internal/app/http/response/domains"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net"
	"strconv"
	"strings"
)

type StatusProcessorService struct {
	fx.In
	Db *gorm.DB
}

func (sps StatusProcessorService) ProcessStatus(cd []*domains.DomainResponse) []*domains.DomainResponse {
	var domain *domains.DomainResponse
	for _, domain = range cd {

		sps.assignVerificationHash(domain)
		sps.assignSpf(domain)

		if domain.Status == models.DomainStatusNew {
			domain.Status = sps.checkOwnership(domain)
		}
		if domain.Status == models.DomainStatusOwnershipVerified {
			domain.Status = sps.checkMx(domain)
		}
		if domain.Status == models.DomainStatusMxSet {
			domain.Status = sps.checkSpf(domain)
		}
		if domain.Status == models.DomainStatusSpfSet {
			domain.Status = sps.checkDkim(domain)
		}

	}

	return cd
}

func (sps StatusProcessorService) assignVerificationHash(domain *domains.DomainResponse) {
	domain.VerificationHash = GetMD5Hash("hash" + strconv.Itoa(domain.UserId) + "domain-verification")
}

func (sps StatusProcessorService) assignSpf(domain *domains.DomainResponse) {
	domain.Spf = "v=spf1 include:proxiedmail.com ~all"
}

func (sps StatusProcessorService) checkMx(domain *domains.DomainResponse) int {
	mxrc, _ := net.LookupMX(domain.Domain)
	for _, mx := range mxrc {
		if mx.Host == "mx.proxiedmail.com." {

			model := domain.GetModel()
			model.Status = models.DomainStatusMxSet
			sps.Db.Save(&model)
			return model.Status
		}
	}
	return domain.Status
}

func (sps StatusProcessorService) checkDkim(domain *domains.DomainResponse) int {
	txts, _ := net.LookupTXT(domain.Domain)
	for _, txt := range txts {

		if txt == domain.DkimKey {

			model := domain.GetModel()
			model.Status = models.DomainStatusDkimSet
			sps.Db.Save(&model)
			return model.Status
		}
	}
	return domain.Status
}

func (sps StatusProcessorService) checkSpf(domain *domains.DomainResponse) int {
	txts, _ := net.LookupTXT(domain.Domain)
	for _, txt := range txts {

		if txt == domain.Spf {

			model := domain.GetModel()
			model.Status = models.DomainStatusSpfSet
			sps.Db.Save(&model)
			return model.Status
		}
	}

	return domain.Status
}

func (sps StatusProcessorService) checkOwnership(domain *domains.DomainResponse) int {
	txts, _ := net.LookupTXT(domain.Domain)
	txtStartWith := "proxiedmail-verification="

	for _, txt := range txts {

		if strings.Contains(txt, txtStartWith) {
			splits := strings.Split(txt, txtStartWith)
			code := splits[1]
			if domain.VerificationHash != code {
				continue
			}

			model := domain.GetModel()
			model.Status = models.DomainStatusOwnershipVerified
			sps.Db.Save(&model)
			return model.Status
		}
	}
	return domain.Status
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
