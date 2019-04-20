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
	PGDBName   string
	LogFile    string
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

	logFile, ok := os.LookupEnv("LOGFILE")
	if !ok {
		logFile = ""
	}

	return Config{
		Port:       port,
		Env:        env,
		PGHost:     pgHost,
		PGPort:     pgPort,
		PGUser:     pgUser,
		PGPassword: pgPassword,
		PGDBName:   pgDBName,
		LogFile:    logFile,
	}
}
