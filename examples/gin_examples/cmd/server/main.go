package main

import (
	"gin_examples/config"
	"gin_examples/http"
	"gin_examples/postgres"
	"gin_examples/service/userService"
)

func main() {
	cfg := config.GetConfig()

	postgresConfig := postgres.DBConfig{
		Host:     cfg.PGHost,
		Port:     cfg.PGPort,
		User:     cfg.PGUser,
		Password: cfg.PGPassword,
		Name:     cfg.PGName,
	}

	repository := postgres.Initialize(postgresConfig)

	server := http.AppServer{
		UserService: userService.New(repository.UserRepository),
	}
	server.Run()
}
