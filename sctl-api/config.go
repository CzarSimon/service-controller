package main // sctl-api

import (
	"database/sql"
	"fmt"

	"github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	_ "github.com/mattn/go-sqlite3"
)

//Config holds configuration values
type Config struct {
	server util.ServerConfig
	minion util.ServerConfig
	db     util.SQLiteConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		minion: getMinionConfig(),
		db:     util.GetSQLiteConfig("./database/sctl-db"),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9104",
	}
}

func getMinionConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Port:     "9105",
	}
}

// GetSchema returns the database schema for sctl-api-server
func GetSchema() []string {
	return []string{sctl.ProjectSchema(), sctl.NodeSchema()}
}

func connectDB(config util.SQLiteConfig) *sql.DB {
	db, dbExists := connectSQLlite(config)
	if !dbExists {
		fmt.Println("New db")
		installSchema(db)
	}
	return db
}

func connectSQLlite(config util.SQLiteConfig) (*sql.DB, bool) {
	dbExists, err := util.FileExists(config.File)
	util.CheckErrFatal(err)
	db, err := sql.Open("sqlite3", config.File)
	util.CheckErrFatal(err)
	err = db.Ping()
	util.CheckErrFatal(err)
	return db, dbExists
}

func installSchema(db *sql.DB) {
	schema := GetSchema()
	for _, tableDef := range schema {
		_, err := db.Exec(tableDef)
		util.CheckErrFatal(err)
	}
}
