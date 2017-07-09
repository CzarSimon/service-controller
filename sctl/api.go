package main // sctl-cli

import (
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/sctl-common"
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
