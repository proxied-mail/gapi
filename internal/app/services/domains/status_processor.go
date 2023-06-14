package domains

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

		if domain.Status == models.DomainStatusNew {
			sps.checkOwnership(domain)
		}
		if domain.Status == models.DomainStatusOwnershipVerified {
			sps.checkMx(domain)
		}
		if domain.Status == models.DomainStatusMxSet {
			sps.checkSpf(domain)
		}

	}

	return cd
}

func (sps StatusProcessorService) assignVerificationHash(domain *domains.DomainResponse) {
	domain.VerificationHash = GetMD5Hash("hash" + strconv.Itoa(domain.UserId) + "domain-verification")
}

func (sps StatusProcessorService) checkMx(domain *domains.DomainResponse) {
	mxrc, _ := net.LookupMX(domain.Domain)
	for _, mx := range mxrc {
		if mx.Host == "mx.proxiedmail.com." {

			model := domain.GetModel()
			model.Status = models.DomainStatusMxSet
			sps.Db.Save(&model)
		}
	}
}

func (sps StatusProcessorService) checkSpf(domain *domains.DomainResponse) {
	txts, _ := net.LookupTXT(domain.Domain)
	for _, txt := range txts {

		fmt.Println(txt)
		if txt == "v=spf1 include:proxiedmail.com ~all" {

			model := domain.GetModel()
			model.Status = models.DomainStatusSpfSet
			sps.Db.Save(&model)
		}
	}
}

func (sps StatusProcessorService) checkOwnership(domain *domains.DomainResponse) {
	txts, _ := net.LookupTXT(domain.Domain)
	txtStartWith := "proxiedmail-verification="

	for _, txt := range txts {

		fmt.Println(txt)
		if strings.Contains(txt, txtStartWith) {
			splits := strings.Split(txt, txtStartWith)
			code := splits[1]
			if domain.VerificationHash != code {
				continue
			}

			fmt.Println("verified")

			model := domain.GetModel()
			model.Status = models.DomainStatusOwnershipVerified
			sps.Db.Save(&model)

		}
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
