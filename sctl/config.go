package main // sctl-cli

import "github.com/CzarSimon/util"

//Config holds configuration values
type Config struct {
	API util.ServerConfig
}

func getConfig() Config {
	return Config{
		API: getAPIConfig(),
	}
}

func getAPIConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9104",
	}
}
