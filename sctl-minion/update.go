package main // sctl-minion

import (
	"log"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//UpdateImage updates the version of the supplied docker image on the minion host
func (env *Env) UpdateImage(res http.ResponseWriter, req *http.Request) {
	env.ValidToken(req)
	var command sctl.Command
	err := util.DecodeJSON(req.Body, &command)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	output, err := command.Execute()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	log.Println(output)
	util.SendOK(res)
}
