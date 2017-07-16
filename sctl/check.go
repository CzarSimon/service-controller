package main // sctl-cli

import (
	"encoding/json"
	"strings"

	"github.com/CzarSimon/util"
	"github.com/fatih/color"
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
	statusPrint(status)
	return nil
}

func statusPrint(status string) {
	if !strings.Contains(status, "Request failed") {
		color.Blue(status)
	} else {
		color.Red(status)
	}
}
