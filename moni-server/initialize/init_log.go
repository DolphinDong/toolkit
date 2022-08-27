package initialize

import (
	"fmt"
	. "github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/utils"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	_ "github.com/zput/zxcTool/ztLog/zt_formatter"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	logFileName     = "moni_server.log"
	logCronFileName = "moni_server_cron.log"
	DEBUG           = "DEBUG"
	INFO            = "INFO"
	ERROR           = "ERROR"
)

// 初始化基础的日志
func initBaseLog() {
	GlobalLoger = logrus.New()

	GlobalLoger.Formatter = &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		NoFieldsColors:  true,
		ShowFullLevel:   true,
	}
}

// 初始化pro版的日志
func initProLog() {
	if GlobalConfig.Log.LogPath == "" {
		GlobalConfig.Log.LogPath = "logs"
	}
	// 创建日志存放的目录
	err := os.MkdirAll(GlobalConfig.Log.LogPath, 0775)
	utils.CheckErr(errors.Wrap(err, "make log dir err"))
	filePath := filepath.Join(GlobalConfig.Log.LogPath, logFileName)
	// 设置日志打印位置
	logfile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	utils.CheckErr(errors.Wrap(err, "open "+filePath+" err"))
	GlobalLoger.SetOutput(io.MultiWriter(logfile, os.Stdout))
	// 设置日志级别
	logLevel := strings.ToUpper(GlobalConfig.Log.LogLevel)
	switch logLevel {
	case DEBUG:
		GlobalLoger.SetLevel(logrus.DebugLevel)
	case INFO:
		GlobalLoger.SetLevel(logrus.InfoLevel)
	case ERROR:
		GlobalLoger.SetLevel(logrus.ErrorLevel)
	default:
		utils.CheckErr(errors.New(fmt.Sprintf("log.log_level=\"%v\" is illegal", GlobalConfig.Log.LogLevel)))
	}
	GlobalLoger.Info("init logger success!!")
}

func initCronLog() {
	GlobalCronLoger = logrus.New()

	GlobalCronLoger.Formatter = &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		NoFieldsColors:  true,
		ShowFullLevel:   true,
	}
	if GlobalConfig.Log.LogPath == "" {
		GlobalConfig.Log.LogPath = "logs"
	}
	// 创建日志存放的目录
	err := os.MkdirAll(GlobalConfig.Log.LogPath, 0775)
	utils.CheckErr(errors.Wrap(err, "make log dir err"))
	filePath := filepath.Join(GlobalConfig.Log.LogPath, logCronFileName)
	// 设置日志打印位置
	logfile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	utils.CheckErr(errors.Wrap(err, "open "+filePath+" err"))
	GlobalCronLoger.SetOutput(io.MultiWriter(logfile))
	// 设置日志级别
	logLevel := strings.ToUpper(GlobalConfig.Log.LogLevel)
	switch logLevel {
	case DEBUG:
		GlobalCronLoger.SetLevel(logrus.DebugLevel)
	case INFO:
		GlobalCronLoger.SetLevel(logrus.InfoLevel)
	case ERROR:
		GlobalCronLoger.SetLevel(logrus.ErrorLevel)
	default:
		utils.CheckErr(errors.New(fmt.Sprintf("log.log_level=\"%v\" is illegal", GlobalConfig.Log.LogLevel)))
	}
	GlobalCronLoger.Info("init  cron logger success!!")
}
