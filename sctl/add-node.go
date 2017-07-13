package main // sctl-cli

import (
	"fmt"
	"path/filepath"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/kardianos/osext"
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
	node := sctl.Node{
		IP:       GetInput("Node IP"),
		IsMaster: false,
		Project:  project.Name,
		OS:       GetInputWithDefault("Node OS", "linux"),
		User:     GetInput("User on master node"),
	}
	fmt.Println(env.SendToAPI("add-node", &node))
	return nil
}

// GetExecPath returns the path of the minion executales folder
func GetExecPath() string {
	execPath, err := osext.ExecutableFolder()
	util.CheckErrFatal(err)
	return filepath.Join(execPath, "executables", "sctl-minion")
}
