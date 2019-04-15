package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reading-go/examples/gin_examples/e"
)

type G struct {
	C *gin.Context
}

type GT = *gin.Context

type GE = *gin.Engine

func (g *G) Send(code int, data interface{}) {

	g.C.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	return
}
