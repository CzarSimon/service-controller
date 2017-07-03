package main

import "github.com/CzarSimon/util"

//Config holds configuration values
type Config struct {
	server util.ServerConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9105",
	}
}
