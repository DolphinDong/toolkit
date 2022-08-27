package middleware

import (
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Token")
		if token == global.GlobalConfig.Token {
			ctx.Next()
			return
		} else {
			ctx.Abort()
			common.ResponseForbidden(ctx, "Authentication failed")
			return
		}

	}
}
