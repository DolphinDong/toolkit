package utils

import (
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 从context 中获取logger
func GetLogger(ctx *gin.Context) *logrus.Entry {
	if logger, exists := ctx.Get(common.LoggerKey); exists {
		if l, ok := logger.(*logrus.Entry); ok {
			return l
		} else {
			return &logrus.Entry{}
		}
	} else {
		return &logrus.Entry{}
	}
}
