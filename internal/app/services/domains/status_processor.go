package domains

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/abrouter/gapi/internal/app/http/response/domains"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/pkg/mxapi"
	"github.com/miekg/dns"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net"
	"strconv"
	"strings"
	"time"
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
			domain.Status, _ = sps.checkOwnership(domain)
		}

		if domain.Status == models.DomainStatusOwnershipVerified {
			domain.Status, _ = sps.checkMx(domain)
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
	domain.VerificationHash = GetMD5Hash("hash2" + strconv.Itoa(domain.UserId) + "domain-verification")
}

func (sps StatusProcessorService) assignSpf(domain *domains.DomainResponse) {
	domain.Spf = "v=spf1 include:spf.proxiedmail.com ~all"
}

func (sps StatusProcessorService) checkMx(domain *domains.DomainResponse) (int, error) {
	mxrc, _ := sps.getResolver().LookupMX(context.Background(), domain.Domain)

	for _, mx := range mxrc {
		if mx.Host == "mx.proxiedmail.com." {

			model := domain.GetModel()
			if model.DkimKey == "" {
				mxapiReponseEntity, err := mxapi.CreateNewUserCatchAllRequest(model.Domain, model.SmtpPassword.String)
				if err != nil || !mxapiReponseEntity.IsCreated {
					return 0, errors.New("Error creating domain on MX")
				}
				dkim, err2 := mxapi.RequestDkim(model.Domain)
				if err2 != nil {
					return 0, err2
				}
				model.DkimKey = dkim.Content
				domain.DkimKey = dkim.Content
			}

			model.Status = models.DomainStatusMxSet
			sps.Db.Save(&model)
			return model.Status, nil
		}
	}
	return domain.Status, nil
}

func (sps StatusProcessorService) checkDkim(domain *domains.DomainResponse) int {

	cnameCheckDomain := "dkim._domainkey." + domain.Domain + "."

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)

	// Note the trailing dot. miekg/dns is very low-level and expects canonical names.
	m.SetQuestion(cnameCheckDomain, dns.TypeCNAME)
	m.RecursionDesired = true
	r, _, error := c.Exchange(m, config.Servers[0]+":"+config.Port)
	if error != nil {
		fmt.Println(error.Error())
		return domain.Status
	}

	if len(r.Answer) == 0 {
		return domain.Status
	}

	if r.Answer[0].(*dns.CNAME).Target == "dkim._domainkey.pxdmail.com." {
		model := domain.GetModel()
		model.Status = models.DomainStatusDkimSet
		sps.Db.Save(&model)
		return model.Status
	}

	return domain.Status
}

func (sps StatusProcessorService) checkSpf(domain *domains.DomainResponse) int {
	txts, _ := sps.getResolver().LookupTXT(context.Background(), domain.Domain)

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

func (sps StatusProcessorService) getResolver() *net.Resolver {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	return r
}

func (sps StatusProcessorService) checkOwnership(domain *domains.DomainResponse) (int, error) {
	txts, _ := sps.getResolver().LookupTXT(context.Background(), domain.Domain)
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

			return model.Status, nil
		}
	}

	return domain.Status, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
