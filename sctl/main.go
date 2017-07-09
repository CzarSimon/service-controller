package main // sctl-cli

import (
	"os"

	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// Env Holds environment values
type Env struct {
	API util.ServerConfig
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	return Env{
		API: config.API,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	app := cli.NewApp()
	app.Commands = []cli.Command{
		env.UpdateCommand(),
	}
	app.Run(os.Args)
}
