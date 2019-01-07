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
	c.AddFunc("@every 20s", func() {
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

	//每天扔僵尸漂流瓶
	c.AddFunc("@every 1m", func() {
		throwZombieDriftBottle()
	})

	//每天检查特殊地址加僵尸账户
	c.AddFunc("@every 24h", func() {
		createZombieAtSpecialLocation()
	})

	//每1分钟查看需要关闭的共享地理位置
	c.AddFunc("@every 1m", func() {
		checkUnclosedSharePositionGroup()
	})

	//c.AddFunc("0 54 16 * * ", func() {
	//	createZombieAtSpecialLocation()
	//})

	//c.AddFunc("@every 1s", func() {
		//birthday := controllers.GetRandomBirthday()
		//util.Logger.Info(birthday)

	//	provinceName, cityName, areaName, longitude, latitude := controllers.GetRandomLocation(810, 849, 0)
	//	util.Logger.Info(provinceName+" "+cityName+" "+areaName+" "+strconv.FormatFloat(longitude, 'f', 6, 64)+" "+strconv.FormatFloat(latitude, 'f', 6, 64))
	//})


	c.Start()
}
