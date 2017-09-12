package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// GetActiveNodes Gets the nodes of the active project and send to the requestor
func (env *Env) GetActiveNodes(res http.ResponseWriter, req *http.Request) {
	nodes, err := GetNodes(env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(nodes)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// GetNodes Retrives the nodes of the active active project
func GetNodes(db *sql.DB) ([]sctl.Node, error) {
	nodes := make([]sctl.Node, 0)
	query := `SELECT USER, IP FROM NODE
							WHERE PROJECT=(SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1)`
	rows, err := db.Query(query)
	if err != nil {
		return nodes, err
	}
	var node sctl.Node
	for rows.Next() {
		err = rows.Scan(&node.User, &node.IP)
		if err != nil {
			return nodes, nil
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
