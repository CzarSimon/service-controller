package main // sctl-minion

import (
	"github.com/CzarSimon/util"
	"github.com/kardianos/osext"
)

// Config holds configuration values
type Config struct {
	server   util.ServerConfig
	SSL      SSLConfig
	DBFolder string
	Lock     LockConfig
}

func getConfig() Config {
	basePath, err := osext.ExecutableFolder()
	util.CheckErrFatal(err)
	return Config{
		server:   getServerConfig(),
		SSL:      getSSLConfig(),
		DBFolder: basePath,
		Lock:     getLockConfig(),
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

// LockConfig Holds config values for the minion lock
type LockConfig struct {
	TokenMaxAge float64
	MaxAttempts int
}

func getLockConfig() LockConfig {
	return LockConfig{
		TokenMaxAge: 360.0,
		MaxAttempts: 3,
	}
}
