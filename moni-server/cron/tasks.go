package cron

import (
	"github.com/DolphinDong/toolkit/moni-server/global"
)

func getTaskFunc() (taskFunc map[string]func()) {
	taskFunc = make(map[string]func())
	taskFunc["test01"] = test01
	taskFunc["test02"] = test02
	return
}
func test01() {
	global.GlobalCronLoger.Infof("test01......")
}
func test02() {
	global.GlobalCronLoger.Infof("test02......")
}
