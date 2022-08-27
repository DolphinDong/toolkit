package initialize

import (
	. "github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/utils"
	"github.com/pkg/errors"
)

// 迁移表
func Migrate() {
	GlobalLoger.Info("start migrate ......")
	err := GlobalDB.AutoMigrate()
	utils.CheckErr(errors.Wrap(err, "migrate error"))
	GlobalLoger.Info("migrate success !!!")
}
