package main // sctl-minion

import (
	"net/http"

	"github.com/CzarSimon/sctl-common"
)

// ValidToken Checks if request token equals the minion token
func (env Env) ValidToken(req *http.Request) bool {
	reqToken := req.Header.Get("Authorization")
	//fmt.Println("minon token:", env.token)
	//fmt.Println("Request token:", reqToken)
	return env.token.Data == reqToken
}

// SetInitalToken Sets a valid token if current one is the intital token
func (env *Env) SetInitalToken(req *http.Request) {
	if env.token.Data == InitalToken {
		env.token.Data = req.Header.Get("Authorization")
	}
}

// SetMasterToken Sets a valid master token in current one is the inital master token and a valid request token was supplied
func (env *Env) SetMasterToken(req *http.Request, project sctl.Project) {
	if env.ValidToken(req) && env.masterToken == InitalToken {
		env.masterToken = project.MasterToken
	}
}
