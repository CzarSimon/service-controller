package main // sctl-cli

import (
	"encoding/json"
	"fmt"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// CheckCommand Checks the status of the cluster or a specific service
func (env Env) CheckCommand() cli.Command {
	return cli.Command{
		Name:    "check",
		Aliases: []string{"c"},
		Usage:   "Checks the status of the cluster or a specific service, use the docker command line api",
		Action:  env.CheckCluster,
	}
}

// CheckCluster Checks the status of the cluster or a specific service
func (env Env) CheckCluster(c *cli.Context) error {
	checkCommand := ArgsToDockerCommand(c)
	jsonCMD, err := json.Marshal(checkCommand)
	util.CheckErrFatal(err)
	status := env.SendToMaster("check", jsonCMD)
	fmt.Println(status)
	return nil
}
