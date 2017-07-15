package main // sctl-cli

import (
	"encoding/json"
	"fmt"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// StopCommand Stops a service running on the cluster
func (env Env) StopCommand() cli.Command {
	return cli.Command{
		Name:   "stop",
		Usage:  "Stops a service running on the cluster",
		Action: env.StopService,
	}
}

// StopService Retrives a the specified service defenintion and issues a stop command to the master
func (env Env) StopService(c *cli.Context) error {
	service := env.GetServiceDef(c)
	stopCommand := service.StopCommand()
	fmt.Println(stopCommand.ToString())
	jsonCMD, err := json.Marshal(stopCommand)
	util.CheckErrFatal(err)
	status := env.SendToMaster("stop", jsonCMD)
	fmt.Println(status)
	return nil
}
