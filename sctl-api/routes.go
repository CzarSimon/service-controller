package main // sctl-api

import (
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Setups a ServeMux with routes an handler functions
func (env *Env) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/init", env.InitProject)
	mux.HandleFunc("/init-master", env.InitProjectMaster)
	mux.HandleFunc("/add-node", env.AddNode)
	mux.HandleFunc("/init-minion", env.InitMinion)
	mux.HandleFunc("/update", env.UpdateImage)
	mux.HandleFunc("/start", env.CommandToMaster)
	mux.HandleFunc("/stop", env.CommandToMaster)
	mux.HandleFunc("/check", env.CommandToMaster)
	mux.HandleFunc("/alter", env.CommandToMaster)
	mux.HandleFunc("/set-env", env.ForwardEnvVar)
	mux.HandleFunc("/active-project", env.ActiveProject)
	mux.HandleFunc("/project-list", env.GetProjectList)
	mux.HandleFunc("/tokens", env.GetTokens)
	mux.HandleFunc("/master", env.GetMasterNode)
	mux.HandleFunc("/ping", util.Ping)
	mux.HandleFunc("/unlock", env.Unlock)
	mux.HandleFunc("/lock", env.Lock)
	return mux
}

// GetTokens Returns the API-server token bundle to the cli
func (env Env) GetTokens(res http.ResponseWriter, req *http.Request) {
	var masterToken string
	query := "SELECT MASTER_TOKEN FROM PROJECT WHERE IS_ACTIVE=1"
	err := env.db.QueryRow(query).Scan(&masterToken)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonRes, err := json.Marshal(env.token.ToBundle(masterToken))
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonRes)
}
