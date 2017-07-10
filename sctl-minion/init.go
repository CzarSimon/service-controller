package main // sctl-minion

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// SetupMaster Initalizes a docker swarm and sets up an overlay network for the cluster
func (env *Env) SetupMaster(res http.ResponseWriter, req *http.Request) {
	env.SetInitalToken(req)
	var project sctl.Project
	err := util.DecodeJSON(req.Body, &project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	env.SetMasterToken(req, project)
	project.SwarmToken, err = SetupSwarmAndNetwork(project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(project)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// SetupSwarmAndNetwork Sets up swarm cluster and overlay network, returns swarm token
func SetupSwarmAndNetwork(project sctl.Project) (string, error) {
	token, err := CreateSwarm()
	if err != nil {
		return token, err
	}
	err = CreateOverlayNetwork(project.Network)
	if err != nil {
		return token, err
	}
	return token, nil
}

// CreateSwarm Creates docker swarm, returns the swarm worker token
func CreateSwarm() (string, error) {
	swarmCommand := sctl.DockerCommand([]string{"swarm", "init"})
	out, err := swarmCommand.Execute()
	if err != nil {
		return out, err
	}
	fmt.Println(out)
	return GetSwarmToken()
}

// GetSwarmToken Returns the swarm worker token
func GetSwarmToken() (string, error) {
	swarmTokenCMD := sctl.DockerCommand([]string{"swarm", "join-token", "worker", "-q"})
	out, err := swarmTokenCMD.Execute()
	if err != nil {
		return out, err
	}
	token := strings.Replace(out, "\n", "", -1)
	return token, nil
}

// CreateOverlayNetwork Creates an overlay network for the cluster
func CreateOverlayNetwork(networkName string) error {
	networkArgs := []string{"network", "create", "--driver", "overlay", networkName}
	createNetwork := sctl.DockerCommand(networkArgs)
	out, err := createNetwork.Execute()
	if err != nil {
		log.Println(out)
		return err
	}
	return nil
}
