package cron

import (
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"log"
)

type CronJob struct {
	taskFunc     map[string]func()
	taskSchedule map[string]string
	cron         *cron.Cron
}

func (c *CronJob) Run() error {
	c.cron = cron.New(cron.WithSeconds())
	for key, fuc := range c.taskFunc {
		if schedule, exist := c.taskSchedule[key]; exist {
			_, err := c.cron.AddFunc(schedule, fuc)
			if err != nil {
				return errors.WithStack(err)
			}
			log.Printf("add task %v ---> %v", key, schedule)
		}
	}
	c.cron.Start()
	select {}
}

func NewCronJob(taskSchedule map[string]string) *CronJob {
	taskFunc := getTaskFunc()
	return &CronJob{
		taskFunc:     taskFunc,
		taskSchedule: taskSchedule,
	}
}
