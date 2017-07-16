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
	nodes, err := env.GetNodes()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	env.SendUpdateToNodes(command, nodes)
	util.SendOK(res)
}

// SendUpdateToNodes Sends an update command to nodes
func (env *Env) SendUpdateToNodes(command sctl.Command, nodes []util.ServerConfig) {
	for _, node := range nodes {
		go env.SendToMinion(node, "update", &command)
	}
}
