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

	//每12小时增加僵尸账户的朋友圈
	c.AddFunc("@every 12h", func() {
		addZombieSocialDynamics()
	})

	//每24小时检查是否有过期的漂流瓶，设置为已失效
	c.AddFunc("@every 24h", func() {
		checkDriftBottleExpiryTime()
	})

	//每1分钟检查是否有过期的漂流瓶，设置为已失效
	c.AddFunc("@every 1m", func() {
		throwZombieDriftBottle()
	})

	//每10分钟增加僵尸账户，位置随机
	//c.AddFunc("0 0,36 17 * * ", func() {
	//	createSuZhouZombie()
	//})

	//c.AddFunc("@every 1s", func() {
	//	provinceName, cityName, areaName, longitude, latitude := controllers.GetRandomLocation(810, 849, 0)
	//	util.Logger.Info(provinceName+" "+cityName+" "+areaName+" "+strconv.FormatFloat(longitude, 'f', 6, 64)+" "+strconv.FormatFloat(latitude, 'f', 6, 64))
	//})


	c.Start()
}
