package initialize

import (
	"encoding/json"
	tutils "github.com/DolphinDong/gotools/utils"
	"github.com/DolphinDong/toolkit/moni-server/config"
	"github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/pkg/errors"
	"os"
)

// 初始化配置文件
func initConfig(configPath string) {
	global.GlobalConfig = new(config.Server)
	content, err := os.ReadFile(configPath)
	if err != nil {
		global.GlobalLoger.Fatalf("%+v", errors.WithStack(err))
	}
	err = json.Unmarshal(content, global.GlobalConfig)
	if err != nil {
		global.GlobalLoger.Fatalf("%+v", errors.WithStack(err))
	}
	// 校验配置文件
	err = tutils.Validate(global.GlobalConfig)
	if err != nil {
		global.GlobalLoger.Fatalf("%+v", errors.WithStack(err))
	}

}
