package main

import (
	"ginexamples/config"
	"ginexamples/http"
	"ginexamples/postgres"
	"ginexamples/service/userservice"
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
	repository.AutoMigrate()

	server := http.AppServer{
		UserService: userservice.New(repository.UserRepository),
	}
	server.Run()
}
