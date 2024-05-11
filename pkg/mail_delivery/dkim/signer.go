package easydkim

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/emersion/go-msgauth/dkim"
	"os"
)

//original https://github.com/metaer/go-easy-dkim-signer

func Sign(data []byte, dkimPrivateKeyFilePath string, selector string, domain string) ([]byte, error) {
	msg := bytes.NewReader(data)
	privateKeyBytes, err := os.ReadFile(dkimPrivateKeyFilePath)
	if err != nil {
		fmt.Println("log file not found")
		return nil, err
	}

	//fmt.Println("Private key bytes", string(privateKeyBytes))
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	result, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey := result

	options := &dkim.SignOptions{
		Domain:   domain,
		Selector: selector,
		Signer:   privateKey,
	}

	var b bytes.Buffer
	if err := dkim.Sign(&b, msg, options); err != nil {
		return nil, fmt.Errorf("dkim signing error: %v", err)
	}

	return b.Bytes(), nil
}
