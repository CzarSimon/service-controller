package main // sctl-minion

import (
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

// InitalToken Holds the inital token value
const InitalToken string = "INITAL_TOKEN"

//Env is the struct for environment objects
type Env struct {
	masterToken string
	token       string
	config      Config
}

func setupEnvironment(config Config) Env {
	return Env{
		masterToken: InitalToken,
		token:       InitalToken,
		config:      config,
	}
}

func main() {
	config := getConfig()
	env := setupEnvironment(config)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/set-env", SetEnvVar)
	http.HandleFunc("/init", env.SetupMaster)
	http.HandleFunc("/reset-token", util.PlaceholderHandler)
	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
