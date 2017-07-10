package main // sctl-api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//AddNode registers a node and installs an sctl-minion on it
func (env *Env) AddNode(res http.ResponseWriter, req *http.Request) {
	var node sctl.Node
	err := json.NewDecoder(req.Body).Decode(&node)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = node.SetToMinion(env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	err = RegisterNode(node, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// RegisterNode Stores a given node in the database
func RegisterNode(node sctl.Node, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO NODE(PROJECT, IP, OS, IS_MASTER) VALUES ($1, $2, $3, $4)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(node.Project, node.IP, node.OS, node.IsMaster)
	if err != nil {
		return err
	}
	return nil
}

// SetupNode performs installation of necessary components on a given node
func SetupNode(node sctl.Node) error {
	fmt.Println("Setting up node: " + node.IP + " for project " + node.Project)
	return nil
}
