package main // sctl-minion

import (
	"fmt"
	"net/http"
	"os"

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
	fmt.Println(envVar)
	envVar.Set()
	fmt.Println(os.Getenv(envVar.Key))
	util.SendOK(res)
}
