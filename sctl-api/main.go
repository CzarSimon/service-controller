package main // sctl-api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	db     *sql.DB
	config Config
	token  string
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	return Env{
		db:     connectDB(config.db),
		config: config,
		token:  sctl.GenerateToken(1),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	http.HandleFunc("/init", env.InitProject)
	http.HandleFunc("/add-node", env.AddNode)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/start", util.PlaceholderHandler)
	http.HandleFunc("/check", util.PlaceholderHandler)
	http.HandleFunc("/alter", util.PlaceholderHandler)
	http.HandleFunc("/set-env", env.ForwardEnvVar)
	http.HandleFunc("/active-project", env.ActiveProject)
	http.HandleFunc("/project-list", env.GetProjectList)
	log.Println("Starting sctl-api-server, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
