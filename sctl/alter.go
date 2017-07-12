package main // sctl-cli

import (
	"encoding/json"
	"fmt"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// AlterCommand Alters the cluster or specific service
func (env Env) AlterCommand() cli.Command {
	return cli.Command{
		Name:    "alter",
		Aliases: []string{"a"},
		Usage:   "Alters the cluster or specific service, use the docker command line api",
		Action:  env.AlterCluster,
	}
}

// AlterCluster Alters the cluster or specific service
func (env Env) AlterCluster(c *cli.Context) error {
	alterCommand := ArgsToDockerCommand(c)
	jsonCMD, err := json.Marshal(alterCommand)
	util.CheckErrFatal(err)
	status := env.SendToMaster("alter", jsonCMD)
	if status == util.StatusOK {
		color.Blue("Successfully executed: " + alterCommand.ToString())
	} else {
		fmt.Println(status)
	}
	return nil
}

// ArgsToDockerCommand Parses cli into a docker command
func ArgsToDockerCommand(c *cli.Context) sctl.Command {
	return sctl.DockerCommand(c.Args())
}
