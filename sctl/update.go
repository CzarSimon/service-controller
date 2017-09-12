package main // sclt-cli

import (
	"fmt"
	"sync"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
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
	service := env.GetServiceDef(c)
	fmt.Println("Updating", service.Name, "image...")
	// _, err := service.PushCommand().Execute()
	// util.CheckErrFatal(err)
	fmt.Println(service.Name, "image:", service.Image, "pushed")
	var wait sync.WaitGroup
	for _, node := range GetNodes(env.API) {
		wait.Add(1)
		go UpdateImageOnNode(service.Image, node, &wait)
	}
	wait.Wait()
	return nil
}

// UpdateImageOnNode Updates image on given node using ssh
func UpdateImageOnNode(image string, node sctl.Node, wait *sync.WaitGroup) {
	defer wait.Done()
	if node.IsLocal() {
		return
	}
	updateCMD := node.SSHCommand("docker", "pull", image)
	_, err := updateCMD.Execute()
	if err != nil {
		fmt.Println("Error occured on node: " + node.IP)
		util.PrintErr(err)
	} else {
		fmt.Printf("%s updated on %s\n", image, node.IP)
	}
}

// GetServiceDef Retrives service definitoin
func (env Env) GetServiceDef(c *cli.Context) sctl.Service {
	service := GetService(c)
	folder, err := env.GetFolder()
	util.CheckErrFatal(err)
	err = service.GetServiceDef(folder)
	util.CheckErrFatal(err)
	return service
}

// GetService Creates a service struct containging a name based on cli arguments
func GetService(c *cli.Context) sctl.Service {
	return sctl.Service{
		Name: c.Args().First(),
	}
}
