package main // sctl-api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//InitProject initalizes a project
func (env *Env) InitProject(res http.ResponseWriter, req *http.Request) {
	var new sctl.Init
	err := util.DecodeJSON(req.Body, &new)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = env.AddProject(new)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// InitProjectMaster Sets up the master node
func (env *Env) InitProjectMaster(res http.ResponseWriter, req *http.Request) {
	var project sctl.Project
	err := util.DecodeJSON(req.Body, &project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = env.SetupMaster(project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// AddProject Registers a given project and the master node
func (env *Env) AddProject(new sctl.Init) error {
	err := new.Project.Insert(env.db)
	if err != nil {
		return err
	}
	err = RegisterNode(new.Master, env.db)
	if err != nil {
		return err
	}
	return nil
}

// SetupMaster Initalizes a minion as master on a specified node
func (env *Env) SetupMaster(project sctl.Project) error {
	master, err := env.GetMaster()
	if err != nil {
		return err
	}
	env.SetupSwarmAndRetriveToken(master, "init", &project)
	InsertSwarmToken(project, env.db)
	return nil
}

// SetupSwarmAndRetriveToken Sets up swarm and returns the swarm token
func (env *Env) SetupSwarmAndRetriveToken(
	master util.ServerConfig, route string, project *sctl.Project) error {
	res, err := env.GetResFromMinion(master, route, project)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	err = util.DecodeJSON(res.Body, project)
	if err != nil {
		return err
	}
	return nil
}

// InsertSwarmToken Stores a retrived swarm token in the database
func InsertSwarmToken(project sctl.Project, db *sql.DB) error {
	if project.SwarmToken == "" {
		return errors.New("Empty swarm token")
	}
	stmt, err := db.Prepare("UPDATE PROJECT SET SWARM_TOKEN=$1 WHERE NAME=$2")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(project.SwarmToken, project.Name)
	if err != nil {
		return err
	}
	return nil
}
