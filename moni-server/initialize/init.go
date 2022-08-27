package initialize

import (
	"flag"
	. "github.com/DolphinDong/toolkit/moni-server/global"
)

const (
	logo = `
███╗   ███╗ ██████╗ ███╗   ██╗██╗      ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
████╗ ████║██╔═══██╗████╗  ██║██║      ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
██╔████╔██║██║   ██║██╔██╗ ██║██║█████╗███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
██║╚██╔╝██║██║   ██║██║╚██╗██║██║╚════╝╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
██║ ╚═╝ ██║╚██████╔╝██║ ╚████║██║      ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝      ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
	Version: %v    Author: 刘冬`

	version = "v1.0.1"
)

func Init() {
	// 初始化log
	initBaseLog()
	// 初始化配置文件
	var configPath string
	flag.StringVar(&configPath, "c", "./moni-server.json", "The configuration file path for moni-server")
	flag.Parse()
	initConfig(configPath)
	// 升级第二部分日志
	initProLog()
	// 初始化定时任务日志
	initCronLog()

	GlobalLoger.Info("starting server....")
	GlobalLoger.Infof(logo, version)
	// 初始化数据库连接
	initMysql()
	// 初始化redis连接
	initRedis()
	// 数据迁移
	Migrate()
	// 初始化定时任务
	initCronjob()
	//RestoreAssets()
	GlobalLoger.Info("server init success!!!")
}
