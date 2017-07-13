package main // sctl-cli

import (
	"fmt"
	"path/filepath"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
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
	folder := GetInputWithDefault("Definition folder", "/Users/simon/workspace/mimir/service-definitions")
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
	fmt.Println(env.SendToAPI("init", &new))
	env.SetupMinonDB()
	env.SetupNode(new.Master)
	WaitForStartUp(new.Master)
	return nil
}

// WaitForStartUp Waits unit user has started minon on node
func WaitForStartUp(node sctl.Node) {
	fmt.Println("Sent executables to node")
	GetInput("Start minion on " + node.IP + " then press enter")
}

// SetupNode Sends executbles and starts them on the node
func (env Env) SetupNode(node sctl.Node) {
	SendExecutables(env.config.Folders, node)
	SendInitd(env.config.Folders, node)
}

// SendExecutables Sends executables to designated destination on node
func SendExecutables(folders FolderConfig, node sctl.Node) {
	send := node.RsyncFolderCMD(folders.Exec, folders.Target)
	out, err := send.Execute()
	util.CheckErrFatal(err)
	if out != "" {
		fmt.Println(out)
	}
}

// SendInitd Sends initialization defiontion to the node
func SendInitd(folders FolderConfig, node sctl.Node) {
	initdFile := filepath.Join(folders.Exec, "sctl-minion.service")
	send := node.RsyncFileCMD(initdFile, folders.Target)
	out, err := send.Execute()
	util.CheckErrFatal(err)
	if out != "" {
		fmt.Println(out)
	}
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
