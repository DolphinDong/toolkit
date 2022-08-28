package cron

import (
	thttp "github.com/DolphinDong/gotools/http"
	"github.com/DolphinDong/toolkit/moni-agent/global"
	"github.com/pkg/errors"
	"log"
)

func getTaskFunc() (taskFunc map[string]func()) {
	taskFunc = make(map[string]func())
	taskFunc["ConnectServer1"] = ConnectServer
	taskFunc["ConnectServer2"] = ConnectServer
	taskFunc["ConnectServer3"] = ConnectServer
	return
}

func ConnectServer() {
	serverClient := thttp.Client{
		BaseUrl: global.GlobalConfig.MoniServerHost,
	}
	header := map[string]string{
		"Token": global.GlobalConfig.MoniServerToken,
	}
	for i := 0; i < 20; i++ {
		response, err := serverClient.Get("/hello", nil, header)
		if err != nil {
			log.Printf("%+v", errors.WithStack(err))
		}
		log.Printf("server response: %+v", string(response))
	}
}
