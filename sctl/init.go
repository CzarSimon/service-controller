package main // sctl-cli

import (
	"fmt"

	"github.com/CzarSimon/sctl-common"
	"github.com/urfave/cli"
)

// InitCommand Cli command for initalizing a project
func (env Env) InitCommand() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initalizes a project",
		Action:  env.InitProject,
	}
}

// InitProject Initalizes a project and sets it to active
func (env Env) InitProject(c *cli.Context) error {
	name := GetInput("Project name")
	folder := GetInput("Definition folder")
	new := sctl.Init{
		Project: sctl.NewProject(name, folder),
		Master: sctl.Node{
			IP:       GetInput("Master IP"),
			IsMaster: true,
			Project:  name,
			OS:       GetInputWithDefault("Node OS", "linux"),
			User:     GetInput("User on master node"),
		},
	}
	env.SendToAPI("init", &new)
	env.SetupNode(new.Master)
	fmt.Println(env.SendToAPI("init-master", &new.Project))
	return nil
}

// GetInput Gets user input from stdin
func GetInput(varName string) string {
	var value string
	fmt.Print(varName + ": ")
	fmt.Scanf("%s", &value)
	return value
}

// GetInputWithDefault Gets user input, if empty returns a supplied default value
func GetInputWithDefault(varName, defaultVal string) string {
	var value string
	fmt.Printf("%s (%s): ", varName, defaultVal)
	fmt.Scanf("%s", &value)
	if value == "" {
		return defaultVal
	}
	return value
}
