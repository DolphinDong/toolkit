package middleware

import (
	"bytes"
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
		logger := global.GlobalLoger.WithField("requestID", requestID)
		ctx.Set(common.LoggerKey, logger)
		ctx.Set(common.RequestIDKey, requestID)
		ctx.Next()
	}
}

// 打印请求信息
func PrintRequestInfo() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 获取context中的logger
		logger := utils.GetLogger(ctx)

		ip := ctx.RemoteIP()
		url := ctx.Request.URL
		// 请求方法为post需要将请求体内的东西打印出来
		if ctx.Request.Method == http.MethodPost {
			all, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				common.ResponseInternalServerErrorWithMsg(ctx, "read request body error")
				logger.Error("read request body error")
				ctx.Abort()
				return
			}
			defer ctx.Request.Body.Close()
			data := string(all)
			// 将内容重新赋值给 ctx.Request.Body
			closer := ioutil.NopCloser(bytes.NewBuffer(all))
			ctx.Request.Body = closer
			logger.Infof("[%v] IP:%v URL:%v requestBody=%v", ctx.Request.Method, ip, url, data)
		} else {
			logger.Infof("[%v] IP:%v URL:%v", ctx.Request.Method, ip, url)
		}

		// 放行
		ctx.Next()
		// 打印响应状态码
		logger.Infof("response code=[%v]", ctx.Writer.Status())
	}
}
