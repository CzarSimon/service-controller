package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/util"
)

//Project contains project metadata
type Project struct {
	Name   string `json:"name"`
	Folder string `json:"folder"`
	Master string `json:"master"`
}

//MakeMasterNode creates node struct for a project master node
func (project Project) MakeMasterNode() Node {
	return Node{
		Project:  project.Name,
		IP:       project.Master,
		OS:       "linux",
		IsMaster: true,
	}
}

//InitProject initalizes a project
func (env *Env) InitProject(res http.ResponseWriter, req *http.Request) {
	var project Project
	err := json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = addProject(project, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

func addProject(project Project, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO PROJECT(NAME, FOLDER, IS_ACTIVE) VALUES ($1,$2,$3)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(project.Name, project.Folder, true)
	if err != nil {
		return err
	}
	return registerNode(project.MakeMasterNode(), db)
}
