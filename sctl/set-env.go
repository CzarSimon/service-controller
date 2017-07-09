package main // sctl-cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// SetEnvCommand Cli command to update a service
func (env Env) SetEnvCommand() cli.Command {
	return cli.Command{
		Name:   "set-env",
		Usage:  "Sets a supplied environment variable on the cluster master. Use as follows: set-env key=val",
		Action: env.SetEnvVar,
	}
}

// SetEnvVar Parses environment variable argument and sends request to the cluster master node
func (env Env) SetEnvVar(c *cli.Context) error {
	args := strings.Split(c.Args().First(), "=")
	if len(args) != 2 {
		fmt.Println("Incorrect formating of argument", c.Args().First())
		fmt.Println("Use set-env --help")
		return errors.New("Incorrct first argument")
	}
	envVar := util.EnvVar{
		Key:   args[0],
		Value: args[1],
	}
	fmt.Println(envVar)
	jsonBody, err := json.Marshal(envVar)
	util.CheckErrFatal(err)
	status := env.SendToMaster("set-env", jsonBody)
	fmt.Println(status)
	return nil
}
