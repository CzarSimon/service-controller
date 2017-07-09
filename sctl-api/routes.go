package main // sctl-api

import (
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

func (env *Env) placeholderHandler(res http.ResponseWriter, req *http.Request) {
	util.SendJSONStringRes(res, req.URL.String())
}

// GetActiveProject Returns the current active project
func (env Env) GetActiveProject(res http.ResponseWriter, req *http.Request) {
	var project sctl.Project
	query := "SELECT NAME, FOLDER, SWARM_TOKEN, NETWORK FROM PROJECT WHERE IS_ACTIVE=1"
	err := env.db.QueryRow(query).Scan(
		&project.Name, &project.Folder, &project.SwarmToken, &project.Network)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}
