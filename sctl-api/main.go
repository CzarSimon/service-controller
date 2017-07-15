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
	token  sctl.Token
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	return Env{
		db:     connectDB(config.db),
		config: config,
		token:  sctl.NewToken(),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	go env.ScheduleTokenRefresh(config.refreshFrequency)

	server := &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: env.SetupRoutes(),
	}

	log.Println("Starting sctl-api, running on port: " + config.server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
