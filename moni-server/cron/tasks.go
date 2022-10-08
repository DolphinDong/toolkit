package cron

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/DolphinDong/gotools/clients/wechat"
	"github.com/DolphinDong/toolkit/moni-server/common"
	"github.com/DolphinDong/toolkit/moni-server/global"
	"github.com/DolphinDong/toolkit/moni-server/models"
	"github.com/pkg/errors"
	"net"
	"time"
)

var (
	AlarmRecipient = []string{"liudong"}
	AlarmTemplate  = `【端口告警】
IP: %v
端口: %v
告警信息: %v
告警时间: %v
`
	AlarmRecoverTemplate = `【端口恢复】
IP: %v
端口: %v
告警信息: %v
恢复时间: %v
`
)

func getTaskFunc() (taskFunc map[string]func()) {
	taskFunc = make(map[string]func())
	taskFunc["onesbot-db-check"] = PortCheck
	return
}

//端口监控
func PortCheck() {
	var hostPorts []*models.HostPort
	messageTmp := "%v connect timeout, please check."
	result := global.GlobalDB.Find(&hostPorts)
	if result.Error != nil {
		global.GlobalCronLoger.Errorf("select host ports error: %v", result.Error)
		return
	}
	for _, hostPort := range hostPorts {
		go func(hp *models.HostPort) {
			message := fmt.Sprintf(messageTmp, hp.Name)
			alarmIdentifier := common.AlarmIdentifier{
				ServerHost:   hp.Host,
				ServerPort:   hp.Port,
				AlarmMessage: message,
			}
			data, err := json.Marshal(alarmIdentifier)
			if err != nil {
				global.GlobalCronLoger.Errorf("Marshal failed:%+v", errors.WithStack(err))
				return
			}
			// 生成MD5
			b := md5.Sum(data)
			m := fmt.Sprintf("moni%x", b)

			addr := fmt.Sprintf("%v:%v", hp.Host, hp.Port)
			conn, err := net.DialTimeout("tcp", addr, time.Second*15)
			// 连接失败
			if err != nil {
				global.GlobalCronLoger.Errorf("%v connect failed", addr)
				redisConnect := global.GlobalRedisPool.Get()
				defer redisConnect.Close()
				result, err := redisConnect.Do("get", m)
				if err != nil {
					global.GlobalCronLoger.Errorf("get redis value failed: %+v", errors.WithStack(err))
					return
				}
				// 告警不存在则直接发送消息到机器人
				if result == nil {
					redisConnect.Do("SETEX", m, 2*60*60, message)
					msg := wechat.RobotMessage{
						MsgType: "text",
					}
					msg.Text.MentionedList = AlarmRecipient
					msg.Text.Content = fmt.Sprintf(AlarmTemplate, hp.Host, hp.Port, message, time.Now().Format("2006-01-02 15:04:05"))
					err = wechat.SendRobotMessage(global.GlobalConfig.RobotKey, global.GlobalConfig.WechatBaseUrl, msg)
					if err != nil {
						global.GlobalCronLoger.Errorf("send robot message failed:%+v", errors.WithStack(err))
						return
					}
				}
				return
			} else { // 连接成功
				redisConnect := global.GlobalRedisPool.Get()
				defer redisConnect.Close()
				result, err := redisConnect.Do("get", m)
				if err != nil {
					global.GlobalCronLoger.Errorf("get redis value failed: %+v", errors.WithStack(err))
					return
				}
				// redis中有告警记录，说明正在告警，需要将告警信息删除
				if result != nil {
					_, err = redisConnect.Do("del", m)
					if err != nil {
						global.GlobalCronLoger.Errorf("delete redis key failed: %+v", errors.WithStack(err))
						return
					}
					msg := wechat.RobotMessage{
						MsgType: "text",
					}
					msg.Text.MentionedList = AlarmRecipient
					message := fmt.Sprintf("%v 端口已恢复~~~", hp.Name)
					msg.Text.Content = fmt.Sprintf(AlarmRecoverTemplate, hp.Host, hp.Port, message, time.Now().Format("2006-01-02 15:04:05"))
					err = wechat.SendRobotMessage(global.GlobalConfig.RobotKey, global.GlobalConfig.WechatBaseUrl, msg)
					if err != nil {
						global.GlobalCronLoger.Errorf("send robot message failed:%+v", errors.WithStack(err))
						return
					}
				}
			}
			defer conn.Close()
		}(hostPort)
	}
}
