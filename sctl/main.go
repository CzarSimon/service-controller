package main // sctl-cli

import (
	"os"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// Env Holds environment values
type Env struct {
	API    util.ServerConfig
	config Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	return Env{
		API:    config.API,
		config: config,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	app := cli.NewApp()
	app.Name = config.App.Name
	app.Usage = config.App.Usage
	app.Version = config.App.Version
	app.Commands = []cli.Command{
		env.UpdateCommand(),
		env.SetEnvCommand(),
		env.InitCommand(),
		env.ProjectCommand(),
		env.CheckCommand(),
		env.AlterCommand(),
		env.StartCommand(),
		env.AddNodeCommand(),
		env.StopCommand(),
		env.LockCommand(),
		env.UnlockCommand(),
	}
	app.Run(os.Args)
}
