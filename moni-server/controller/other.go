package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OtherController struct {
}

func (*OtherController) NoRouter(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"msg":  "Resource not found",
	})
}

func (*OtherController) GetHelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "hello  moni-server!!",
	})
}
