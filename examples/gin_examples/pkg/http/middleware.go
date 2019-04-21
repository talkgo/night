package http

import (
	"ginexamples"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(provider ginexamples.UserAuthenticationProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		user, err := provider.CheckAuthentication(sessionID)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("userID", user.ID)
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
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func Logger(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Printf("%s: %s %s", c.Request.RemoteAddr, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
