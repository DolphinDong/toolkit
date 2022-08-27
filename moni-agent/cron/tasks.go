package cron

import "fmt"

func getTaskFunc() (taskFunc map[string]func()) {
	taskFunc = make(map[string]func())
	taskFunc["test01"] = test01
	taskFunc["test02"] = test02
	return
}
func test01() {
	fmt.Println("test01")
}
func test02() {
	fmt.Println("test02")
}
