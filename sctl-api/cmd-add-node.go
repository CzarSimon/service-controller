package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CzarSimon/util"
)

//Node contains node metadata
type Node struct {
	Project  string `json:"project"`
	IP       string `json:"ip"`
	OS       string `json:"os"`
	IsMaster bool   `json:"isMaster"`
}

//SetToMinion sets a node struct to hold values of a minion
func (node *Node) SetToMinion(db *sql.DB) error {
	if node.Project == "" {
		projectName, err := getActiveProject(db)
		if err != nil {
			return err
		}
		node.Project = projectName
	}
	node.IsMaster = false
	node.OS = "linux"
	return nil
}

//AddNode registers a node and installs an sctl-minion on it
func (env *Env) AddNode(res http.ResponseWriter, req *http.Request) {
	var node Node
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
	err = registerNode(node, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

func registerNode(node Node, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO NODE(PROJECT, IP, OS, IS_MASTER) VALUES ($1, $2, $3, $4)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(node.Project, node.IP, node.OS, node.IsMaster)
	if err != nil {
		return err
	}
	return setupNode(node)
}

func setupNode(node Node) error {
	fmt.Println("Setting up node: " + node.IP + " for project " + node.Project)
	return nil
}
