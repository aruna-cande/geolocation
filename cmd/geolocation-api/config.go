package main

import (
	"os"
	"strconv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresHost     string
	PostgresPort     int64
	ApiPort          int64
}

func NewConfig() Config {
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	if err != nil {
		return Config{}
	}
	apiPort, err := strconv.ParseInt(os.Getenv("API_PORT"), 10, 64)
	if err != nil {
		return Config{}
	}
	return Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     port,
		ApiPort:          apiPort,
	}
}
