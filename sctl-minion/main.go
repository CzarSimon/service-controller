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
	lock        MinionLock
	config      Config
}

// SetupEnv Initalizes environment based on config
func SetupEnv(config Config) Env {
	tokens := sctl.GetTokenBundle(".")
	return Env{
		masterToken: tokens.Master,
		token:       tokens.Auth,
		lock:        NewLock(config.Lock),
		config:      config,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	//config.SSL.CertGen()

	//certManager := GetCertManager(config.SSL)
	server := &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: env.SetupRoutes(),
	}
	/*
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			ServerName:     "sctl",
		},
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	*/
	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	err := server.ListenAndServe() //TLS(config.SSL.Cert, config.SSL.Key)
	util.CheckErr(err)
}
