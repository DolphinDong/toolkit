package initialize

import (
	"fmt"
	. "github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/utils"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)

const (
	defaultMaxIdle        = 5
	defaultMaxActive      = 100
	defaultConnectTimeout = 10  //秒
	defaultIdleTimeout    = 100 //秒
)

func initRedis() {
	GlobalLoger.Info("start init redis pool.....")
	maxIdle := GlobalConfig.Redis.PoolMaxIdle
	if maxIdle == 0 {
		maxIdle = defaultMaxIdle
	}

	maxActive := GlobalConfig.Redis.PoolMaxActive
	if maxActive == 0 {
		maxActive = defaultMaxActive
	}
	if GlobalConfig.Redis.ConnectTimeOut == 0 {
		GlobalConfig.Redis.ConnectTimeOut = defaultConnectTimeout
	}

	if GlobalConfig.Redis.IdleTimeout == 0 {
		GlobalConfig.Redis.IdleTimeout = defaultIdleTimeout
	}
	address := fmt.Sprintf("%v:%v", GlobalConfig.Redis.RedisHost, GlobalConfig.Redis.RedisPort)
	GlobalRedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Second * time.Duration(GlobalConfig.Redis.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				address,
				redis.DialPassword(GlobalConfig.Redis.RedisPassword),
				redis.DialDatabase(GlobalConfig.Redis.RedisDB),
				redis.DialConnectTimeout(time.Second*time.Duration(GlobalConfig.Redis.ConnectTimeOut)))
		},
	}
	_, err := GlobalRedisPool.Get().Do("ping")
	utils.CheckErr(errors.Wrap(err, "init redis pool error"))
	GlobalLoger.Info("start init redis pool success!!")
}
