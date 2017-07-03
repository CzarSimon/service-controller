package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	token  string
	config Config
}

func setupEnvironment(config Config) Env {
	return Env{
		token:  "dummy-token",
		config: config,
	}
}

func main() {
	config := getConfig()
	env := setupEnvironment(config)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/reset-token", util.PlaceholderHandler)
	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
