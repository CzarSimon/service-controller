package main // sctl-cli

import (
	"fmt"
	"strings"

	sctl "github.com/CzarSimon/sctl-common"
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
		env.ListProjects()
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
	if strings.Contains(status, "success") {
		fmt.Println("Set", targetProject.Name, "to active")
	} else {
		fmt.Println(status)
	}
}

// ListProjects Lists projects
func (env Env) ListProjects() {
	projects := env.GetFromAPI("project-list")
	fmt.Println(projects)
}

// IsListOption Determines if first argument is the list option
func IsListOption(c *cli.Context) bool {
	const ListOption string = "ls"
	return c.Args().First() == ListOption
}
