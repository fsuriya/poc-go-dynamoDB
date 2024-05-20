package app

import (
	"poc-go-dynamodb/pkg/error"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: error.GetMsg(errCode),
		Data:    data,
	})
}
