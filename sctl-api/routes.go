package main // sctl-api

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Setups a ServeMux with routes an handler functions
func (env *Env) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/init", env.InitProject)
	mux.HandleFunc("/add-node", env.AddNode)
	mux.HandleFunc("/update", env.UpdateImage)
	mux.HandleFunc("/start", util.PlaceholderHandler)
	mux.HandleFunc("/check", env.CommandToMaster)
	mux.HandleFunc("/alter", env.CommandToMaster)
	mux.HandleFunc("/set-env", env.ForwardEnvVar)
	mux.HandleFunc("/active-project", env.ActiveProject)
	mux.HandleFunc("/project-list", env.GetProjectList)
	return mux
}
