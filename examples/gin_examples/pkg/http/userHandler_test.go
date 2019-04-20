package http

import (

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestAppServer_RegisterUserHandler(t *testing.T) {
	mockService := mock.UserService{}
	mockService.CreateUserFn = func(u *ginexamples.User, password string) (*ginexamples.User, error) {
		if len(password) < 8 {
			return nil, errors.New("too short")
		}
		u.ID = 1
		return u, nil
	}

	appServer := AppServer{UserService: &mockService, Logger: log.New(os.Stdout, "", 0)}

	var testCases = []struct {
		name   string
		body   string
		status int
	}{
		{"no email", `{"password":"password"}`, 400},
		{"no password", `{"email":"e@mail.com"}`, 400},
		{"password too short", `{"email":"e@mail.com","password":"pass"}`, 500},
		{"malformed json", `{"email":"e@mail.com","password":"pass"`, 400},
		{"success", `{"email":"e@mail.com","password":"password"}`, 200},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			mockService.CreateUserFnInvoked = false

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.POST("/", appServer.RegisterUserHandler)

			req := httptest.NewRequest("POST", "/", strings.NewReader(v.body))
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			if v.status == 400 {
				assert.Equal(t, v.status, resp.Code, "bad response code")
				assert.False(t, mockService.CreateUserFnInvoked, "CreateUser was invoked when it should not")
				assert.Empty(t, resp.Body, "body should be empty")
				return
			}

			assert.True(t, mockService.CreateUserFnInvoked, "CreateUser was not invoked when it should")
			assert.Equal(t, v.status, resp.Code, "bad response code")

			if v.status == 200 {
				assert.NotEmptyf(t, resp.HeaderMap["Set-Cookie"], "cookie was not set on successful registration")
			}
		})
	}
}

func TestAppServer_LoginUserHandler(t *testing.T) {
	mockService := mock.UserService{}
	mockService.LoginFn = func(email string, password string) (*ginexamples.User, error) {
		if password != "password" {
			return nil, errors.New("bad login")
		}
		return &ginexamples.User{Model: gorm.Model{ID: 1}}, nil
	}
	appServer := AppServer{UserService: &mockService, Logger: log.New(os.Stdout, "", 0)}

	var testCases = []struct {
		name   string
		body   string
		status int
	}{
		{"no email", `{"password":"password"}`, 400},
		{"no password", `{"email":"e@mail.com"}`, 400},
		{"bad password", `{"email":"e@mail.com","password":"pass"}`, 401},
		{"malformed json", `{"email":"e@mail.com","password":"pass"`, 400},
		{"success", `{"email":"e@mail.com","password":"password"}`, 200},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			mockService.LoginFnInvoked = false

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.POST("/", appServer.LoginUserHandler)

			req := httptest.NewRequest("POST", "/", strings.NewReader(v.body))
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			if v.status == 400 {
				assert.Equal(t, v.status, resp.Code, "bad response code")
				assert.False(t, mockService.LoginFnInvoked, "Login was invoked when it should not")
				assert.Empty(t, resp.Body, "body should be empty")
				return
			}

			assert.True(t, mockService.LoginFnInvoked, "Login was not invoked when it should")
			assert.Equal(t, v.status, resp.Code, "bad response code")

			if v.status == 200 {
				assert.NotEmptyf(t, resp.HeaderMap["Set-Cookie"], "cookie was not set on successful registration")
			}
		})
	}
}

func TestAppServer_LogoutUserHandler(t *testing.T) {
	mockService := mock.UserService{}
	mockService.LogoutFn = func(sessionID string) error {
		if sessionID != "session" {
			return errors.New("not found")
		}
		return nil
	}
	appServer := AppServer{UserService: &mockService, Logger: log.New(os.Stdout, "", 0)}

	var testCases = []struct {
		name      string
		sessionID string
		status    int
	}{
		{"no sessionID", "", 200},
		{"bad sessionID", "535510N", 500},
		{"success", "session", 200},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			mockService.LogoutFnInvoked = false

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.POST("/", appServer.LogoutUserHandler)

			req := httptest.NewRequest("POST", "/", nil)
			c := http.Cookie{Name: "sessionID", Value: v.sessionID}
			req.AddCookie(&c)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, v.status, resp.Code, "bad response code")
			assert.Empty(t, resp.Body, "body should be empty")
			if v.sessionID == "" {
				assert.False(t, mockService.LogoutFnInvoked, "Logout was invoked when it should not")
				assert.Equal(t, v.status, resp.Code, "bad response code")
				return
			}
			assert.True(t, mockService.LogoutFnInvoked, "Logout was not invoked when it should")

			if v.status == 200 {
				assert.NotEmptyf(t, resp.HeaderMap["Set-Cookie"], "cookie was not set on successful registration")
			}
		})
	}
}

func TestAppServer_GetUserHandler(t *testing.T) {
	mockService := mock.UserService{}
	mockService.GetUserFn = func(id string) (*ginexamples.User, error) {
		if id != "userId" {
			return nil, errors.New("not found")
		}
		return &ginexamples.User{Model: gorm.Model{ID: 1}}, nil
	}
	appServer := AppServer{UserService: &mockService, Logger: log.New(os.Stdout, "", 0)}

	var testCases = []struct {
		name   string
		userID string
		status int
	}{
		{"no userID", "", 400},
		{"bad userID", "someId", 404},
		{"success", "userId", 200},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			mockService.GetUserFnInvoked = false

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/", appServer.GetUserHandler)
			r.GET("/:id", appServer.GetUserHandler)

			req := httptest.NewRequest("GET", "/"+v.userID, nil)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, v.status, resp.Code, "bad response code")
			if v.userID == "" {
				assert.False(t, mockService.GetUserFnInvoked, "GetUser was invoked when it should not")
				assert.Equal(t, v.status, resp.Code, "bad response code")
				assert.Empty(t, resp.Body, "body should not be empty")
				return
			}
			assert.True(t, mockService.GetUserFnInvoked, "GetUser was not invoked when it should")
			if v.status == 200 {
				assert.NotEmptyf(t, resp.Body, "body should not be empty")
			}
		})
	}
}
func TestAppServer_GetMeHandler(t *testing.T) {
	mockService := mock.UserService{}
	mockService.GetUserFn = func(id string) (*ginexamples.User, error) {
		if id != "1" {
			return nil, errors.New("not found")
		}
		return &ginexamples.User{Model: gorm.Model{ID: 1}, Email: "heisenberg@gmail.com", Name: "heisenberg"}, nil
	}
	appServer := AppServer{UserService: &mockService, Logger: log.New(os.Stdout, "", 0)}

	var testCases = []struct {
		name   string
		userID string
		status int
	}{
		{"no userID", "", 400},
		{"bad userID", "someId", 404},
		{"success", "1", 200},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			mockService.GetUserFnInvoked = false

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.Use(func(c *gin.Context) {
				c.Set("userID", v.userID)
			})
			r.GET("/", appServer.GetMeHandler)
			req := httptest.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, v.status, resp.Code, "bad response code")
			if v.userID == "" {
				assert.False(t, mockService.GetUserFnInvoked, "GetUser was invoked when it should not")
				assert.Equal(t, v.status, resp.Code, "bad response code")
				assert.Empty(t, resp.Body, "body should not be empty")
				return
			}
			assert.True(t, mockService.GetUserFnInvoked, "GetUser was not invoked when it should")
			if v.status == 200 {
				assert.NotEmptyf(t, resp.Body, "body should not be empty")
				assert.Contains(t, resp.Body.String(), "heisenberg@gmail.com", "Response Body does not contain profilePicture")
			}
		})
	}
}
