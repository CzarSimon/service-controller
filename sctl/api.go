package main // sctl-cli

import (
	"bytes"
	"encoding/json"
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

// SendCommandToNodes Sends a command to be executed on all nodes
func (env Env) SendCommandToNodes(route string, command sctl.Command) string {
	cmd, err := json.Marshal(command)
	util.CheckErrFatal(err)
	res, err := http.Post(
		env.API.ToURL(route), "application/json", bytes.NewBuffer(cmd))
	defer res.Body.Close()
	util.CheckErrFatal(err)
	if res.StatusCode == http.StatusOK {
		return "Successfully sent command"
	}
	return "Failed to send command"
}

// SendToMaster Sends a json body to the cluster master node
func (env Env) SendToMaster(route string, jsonBody []byte) string {
	res, err := http.Post(
		env.API.ToURL(route), "application/json", bytes.NewBuffer(jsonBody))
	defer res.Body.Close()
	util.CheckErrFatal(err)
	if res.StatusCode == http.StatusOK {
		return "Successfully sent command"
	}
	return "Failed to send command"
}
