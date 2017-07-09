package main // sctl-api

import (
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/util"
)

// ForwardEnvVar Forwards a request to set an environment variable to cluster master
func (env *Env) ForwardEnvVar(res http.ResponseWriter, req *http.Request) {
	var envVar util.EnvVar
	err := util.DecodeJSON(req.Body, &envVar)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	master, err := env.GetMaster()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(envVar)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	env.SendToMinion(master, "set-env", jsonBody)
	util.SendOK(res)
}
