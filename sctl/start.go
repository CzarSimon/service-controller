package main // sctl-cli

import (
	"encoding/json"
	"fmt"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// StartCommand Command for service start on the cluster
func (env Env) StartCommand() cli.Command {
	return cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Retrives the service specification for the specified service and starts it on the cluster",
		Action:  env.StartService,
	}
}

// StartService Retrives the service specification for the specified service and starts it on the cluster
func (env Env) StartService(c *cli.Context) error {
	service := env.GetServiceDef(c)
	project, err := env.GetProject()
	util.CheckErrFatal(err)
	startCommand := service.RunCommand(project.Network)
	fmt.Println(startCommand.ToString())
	jsonCMD, err := json.Marshal(startCommand)
	util.CheckErrFatal(err)
	status := env.SendToMaster("start", jsonCMD)
	fmt.Println(status)
	return nil
}
