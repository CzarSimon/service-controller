package main // sclt-cli

import (
	"github.com/CzarSimon/sctl-common"
	"github.com/urfave/cli"
)

// UpdateCommand Cli command to update a service
func (env Env) UpdateCommand() cli.Command {
	return cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "Push image to registry and update image on nodes",
		Action:  env.UpdateImage,
	}
}

// UpdateImage pushes and image to remote registyr and updatas image on cluster nodes
func (env Env) UpdateImage(c *cli.Context) error {
	service := GetService(c)
	return nil
}

// GetService Creates a service struct containging a name based on cli arguments
func GetService(c *cli.Context) sctl.Service {
	return sctl.Service{
		Name: c.Args().First(),
	}
}
