package main // sctl-minion

import "github.com/CzarSimon/util"

// Config holds configuration values
type Config struct {
	server   util.ServerConfig
	SSL      SSLConfig
	DBFolder string
}

func getConfig() Config {
	return Config{
		server:   getServerConfig(),
		SSL:      getSSLConfig(),
		DBFolder: ".",
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "https",
		Host:     "localhost",
		Port:     "9105",
	}
}

// SSLConfig Path info for SSL key and certificate
type SSLConfig struct {
	Folder string
	Key    string
	Cert   string
}

func getSSLConfig() SSLConfig {
	return SSLConfig{
		Folder: "ssl",
		Key:    "./ssl/sctl.key",
		Cert:   "./ssl/sctl.crt",
	}
}
