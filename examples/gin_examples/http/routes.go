package http

import (
	"github.com/gin-gonic/gin"
)

func (a *AppServer) publicRoutes(router *gin.RouterGroup) {
	router.POST("/register", a.RegisterUserHandler)
	router.POST("/login", a.LoginUserHandler)
}
