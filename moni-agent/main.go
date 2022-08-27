package main

import (
	"flag"
	"github.com/DolphinDong/toolkit/moni-agent/Initialize"
	"github.com/DolphinDong/toolkit/moni-agent/cron"
	"github.com/DolphinDong/toolkit/moni-agent/global"
	"github.com/DolphinDong/toolkit/moni-agent/utils"
)

func init() {
	var configPath string
	flag.StringVar(&configPath, "c", "./moni-agent.json", "config file path")
	flag.Parse()
	// 初始化配置文件
	err := Initialize.InitConfig(configPath)
	utils.CheckErr(err)
}
func main() {
	err := cron.NewCronJob(global.GlobalConfig.TaskSchedule).Run()
	utils.CheckErr(err)
}
