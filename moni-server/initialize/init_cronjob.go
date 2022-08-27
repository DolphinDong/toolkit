package initialize

import (
	"github.com/DolphinDong/toolkit/moni-server/cron"
	"github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/utils"
)

func initCronjob() {
	global.GlobalLoger.Info("start init cron job.....")
	err := cron.NewCronJob(global.GlobalConfig.TaskSchedule).Run()
	utils.CheckErr(err)
	global.GlobalLoger.Info("start init cron job success !!")
}
