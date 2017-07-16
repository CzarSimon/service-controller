package main // sclt-api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// SendToMinion Sends data in json format to specified minion
func (env *Env) SendToMinion(minion util.ServerConfig, route string, data interface{}) {
	res, err := env.postToNode(minion, route, data)
	defer res.Body.Close()
	if err != nil {
		util.LogErr(err)
		return
	}
	if res.StatusCode == http.StatusOK {
		return
	}
	log.Printf("%d response\n", res.StatusCode)
}

// SendToAllNodes Sends data to all nodes in the cluster
func (env *Env) SendToAllNodes(route string, data interface{}) error {
	nodes, err := env.GetNodes()
	if err != nil {
		return err
	}
	for _, node := range nodes {
		go env.SendToMinion(node, route, data)
	}
	return nil
}

// GetResFromMinion Sends data in json format to specified minion and returns the response
func (env *Env) GetResFromMinion(minion util.ServerConfig, route string, data interface{}) (*http.Response, error) {
	res, err := env.postToNode(minion, route, data)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		return res, errors.New(res.Status)
	}
	return res, nil
}

// postToNode Issues a post request to a specified minon and returns the result and error
func (env *Env) postToNode(minion util.ServerConfig, route string, data interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return &http.Response{}, err
	}
	req, err := http.NewRequest(http.MethodPost, minion.ToURL(route), bytes.NewBuffer(jsonBody))
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	env.SetReqToken(req)
	client := sslClient()
	return client.Do(req)
}

// SetReqToken Sets the authorization header of the request
func (env *Env) SetReqToken(req *http.Request) {
	env.reqCount++
	req.Header.Set("Authorization", env.token.Data)
	env.reqCount--
}

// sslClient Returns a client that can make request to minion using self self-signed certificate
func sslClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// GetNodes Retrives node configurations for the active project
func (env *Env) GetNodes() ([]util.ServerConfig, error) {
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
func (env *Env) GetMaster() (util.ServerConfig, error) {
	master := env.config.minion
	query := "SELECT IP FROM NODE WHERE IS_MASTER=1 AND PROJECT=(SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1)"
	err := env.db.QueryRow(query).Scan(&master.Host)
	if err != nil {
		return master, err
	}
	return master, nil
}

// CommandToMaster Redirects a command to the master node and returns the response
func (env *Env) CommandToMaster(res http.ResponseWriter, req *http.Request) {
	var command sctl.Command
	err := util.DecodeJSON(req.Body, &command)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	master, err := env.GetMaster()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	masterResponse, err := env.GetResFromMinion(master, forwardRoute(req), command)
	defer masterResponse.Body.Close()
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	responseMessage, err := ioutil.ReadAll(masterResponse.Body)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendPlainTextRes(res, string(responseMessage))
}

// forwardRoute Creates a route that can be passed to ServerConfig.ToURL
func forwardRoute(req *http.Request) string {
	return strings.Replace(req.URL.Path, "/", "", -1)
}
