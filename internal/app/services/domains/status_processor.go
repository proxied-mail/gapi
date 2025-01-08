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
	"sync"
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
	mxrc, _ := sps.lookupMX2(context.Background(), domain.Domain)

	for _, mx := range mxrc {
		if mx.Mx == "mx.proxiedmail.com." {
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
	txts, _ := sps.getTXTRecords(domain.Domain)

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
	servers := []string{"8.8.8.8:53", "1.1.1.1:53"}
	timeout := 2 * time.Second

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			done := make(chan net.Conn, 1)

			for _, server := range servers {
				go func(server string) {
					d := net.Dialer{Timeout: timeout}
					conn, err := d.DialContext(ctx, network, server)
					if err == nil {
						done <- conn
					}
				}(server)
			}

			select {
			case conn := <-done:
				return conn, nil
			case <-time.After(timeout):
				return nil, fmt.Errorf("timeout resolving DNS")
			}
		},
	}

	return r
}

func (sps StatusProcessorService) getResolver2() *dns.Client {
	return &dns.Client{
		Timeout: time.Second * 10,
	}
}

func (sps StatusProcessorService) lookupMX2(ctx context.Context, domain string) ([]*dns.MX, error) {
	c := sps.getResolver2()

	// Create a DNS message for MX lookup
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeMX)

	// Use a reliable DNS server (e.g., Google's 8.8.8.8)
	dnsServer := "8.8.8.8:53"

	// Perform the query
	resp, _, err := c.ExchangeContext(ctx, msg, dnsServer)
	if err != nil {
		return nil, fmt.Errorf("DNS query failed: %w", err)
	}

	if resp == nil || resp.Rcode != dns.RcodeSuccess {
		return nil, fmt.Errorf("failed to get a valid DNS response")
	}

	// Parse MX records from the response
	var mxRecords []*dns.MX
	for _, answer := range resp.Answer {
		if mx, ok := answer.(*dns.MX); ok {
			mxRecords = append(mxRecords, mx)
		}
	}

	return mxRecords, nil
}

func (sps StatusProcessorService) checkOwnership(domain *domains.DomainResponse) (int, error) {
	txts, _ := sps.getTXTRecords(domain.Domain)
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

func (sps StatusProcessorService) getTXTRecords(domain string) ([]string, error) {
	servers := []string{
		"8.8.8.8:53", "8.8.4.4:53", // Google Public DNS
		"1.1.1.1:53", "1.0.0.1:53", // Cloudflare DNS
		"9.9.9.9:53", "149.112.112.112:53", // Quad9 DNS
		"208.67.222.222:53", "208.67.220.220:53", // OpenDNS
		"185.228.168.9:53", "185.228.169.9:53", // CleanBrowsing
		"94.140.14.14:53", "94.140.15.15:53", // AdGuard DNS
	}

	timeout := 2 * time.Second

	var wg sync.WaitGroup
	results := make(chan []string, len(servers))
	errorsList := make(chan error, len(servers))

	// Query each DNS server concurrently
	for _, server := range servers {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()

			r := &net.Resolver{
				PreferGo: true,
				Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
					d := net.Dialer{Timeout: timeout}
					return d.DialContext(ctx, network, server)
				},
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			txts, err := r.LookupTXT(ctx, domain)
			if err != nil {
				errorsList <- err
				return
			}

			results <- txts
		}(server)
	}

	// Wait for all queries to complete
	go func() {
		wg.Wait()
		close(results)
		close(errorsList)
	}()

	// Collect results
	var allRecords []string
	var lastError error

	for res := range results {
		allRecords = append(allRecords, res...)
	}
	for err := range errorsList {
		lastError = err
	}

	// Deduplicate records (optional)
	uniqueRecords := make(map[string]struct{})
	for _, record := range allRecords {
		uniqueRecords[record] = struct{}{}
	}

	mergedRecords := make([]string, 0, len(uniqueRecords))
	for record := range uniqueRecords {
		mergedRecords = append(mergedRecords, record)
	}

	if len(mergedRecords) == 0 && lastError != nil {
		return nil, lastError
	}

	return mergedRecords, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
