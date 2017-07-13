package main // sctl-minion

import (
	"log"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	masterToken string
	token       sctl.Token
	config      Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	tokens := sctl.GetTokenBundle(".")
	return Env{
		masterToken: tokens.Master,
		token:       tokens.Auth,
		config:      config,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	config.SSL.CertGen()

	server := &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: env.SetupRoutes(),
	}

	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	err := server.ListenAndServeTLS(config.SSL.Cert, config.SSL.Key)
	util.CheckErr(err)
}
