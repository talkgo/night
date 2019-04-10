package http

import "github.com/gin-gonic/gin"

func (a *AppServer) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
