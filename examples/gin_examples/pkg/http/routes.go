package http

import (
	"github.com/gin-gonic/gin"
)

func (a *AppServer) privateRoutes(router *gin.RouterGroup) {
	router.GET("me", a.GetMeHandler)
	router.GET("users/:id", a.GetUserHandler)
}

func (a *AppServer) publicRoutes(router *gin.RouterGroup) {
	router.POST("/register", a.RegisterUserHandler)
	router.POST("/login", a.LoginUserHandler)
	router.POST("/logout", a.LogoutUserHandler)
}
