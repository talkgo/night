package http

import (
	"ginexamples"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AppServer) RegisterUserHandler(c *gin.Context) {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var (
		userModel ginexamples.User
		req       request
	)

	err := c.BindJSON(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel.Email = req.Email
	userModel.Name = req.Name

	user, err := a.UserService.CreateUser(&userModel, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}
