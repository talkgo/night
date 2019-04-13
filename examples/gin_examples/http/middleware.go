package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Printf("%s: %s %s", c.Request.RemoteAddr, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			return
		}

		c.Next()
	}
}
