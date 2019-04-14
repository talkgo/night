package http

import (
	"fmt"
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

	var (
		userModel ginexamples.User
		req       request
	)

	err := c.BindJSON(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	userModel.Email = req.Email
	userModel.Name = req.Name

	user, err := a.UserService.CreateUser(&userModel, req.Password)
	if err != nil {
		a.Logger.Printf("error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	setCookie(c, user.SessionID)
	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
	})
}

func (a *AppServer) LoginUserHandler(c *gin.Context) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req request
	err := c.BindJSON(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := a.UserService.Login(req.Email, req.Password)
	if err != nil {
		a.Logger.Printf("error logging in: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	setCookie(c, user.SessionID)
	c.JSON(http.StatusOK, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}

func (a *AppServer) LogoutUserHandler(c *gin.Context) {
	sessionID, err := c.Cookie("sessionID")
	if err != nil || sessionID == "" {
		c.Status(http.StatusOK)
		return
	}

	err = a.UserService.Logout(sessionID)
	if err != nil {
		a.Logger.Printf("error logging out user %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	setCookie(c, "")
	c.Status(http.StatusOK)
}

func (a *AppServer) GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := a.UserService.GetUser(id)
	if err != nil {
		a.Logger.Printf("error getting user %v", err)
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}

func (a *AppServer) GetMeHandler(c *gin.Context) {
	id, exists := c.Get("userID")
	if id == "" || exists == false {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := a.UserService.GetUser(fmt.Sprintf("%v", id))
	if err != nil {
		a.Logger.Printf("error getting user %v", err)
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}

func setCookie(c *gin.Context, value string) {
	c.SetCookie("sessionID", value, 86400, "/", "localhost", false, true)
}
