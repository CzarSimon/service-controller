package main // sclt-api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// SendToMinion Sends data in json format to specified minion
func (env Env) SendToMinion(minion util.ServerConfig, route string, jsonBody []byte) {
	req, err := http.NewRequest("POST", minion.ToURL(route), bytes.NewBuffer(jsonBody))
	if err != nil {
		util.LogErr(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", env.token)
	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		util.LogErr(err)
		return
	}
	if res.StatusCode == http.StatusOK {
		return
	}
	log.Println("Non 200 response")
}

// CommandToNodes Redirect given command to nodes
func (env Env) CommandToNodes(res http.ResponseWriter, req *http.Request) {
	var command sctl.Command
	err := util.DecodeJSON(req.Body, &command)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	fmt.Println(command)
	util.SendOK(res)
}

// GetNodes Retrives node configurations for the active project
func (env Env) GetNodes() ([]util.ServerConfig, error) {
	nodes := make([]util.ServerConfig, 0)
	query := "SELECT IP FROM NODE WHERE PROJECT=(SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1)"
	rows, err := env.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nodes, err
	}
	node := env.config.minion
	for rows.Next() {
		err = rows.Scan(&node.Host)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// GetMaster Retrives node configurations for the master node of the active project
func (env Env) GetMaster() (util.ServerConfig, error) {
	master := env.config.minion
	query := "SELECT IP FROM NODE WHERE IS_MASTER=1 AND PROJECT=(SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1)"
	err := env.db.QueryRow(query).Scan(&master.Host)
	if err != nil {
		return master, err
	}
	return master, nil
}
