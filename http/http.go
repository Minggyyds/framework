package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func Res(c *gin.Context, code int, msg string, data interface{}) {
	httpCode := http.StatusOK
	if code > 2000 {
		httpCode = http.StatusBadGateway
	}
	c.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
