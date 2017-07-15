package main // sctl-cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// GetProject Gets the active project metadata from API-server
func (env Env) GetProject() (sctl.Project, error) {
	var project sctl.Project
	res, err := http.Get(env.API.ToURL("active-project"))
	defer res.Body.Close()
	if err != nil {
		return project, err
	}
	err = json.NewDecoder(res.Body).Decode(&project)
	if err != nil {
		return project, err
	}
	return project, nil
}

// GetFolder Gets the configuration folder of the active project from API-server
func (env Env) GetFolder() (string, error) {
	project, err := env.GetProject()
	if err != nil {
		return "", err
	}
	return project.Folder, nil
}

// GetMaster Gets the master node from the api
func GetMaster(API util.ServerConfig) sctl.Node {
	var master sctl.Node
	res, err := http.Get(API.ToURL("master"))
	defer res.Body.Close()
	util.CheckErrFatal(err)
	err = util.DecodeJSON(res.Body, &master)
	util.CheckErrFatal(err)
	return master
}

// SendCommandToNodes Sends a command to be executed on all nodes
func (env Env) SendCommandToNodes(route string, command sctl.Command) string {
	cmd, err := json.Marshal(command)
	util.CheckErrFatal(err)
	res, err := http.Post(
		env.API.ToURL(route), "application/json", bytes.NewBuffer(cmd))
	return handlePostResponse(res, err)
}

// SendToMaster Sends a json body to the cluster master node
func (env Env) SendToMaster(route string, jsonBody []byte) string {
	res, err := http.Post(
		env.API.ToURL(route), "application/json", bytes.NewBuffer(jsonBody))
	return handlePostResponse(res, err)
}

// SendToAPI Sends supplied data to the API-server on a specified route
func (env Env) SendToAPI(route string, data interface{}) string {
	jsonBody, err := json.Marshal(data)
	util.CheckErrFatal(err)
	res, err := http.Post(
		env.API.ToURL(route), "application/json", bytes.NewBuffer(jsonBody))
	return handlePostResponse(res, err)
}

// GetTokens Returns the token bundle of the API server
func (env Env) GetTokens() sctl.TokenBundle {
	res, err := http.Get(env.API.ToURL("tokens"))
	defer res.Body.Close()
	util.CheckErrFatal(err)
	var tokens sctl.TokenBundle
	err = util.DecodeJSON(res.Body, &tokens)
	util.CheckErrFatal(err)
	return tokens
}

// GetFromAPI Performs a get request against the master and returns the output
func GetFromAPI(API util.ServerConfig, route string) string {
	res, err := http.Get(API.ToURL(route))
	return handlePostResponse(res, err)
}

func handlePostResponse(res *http.Response, err error) string {
	if res == nil {
		return "No response, is the API server running?"
	}
	defer res.Body.Close()
	util.CheckErrFatal(err)
	if res.StatusCode == http.StatusOK {
		responseMessage, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "Command successfully sent"
		}
		return string(responseMessage)
	}
	return "Request failed " + res.Status
}
