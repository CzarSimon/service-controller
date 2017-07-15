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
	mux.HandleFunc("/start", env.Auth(RunCommand))
	mux.HandleFunc("/stop", env.Auth(RunCommand))
	mux.HandleFunc("/update", env.Auth(RunCommand))
	mux.HandleFunc("/alter", env.Auth(RunCommand))
	mux.HandleFunc("/check", env.Auth(RunCommand))
	mux.HandleFunc("/init", env.Auth(env.SetupMaster))
	mux.HandleFunc("/join-swarm", env.Auth(RunCommand))
	mux.HandleFunc("/ping", env.Auth(util.Ping))
	mux.HandleFunc("/unlock", env.LockHandler)
	mux.HandleFunc("/lock", env.LockHandler)
	mux.HandleFunc("/reset-token", env.Auth(env.ResetToken))
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
	fmt.Println(cmd.ToString())
	output, err := cmd.Execute()
	if err != nil {
		fmt.Println(output)
		util.SendErrRes(res, err)
		return
	}
	util.SendPlainTextRes(res, output)
}
