package config

import (
	"os"
)

// Config contains all the environment varialbes.
type Config struct {
	Port       string
	Env        string
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGName     string
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

// GetConfig returns the environment varialbes.
func GetConfig() Config {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "development"
	}

	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		pgHost = "localhost"
	}

	pgPort, ok := os.LookupEnv("PG_PORT")
	if !ok {
		pgPort = "5432"
	}

	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		pgUser = "postgres"
	}

	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		pgPassword = ""
	}

	pgDBName, ok := os.LookupEnv("PG_DB_NAME")
	if !ok {
		pgDBName = "ginexamples"
	}

	return Config{
		Port:       port,
		Env:        env,
		PGHost:     pgHost,
		PGPort:     pgPort,
		PGUser:     pgUser,
		PGPassword: pgPassword,
		PGName:     pgDBName,
	}
}
