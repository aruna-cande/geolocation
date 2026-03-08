package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration for the geolocation API.
type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresHost     string
	PostgresPort     int64
	APIPort          int64
}

// NewConfig reads configuration from environment variables.
// Returns an error if required numeric values cannot be parsed.
func NewConfig() (Config, error) {
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("parsing POSTGRES_PORT: %w", err)
	}
	apiPort, err := strconv.ParseInt(os.Getenv("API_PORT"), 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("parsing API_PORT: %w", err)
	}
	return Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     port,
		APIPort:          apiPort,
	}, nil
}
