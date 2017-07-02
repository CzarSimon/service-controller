package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CzarSimon/util"
)

//Service contains info about the service whose image will be updated
type Service struct {
	Image string
}

//UpdateImage updates a specified image on all nodes in a project cluster
func (env *Env) UpdateImage(res http.ResponseWriter, req *http.Request) {
	var service Service
	err := json.NewDecoder(req.Body).Decode(&service)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	nodes, err := getNodes(env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	sendUpdateToNodes(service, nodes, env.config.minion)
	util.SendOK(res)
}

func sendUpdateToNodes(service Service, nodes []string, minion util.ServerConfig) {
	for _, node := range nodes {
		minion.Host = node
		jsonBody, err := json.Marshal(service)
		if err != nil {
			util.LogErr(err)
			break
		}
		go sendToMinion(minion, "update", jsonBody)
	}
}

func getNodes(db *sql.DB) ([]string, error) {
	nodes := make([]string, 0)
	query := "SELECT IP FROM NODE WHERE PROJECT=(SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1)"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nodes, nil
	}
	var node string
	for rows.Next() {
		err = rows.Scan(&node)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
