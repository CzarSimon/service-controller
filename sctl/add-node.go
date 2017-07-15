package main // sctl-cli

import (
	"fmt"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// AddNodeCommand Cli command adding a new node
func (env Env) AddNodeCommand() cli.Command {
	return cli.Command{
		Name:   "add-node",
		Usage:  "Registers a new node and starts a minon on it",
		Action: env.AddNode,
	}
}

// AddNode Registers a new node and starts a minon on it
func (env Env) AddNode(c *cli.Context) error {
	project, err := env.GetProject()
	util.CheckErrFatal(err)
	master := GetMaster(env.API)
	util.CheckErrFatal(err)
	node := GetNewNode(project)
	env.SendToAPI("add-node", &node)
	cmd := sctl.MinionCommand{
		Minion:  node,
		Command: JoinSwarmCommand(project, master),
	}
	env.SetupNode(node)
	fmt.Println(env.SendToAPI("init-minion", cmd))
	return nil
}

// GetNewNode Retrives configuraton information about the new node
func GetNewNode(project sctl.Project) sctl.Node {
	return sctl.Node{
		IP:       GetInput("Node IP"),
		IsMaster: false,
		Project:  project.Name,
		OS:       GetInputWithDefault("Node OS", "linux"),
		User:     GetInput("User on master node"),
	}
}

// JoinSwarmCommand Creates a command to join a swarm
func JoinSwarmCommand(project sctl.Project, master sctl.Node) sctl.Command {
	args := []string{"swarm", "join", "--token", project.SwarmToken, master.IP + ":2377"}
	return sctl.DockerCommand(args)
}
