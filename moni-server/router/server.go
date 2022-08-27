package router

import (
	"fmt"
	. "github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/initialize"
	"github.com/DolphinDong/toolkit/moni-server/middleware"
	"github.com/DolphinDong/toolkit/moni-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func init() {
	initialize.Init()
}
func Run() {
	engine := gin.New()

	// 注册中间件
	registerMiddleware(engine)
	// 注册路由
	registerRouter(engine)
	err := engine.Run(fmt.Sprintf("%v:%v", GlobalConfig.ServerHost, GlobalConfig.ServerPort))
	utils.CheckErr(errors.Wrap(err, "start server error!!"))
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) {
	// 日志初始化
	engine.Use(middleware.LogMiddleware())
	// recover
	engine.Use(middleware.Recover())
	// 打印日志
	engine.Use(middleware.PrintRequestInfo())
	// auth
	engine.Use(middleware.Auth())
}
