package main

import (
	"database/sql"
	"log"
	"net/http"
)

//Env is the struct for environment objects
type Env struct {
	db     *sql.DB
	config Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	return Env{
		db:     connectDB(config.db),
		config: config,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	http.HandleFunc("/init", env.InitProject)
	http.HandleFunc("/add-node", env.AddNode)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/start", env.placeholderHandler)
	http.HandleFunc("/check", env.placeholderHandler)
	http.HandleFunc("/alter", env.placeholderHandler)
	log.Println("Starting sctl-api-server, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
