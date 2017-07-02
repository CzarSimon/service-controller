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

func setupEnvironment(config Config) Env {
	return Env{
		db:     connectDB(config.db),
		config: config,
	}
}

func getActiveProject(db *sql.DB) (string, error) {
	var projectName string
	query := "SELECT NAME FROM PROJECT WHERE IS_ACTIVE=1"
	err := db.QueryRow(query).Scan(&projectName)
	return projectName, err
}

func main() {
	config := getConfig()
	env := setupEnvironment(config)
	http.HandleFunc("/init", env.InitProject)
	http.HandleFunc("/add-node", env.AddNode)
	http.HandleFunc("/update", env.UpdateImage)
	http.HandleFunc("/start", env.placeholderHandler)
	http.HandleFunc("/check", env.placeholderHandler)
	http.HandleFunc("/alter", env.placeholderHandler)
	log.Println("Starting sctl-api-server, running on port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
