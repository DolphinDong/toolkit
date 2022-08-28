package controller

import (
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/DolphinDong/toolkit/moni-server/global"
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
	conn := global.GlobalRedisPool.Get()
	defer conn.Close()
	times, err := conn.Do("INCR", "times")
	if err != nil {
		common.ResponseInternalServerErrorWithMsg(ctx, "命令执行失败")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  times,
	})
}
