package Initialize

import (
	"encoding/json"
	tutils "github.com/DolphinDong/gotools/utils"
	"github.com/DolphinDong/toolkit/moni-agent/config"
	"github.com/DolphinDong/toolkit/moni-agent/global"
	"github.com/pkg/errors"
	"os"
)

// 初始化配置文件
func InitConfig(configPath string) error {
	global.GlobalConfig = new(config.Config)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return errors.WithStack(err)
	}
	err = json.Unmarshal(content, global.GlobalConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	// 校验配置文件
	err = tutils.Validate(global.GlobalConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
