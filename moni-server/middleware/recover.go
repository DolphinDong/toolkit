package middleware

import (
	"fmt"
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/DolphinDong/toolkit/moni-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 捕获错误
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := utils.GetLogger(ctx)
				common.ResponseInternalServerErrorWithMsg(ctx, "server Error: "+fmt.Sprintf("%v", err))
				if err2, ok := err.(error); ok {
					logger.Errorf("%+v", errors.WithStack(err2))
				} else {
					logger.Errorf("%+v", errors.New(fmt.Sprintf("%+v", err)))
				}
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
