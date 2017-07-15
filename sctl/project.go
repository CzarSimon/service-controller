package main // sctl-cli

import (
	"fmt"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// ProjectCommand Lists all or swithes the active project
func (env Env) ProjectCommand() cli.Command {
	return cli.Command{
		Name:    "project",
		Aliases: []string{"p"},
		Usage:   "Switch or list project (PROJECT_NAME | ls)",
		Action:  env.HandleProjects,
	}
}

// HandleProjects Lists all or switches the active project
func (env Env) HandleProjects(c *cli.Context) error {
	if IsListOption(c) {
		ListProjects(env.API)
	} else {
		targetProject := sctl.Project{
			Name: c.Args().First(),
		}
		env.SwitchProject(targetProject)
	}
	return nil
}

// SwitchProject Switches the active project to the supplied one
func (env Env) SwitchProject(targetProject sctl.Project) {
	status := env.SendToAPI("active-project", &targetProject)
	if status == util.StatusOK {
		fmt.Println("Set", targetProject.Name, "to active")
	} else {
		fmt.Println(status)
	}
}

// ListProjects Lists projects
func ListProjects(API util.ServerConfig) {
	projects := GetFromAPI(API, "project-list")
	fmt.Println(projects)
}

// IsListOption Determines if first argument is the list option
func IsListOption(c *cli.Context) bool {
	const ListOption string = "ls"
	return c.Args().First() == ListOption
}
