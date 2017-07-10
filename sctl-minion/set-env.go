package main // sctl-minion

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// SetEnvVar Sets supplied environment variable
func SetEnvVar(res http.ResponseWriter, req *http.Request) {
	var envVar util.EnvVar
	err := util.DecodeJSON(req.Body, &envVar)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	envVar.Set()
	util.SendOK(res)
}
