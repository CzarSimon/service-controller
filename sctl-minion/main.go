package main // sctl-minion

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	masterToken string
	token       sctl.Token
	lock        MinionLock
	config      Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	tokens := sctl.GetTokenBundle(config.DBFolder)
	return Env{
		masterToken: tokens.Master,
		token:       tokens.Auth,
		lock:        NewLock(config.Lock),
		config:      config,
	}
}

// SetupServer Genreates certificates and returns tls configured server with routes setup
func SetupServer(env Env, config Config) *http.Server {
	config.SSL.CertGen()
	return &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: env.SetupRoutes(),
		TLSConfig: &tls.Config{
			ServerName: "sctl-minion",
		},
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
}

func main() {
	fmt.Println(os.Getenv("USER"))
	fmt.Println(strings.Replace(os.Getenv("PATH"), ":", "\n", -1))
	config := getConfig()
	env := SetupEnv(config)
	server := SetupServer(env, config)

	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	err := server.ListenAndServeTLS(config.SSL.Cert, config.SSL.Key)
	util.CheckErr(err)
}
