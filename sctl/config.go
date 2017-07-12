package main // sctl-cli

import "github.com/CzarSimon/util"

//Config holds configuration values
type Config struct {
	API      util.ServerConfig
	Version  string
	AppName  string
	AppUsage string
}

func getConfig() Config {
	return Config{
		API:      getAPIConfig(),
		Version:  "0.0.1",
		AppName:  "service controller (sctl)",
		AppUsage: "Command line tool for simplifying running services using docker swarm",
	}
}

func getAPIConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9104",
	}
}
