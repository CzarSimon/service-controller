package main // sctl-minion

import (
	"log"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// InitalToken Holds the inital token value
const InitalToken string = "INITAL_TOKEN"

//Env is the struct for environment objects
type Env struct {
	masterToken string
	token       sctl.Token
	config      Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	token := sctl.NewToken(1)
	token.Data = InitalToken
	return Env{
		masterToken: InitalToken,
		token:       token,
		config:      config,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)

	server := &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: env.SetupRoutes(),
	}

	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
