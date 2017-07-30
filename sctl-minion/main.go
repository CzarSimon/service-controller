package main // sctl-minion

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	sctl "github.com/CzarSimon/sctl-common"
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

func testEnv() {
	path := os.Getenv("PATH")
	fmt.Println(strings.Replace(path, ":", "\n", -1))
	fmt.Println(os.Getenv("USER"))
	checkDocker := sctl.Command{
		Main: "which",
		Args: []string{"docker"},
	}
	dockerPath, _ := checkDocker.Execute()
	fmt.Println(dockerPath)
	listImages := sctl.DockerCommand([]string{"images"})
	fmt.Println(listImages.ToString())
	out, _ := listImages.Execute()
	fmt.Println(out)
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
	testEnv()
	config := getConfig()
	env := SetupEnv(config)
	server := SetupServer(env, config)

	log.Println("Starting sctl-minion, running on port: " + config.server.Port)
	err := server.ListenAndServeTLS(config.SSL.Cert, config.SSL.Key)
	util.CheckErr(err)
}
