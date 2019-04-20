package http

import (
	"ginexamples"
	"log"

	"github.com/gin-gonic/gin"
)

// AppServer contains the information to run a server.
type AppServer struct {
	UserService ginexamples.UserService
	Logger      *log.Logger
	route       *gin.Engine
}

func (a *AppServer) initialize() {
	gin.DisableConsoleColor()
	route := gin.New()
	public := route.Group("/api", Logger(a.Logger), CORS())
	private := route.Group("/api/v1", Logger(a.Logger), CORS(), NewAuthMiddleware(a.UserService))
	a.publicRoutes(public)
	a.privateRoutes(private)

	a.route = route
}

// Run will start the gin server.
func (a *AppServer) Run() {
	a.initialize()
	a.route.Run()
}
