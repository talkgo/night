package http

import (
	"gin_examples"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	UserService gin_examples.UserService
	route       *gin.Engine
}

func (a *AppServer) initialize() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	// Write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	route := gin.New()

	a.publicRoutes(route)

	a.route = route
}

func (a *AppServer) Run() {
	a.initialize()
	a.route.Run()
}
