package main // sctl-minion

import (
	"fmt"
	"net/http"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// SetupRoutes Setups a ServeMux with routes an handler functions
func (env *Env) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/update", RunCommand)
	mux.HandleFunc("/set-env", SetEnvVar)
	mux.HandleFunc("/init", env.SetupMaster)
	mux.HandleFunc("/reset-token", util.PlaceholderHandler)
	mux.HandleFunc("/alter", RunCommand)
	mux.HandleFunc("/check", RunCommand)
	mux.HandleFunc("/start", RunCommand)
	return mux
}

// RunCommand Executes a command on the minion and returns the output
func RunCommand(res http.ResponseWriter, req *http.Request) {
	var cmd sctl.Command
	err := util.DecodeJSON(req.Body, &cmd)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	output, err := cmd.Execute()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	fmt.Println(cmd.ToString())
	util.SendPlainTextRes(res, output)
}
