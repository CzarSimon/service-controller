package main // sctl-minion

import (
	"log"
	"net/http"
	"os"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	token  string
	config Config
}

func setupEnvironment(config Config) Env {
	initalToken := os.Getenv("MINON_TOKEN")
	os.Setenv("MINION_TOKEN", "")
	return Env{
		token:  initalToken,
		config: config,
	}
}

func main() {
	config := getConfig()
	env := setupEnvironment(config)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/set-env", SetEnvVar)
	http.HandleFunc("/reset-token", util.PlaceholderHandler)
	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
