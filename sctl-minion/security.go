package main // sctl-minion

import (
	"net/http"
	"os"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// ValidToken Checks if request token equals the minion token
func (env Env) ValidToken(req *http.Request) bool {
	reqToken := req.Header.Get("Authorization")
	return env.token.Data == reqToken
}

// CertificateCommand Creates the command for creation of an rsa certificate and key
func (ssl SSLConfig) CertificateCommand() sctl.Command {
	subject := "/C=SE/ST=Stockholm/L=Stockholm/O=sctl/OU/=security/CN=sctl.minion"
	args := []string{
		"req", "-x509", "-newkey", "rsa:4096", "-keyout", ssl.Key,
		"-out", ssl.Cert, "-days", "100", "-nodes", "-subj", subject}
	return sctl.Command{
		Main: "openssl",
		Args: args,
	}
}

// CertGen Generates self-signed ssl certificates
func (ssl SSLConfig) CertGen() {
	err := os.MkdirAll(ssl.Folder, os.ModePerm)
	util.CheckErrFatal(err)
	_, err = ssl.CertificateCommand().Execute()
	util.CheckErrFatal(err)
}
