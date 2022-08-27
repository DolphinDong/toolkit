package global

import (
	"github.com/DolphinDong/toolkit/moni-server/config"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	GlobalConfig    *config.Server
	GlobalLoger     *logrus.Logger
	GlobalCronLoger *logrus.Logger
	GlobalDB        *gorm.DB
	GlobalRedisPool *redis.Pool
)
