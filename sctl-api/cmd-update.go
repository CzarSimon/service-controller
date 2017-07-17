package main

import (
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//UpdateImage updates a specified image on all nodes in a project cluster
func (env *Env) UpdateImage(res http.ResponseWriter, req *http.Request) {
	var command sctl.Command
	err := util.DecodeJSON(req.Body, &command)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = env.SendToAllNodes("update", &command)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}
