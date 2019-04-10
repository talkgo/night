package http

import (
	"github.com/gin-gonic/gin"
)

func (a *AppServer) publicRoutes(r *gin.Engine) {
	v1 := r.Group("api")
	v1.GET("/", a.HelloHandler)
}
