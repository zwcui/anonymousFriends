package task

import (
	"time"
	"github.com/robfig/cron"
)

//初始化定时任务
func init() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.NewWithLocation(location)

	//其他参考cron表达式
	c.AddFunc("@every 5m", func() {
		TestTimedTask()
	})

	//每分钟修改僵尸账户位置
	c.AddFunc("@every 1m", func() {
		ZombieMoveTask()
	})

	c.Start()
}
