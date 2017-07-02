package main

import (
	"database/sql"
	"fmt"

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

func getSchema() []string {
	return []string{
		`CREATE TABLE PROJECT(
      NAME VARCHAR(50) PRIMARY KEY,
      FOLDER VARCHAR(300),
      SWARM_TOKEN VARCHAR(100),
      IS_ACTIVE BOOLEAN
    )`,
		`CREATE TABLE NODE(
      PROJECT VARCHAR(50),
      IP VARCHAR(50),
      OS VARCHAR(10) DEFAULT 'linux',
      IS_MASTER BOOLEAN,
      FOREIGN KEY (PROJECT) REFERENCES PROJECT(NAME),
      PRIMARY KEY (PROJECT, IP)
    )`,
	}
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
	schema := getSchema()
	for _, tableDef := range schema {
		_, err := db.Exec(tableDef)
		util.CheckErrFatal(err)
	}
}
