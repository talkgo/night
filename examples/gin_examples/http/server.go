package http

import (
	"ginexamples"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// AppServer contains the information to run a server.
type AppServer struct {
	UserService ginexamples.UserService
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

// Run will start the gin server.
func (a *AppServer) Run() {
	a.initialize()
	a.route.Run()
}
