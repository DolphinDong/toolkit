package router

import (
	"github.com/DolphinDong/toolkit/moni-server/controller"
	"github.com/gin-gonic/gin"
)

func registerRouter(engine *gin.Engine) {

	{
		otherController := &controller.OtherController{}
		engine.NoRoute(otherController.NoRouter)
		engine.GET("/hello", otherController.GetHelloWorld)
	}
	// 其他路由

}
