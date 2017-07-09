package main // sctl-minion

import "github.com/CzarSimon/util"

// Config holds configuration values
type Config struct {
	server util.ServerConfig
	SSL    SSLConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		SSL:    getSSLConfig(),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     "localhost",
		Port:     "9105",
	}
}

// SSLConfig Path info for SSL key and certificate
type SSLConfig struct {
	Key  string
	Cert string
}

func getSSLConfig() SSLConfig {
	return SSLConfig{
		Key:  "./ssl/minion.key",
		Cert: "./ssl/minion.cert",
	}
}
