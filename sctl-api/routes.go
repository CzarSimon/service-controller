package main // sctl-api

import (
	"net/http"

	"github.com/CzarSimon/util"
)

func (env *Env) placeholderHandler(res http.ResponseWriter, req *http.Request) {
	util.SendJSONStringRes(res, req.URL.String())
}
