package main // sctl-api

import (
	"database/sql"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//InitProject initalizes a project
func (env *Env) InitProject(res http.ResponseWriter, req *http.Request) {
	var project sctl.Project
	err := util.DecodeJSON(req.Body, &project)
	//err := json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	masterNode := project.MakeMasterNode()
	project = sctl.NewProject(project.Name, project.Folder)
	err = AddProject(project, masterNode, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// AddProject Registers a given project and the master node and sets up the master node
func AddProject(project sctl.Project, master sctl.Node, db *sql.DB) error {
	err := project.Insert(db)
	if err != nil {
		return err
	}
	return RegisterNode(master, db)
}
