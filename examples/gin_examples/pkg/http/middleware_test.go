package http

import (
	"bytes"
	"errors"
	"ginexamples"
	"ginexamples/pkg/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthMiddleware(t *testing.T) {
	var as = &mock.UserAuthenticationProvider{
		CheckAuthenticationFn: func(sessionID string) (*ginexamples.User, error) {
			if sessionID == "" {
				return nil, errors.New("empty sessionID")
			}
			if sessionID == "invalid" {
				return nil, errors.New("not found")
			}
			return &ginexamples.User{}, nil
		},
	}

	var testCases = []struct {
		name   string
		enter  bool
		status int
		input  string
		cookie bool
	}{
		{"empty sessionid", false, 403, "", true},
		{"invalid sessionid", false, 403, "invalid", true},
		{"no cookie", false, 403, "whatever", false},
		{"success", true, 200, "valid", true},
	}

	for _, v := range testCases {
		handlerEntered := false

		var testHandler gin.HandlerFunc = func(c *gin.Context) {
			assert.True(t, v.enter, "handler should not have been entered")
			handlerEntered = true
			c.Status(http.StatusOK)
		}

		t.Run(v.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.Use(NewAuthMiddleware(as), testHandler)

			req := httptest.NewRequest("", "/", nil)
			if v.cookie {
				req.AddCookie(&http.Cookie{Name: "sessionID", Value: v.input})
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)
			assert.Equal(t, v.status, resp.Code, "status code does not match")
			assert.True(t, as.CheckAuthenticationFnInvoked, "authentication was not actually checked")
			assert.Equal(t, v.enter, handlerEntered, "handler was not called as expected")
		})
	}
}

func TestLogger(t *testing.T) {
	var testCases = []struct {
		name   string
		method string
		path   string
	}{
		{"login", "POST", "/api/login"},
		{"getCourses", "GET", "/api/v1/me"},
		{"CORS", "OPTIONS", "/"},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			handlerEntered := false

			var testHandler gin.HandlerFunc = func(c *gin.Context) {
				handlerEntered = true
				c.Status(http.StatusOK)
			}

			b := &bytes.Buffer{}
			l := log.New(b, "", 0)
			logMiddleWare := Logger(l)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.Use(logMiddleWare, testHandler)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(v.method, v.path, nil)
			r.ServeHTTP(resp, req)
			assert.True(t, handlerEntered, "inner handler was not entered")
			assert.Contains(t, b.String(), v.path, "log does not contain path")
			assert.Contains(t, b.String(), v.method, "log does not contain method")
		})
	}
}

func TestCORS(t *testing.T) {
	var testCases = []struct {
		name   string
		method string
		enter  bool
	}{
		{"GET", "GET", true},
		{"PUT", "PUT", true},
		{"POST", "POST", true},
		{"OPTIONS", "OPTIONS", false},
	}

	var expectedHeader = map[string]string{
		"Access-Control-Allow-Origin":      "http://localhost:8080",
		"Access-Control-Allow-Methods":     "POST, GET, OPTIONS, PUT, DELETE",
		"Access-Control-Allow-Headers":     "Accept, Content-Type, Content-Length, Accept-Encoding",
		"Access-Control-Allow-Credentials": "true",
	}

	for _, v := range testCases {
		handlerEntered := false

		var testHandler gin.HandlerFunc = func(c *gin.Context) {
			assert.True(t, v.enter, "handler should not have been entered")
			handlerEntered = true
			c.Status(http.StatusOK)
		}

		t.Run(v.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.Use(CORS(), testHandler)

			req := httptest.NewRequest(v.method, "/", nil)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			res := resp.Result()
			h := res.Header

			for k, v := range expectedHeader {
				assert.Equalf(t, v, h.Get(k), "header %s does not match expected value", k)
			}
			assert.Equal(t, http.StatusOK, res.StatusCode, "http status not 200")
			assert.Equal(t, v.enter, handlerEntered, "next handler was not called as expected")
		})
	}
}
