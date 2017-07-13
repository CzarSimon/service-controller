package main // sctl-cli

import (
	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

//Config holds configuration values
type Config struct {
	API     util.ServerConfig
	DB      util.SQLiteConfig
	App     AppConfig
	Folders FolderConfig
}

// FolderConfig holds local and node folder info
type FolderConfig struct {
	Exec   string
	Target string
	Init   string
}

func getConfig() Config {
	return Config{
		API: getAPIConfig(),
		App: getAppConfig(),
		Folders: FolderConfig{
			Exec:   "./executables/sctl-minion",
			Target: "../sctl-minion",
			Init:   "../sctl-minion",
		},
	}
}

// AppConfig holds app metadata
type AppConfig struct {
	Version string
	Name    string
	Usage   string
}

func getAppConfig() AppConfig {
	return AppConfig{
		Version: "0.0.1",
		Name:    "service controller (sctl)",
		Usage:   "Command line tool for simplifying running services using docker swarm",
	}
}

func getAPIConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9104",
	}
}

// SetupMinonDB Sets up and populates a minion db for tokens
func (env Env) SetupMinonDB() {
	tokens := env.GetTokens()
	db := sctl.ConnectTokenDB(env.config.Folders.Exec)
	defer db.Close()
	_, err := db.Exec("DROP TABLE IF EXISTS TOKEN")
	util.CheckErrFatal(err)
	_, err = db.Exec(sctl.TokenSchema())
	util.CheckErrFatal(err)
	stmt, err := db.Prepare("INSERT INTO TOKEN(AUTH, MASTER, AUTH_TIMESTAMP) VALUES($1,$2,$3)")
	util.CheckErrFatal(err)
	_, err = stmt.Exec(tokens.Auth.Data, tokens.Master, tokens.Auth.Timestamp)
	util.CheckErrFatal(err)
}
