package main // sctl-minion

import (
	"net/http"
	"os"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"golang.org/x/crypto/acme/autocert"
)

// ValidToken Checks if request token equals the minion token
func (env Env) ValidToken(req *http.Request) bool {
	reqToken := req.Header.Get("Authorization")
	return env.token.Data == reqToken
}

// CertificateCommand Creates the command for creation of an rsa certificate and key
func (ssl SSLConfig) CertificateCommand() sctl.Command {
	subject := "/C=SE/ST=Stockholm/L=Stockholm/O=sctl/OU=security/CN=sctl.com"
	args := []string{
		"req", "-x509", "-newkey", "rsa:4096", "-keyout", ssl.Key,
		"-out", ssl.Cert, "-days", "100", "-nodes", "-subj", subject}
	return sctl.Command{
		Main: "openssl",
		Args: args,
	}
}

// GetCertManager Sets up an autocert certificate manager
func GetCertManager(ssl SSLConfig) autocert.Manager {
	return autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("sctl.com"), //your domain here
		Cache:      autocert.DirCache("certs"),
	}
}

// CertGen Generates self-signed ssl certificates
func (ssl SSLConfig) CertGen() {
	err := os.MkdirAll(ssl.Folder, os.ModePerm)
	util.CheckErrFatal(err)
	_, err = ssl.CertificateCommand().Execute()
	util.CheckErrFatal(err)
}
