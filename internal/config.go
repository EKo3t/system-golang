package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Postgres struct {
		host string
		port string
		username string
		password string
		dbName string
	}
	Server struct {
		host string
		port uint16
	}
}

func LoadAppConfig(configPath string) (*AppConfig, error) {
	var appConfig AppConfig

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Can't open or read file at specified path: %s", configPath)
	}

	err = yaml.Unmarshal(file, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
