package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//InitProject initalizes a project
func (env *Env) InitProject(res http.ResponseWriter, req *http.Request) {
	var project sctl.Project
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

func addProject(project sctl.Project, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO PROJECT(NAME, FOLDER, IS_ACTIVE) VALUES ($1,$2,$3)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(project.Name, project.Folder, true)
	if err != nil {
		return err
	}
	return RegisterNode(project.MakeMasterNode(), db)
}
