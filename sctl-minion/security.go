package main // sctl-minion

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// ResetToken Updates the minion token
func (env *Env) ResetToken(res http.ResponseWriter, req *http.Request) {
	var newToken sctl.Token
	err := util.DecodeJSON(req.Body, &newToken)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	env.token = newToken
	util.SendOK(res)
}

// LockHandler Handles locking and unlocking
func (env *Env) LockHandler(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/lock":
		env.lock.Close()
		util.SendOK(res)
		break
	case "/unlock":
		if env.Unlock(req) {
			util.SendOK(res)
		} else {
			util.SendUnauthorized(res)
		}
		break
	default:
		break
	}
}

// MinionLock Holds lock state
type MinionLock struct {
	Open               bool
	FailedAuthAttempts int
	MaxTokenAge        float64
	MaxAttempts        int
}

// NewLock Returns a new open lock with 1 auth attempt left
func NewLock(config LockConfig) MinionLock {
	return MinionLock{
		Open:               true,
		FailedAuthAttempts: config.MaxAttempts - 1,
		MaxTokenAge:        config.TokenMaxAge,
		MaxAttempts:        config.MaxAttempts,
	}
}

// Close Closes the MinionLock
func (lock *MinionLock) Close() {
	lock.Open = false
	lock.FailedAuthAttempts = lock.MaxAttempts + 1
}

// RegisterFail registers an authorization failure
func (lock *MinionLock) RegisterFail() {
	lock.FailedAuthAttempts++
	if lock.FailedAuthAttempts >= lock.MaxAttempts {
		lock.Close()
	}
}

// ValidToken Checks if request token equals the minion token
func (env *Env) ValidToken(req *http.Request) bool {
	if !env.token.Valid(env.lock.MaxTokenAge) {
		fmt.Println(env.token.Timestamp, "To Old")
		env.lock.Close()
		return false
	}
	authSuccess := env.token.Data == req.Header.Get("Authorization")
	if authSuccess {
		env.lock.FailedAuthAttempts = 0
	} else {
		env.lock.RegisterFail()
	}
	return authSuccess
}

// Unlock Compares a candidate token a against the master token and unlocks the minion if they match
func (env *Env) Unlock(req *http.Request) bool {
	var candidate sctl.TokenBundle
	err := util.DecodeJSON(req.Body, &candidate)
	success := err == nil && env.masterToken == candidate.Master
	if success {
		env.token = candidate.Auth
		env.lock.Open = true
		env.lock.FailedAuthAttempts = env.config.Lock.MaxAttempts - 1
	}
	return success
}

// CertificateCommand Creates the command for creation of an rsa certificate and key
func (ssl SSLConfig) CertificateCommand() sctl.Command {
	subject := "/C=SE/ST=Stockholm/L=Stockholm/O=sctl/OU=security/CN=sctl-minion"
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

// Auth Checks if a request is made with a valid token
func (env *Env) Auth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		if !env.lock.Open {
			logAuthStatus("request made to a locked node", req)
			util.SendErrRes(res, errors.New("node locked"))
		} else if env.ValidToken(req) {
			logAuthStatus("auth success", req)
			handler(res, req)
		} else {
			logAuthStatus("auth failue", req)
			util.SendUnauthorized(res)
		}
	}
}

// logAuthStatus Logs outcome of authorization challange
func logAuthStatus(msg string, req *http.Request) {
	log.Printf("%s from: %s, %s", req.URL.Path, req.RemoteAddr, msg)
}
