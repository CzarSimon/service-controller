package main // sctl-api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

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

// SwitchProject Switches the active project
func (env Env) SwitchProject(res http.ResponseWriter, req *http.Request) {
	var project sctl.Project
	err := util.DecodeJSON(req.Body, &project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	activateStmt, err := env.db.Prepare("UPDATE PROJECT SET IS_ACTIVE=1 WHERE NAME=$1")
	defer activateStmt.Close()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	_, err = activateStmt.Exec(project.Name)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	deactivateStmt, err := env.db.Prepare("UPDATE PROJECT SET IS_ACTIVE=0 WHERE NAME!=$1")
	defer deactivateStmt.Close()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	_, err = deactivateStmt.Exec(project.Name)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// ActiveProject Retrives or switches the active project
func (env Env) ActiveProject(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		env.GetActiveProject(res, req)
		break
	case http.MethodPost:
		env.SwitchProject(res, req)
		break
	default:
		util.SendErrRes(res, errors.New("Unsupported method"))
	}
}

// GetProjectList Returns a list of all projects
func (env Env) GetProjectList(res http.ResponseWriter, req *http.Request) {
	projects := make([]sctl.Project, 0)
	rows, err := env.db.Query("SELECT NAME, IS_ACTIVE FROM PROJECT")
	defer rows.Close()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	var project sctl.Project
	for rows.Next() {
		err := rows.Scan(&project.Name, &project.IsActive)
		if err != nil {
			util.SendErrRes(res, err)
			return
		}
		projects = append(projects, project)
	}
	projectList := formatProjectList(projects)
	util.SendPlainTextRes(res, projectList)
}

// formatProjectList Formats the list of projects to be displayed to the user
func formatProjectList(projects []sctl.Project) string {
	projectList := make([]string, 0)
	var projectRow string
	for _, project := range projects {
		if project.IsActive {
			projectRow = fmt.Sprintf("%s  -  active", project.Name)
		} else {
			projectRow = project.Name
		}
		projectList = append(projectList, projectRow)
	}
	return strings.Join(projectList, "\n")
}
