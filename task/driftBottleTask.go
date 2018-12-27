/*
@Time : 2018/12/27 上午10:12 
@Author : zwcui
@Software: GoLand
*/
package task

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"anonymousFriends/util"
)

//用户漂流瓶少于这个阈值则新增僵尸漂流瓶
const UserDriftBottleNumberLimit = 10

//检查是否有过期的漂流瓶，设置为已失效
func checkDriftBottleExpiryTime(){
	util.Logger.Info("定时任务：检查是否有过期的漂流瓶，设置为已失效")
	var driftBottleList []models.DriftBottle
	base.DBEngine.Table("drift_bottle").Where("status != 3").And("expiry_time != 0").And("expiry_time < ?", util.UnixOfBeijingTime()).Find(&driftBottleList)
	for _, driftBottle := range driftBottleList {
		driftBottle.Status = 3
		driftBottle.Remark += "定时任务已失效;"
		base.DBEngine.Table("drift_bottle").Where("bottle_id=?", driftBottle.BottleId).Cols("status", "remark").Update(&driftBottle)
	}
}

//扔僵尸漂流瓶
func throwZombieDriftBottle(){
	util.Logger.Info("定时任务：扔僵尸漂流瓶")
	userDriftBottleNumber, _ := base.DBEngine.Table("drift_bottle").Where("status != 3").And("status != 2").Count(new(models.DriftBottle))
	if userDriftBottleNumber < UserDriftBottleNumberLimit {
		var zombieDriftBottleList []models.ZombieDriftBottle
		base.DBEngine.Table("zombie_drift_bottle").Where("status=0").Desc("created").Limit(int(UserDriftBottleNumberLimit - userDriftBottleNumber), 0).Find(&zombieDriftBottleList)
		if zombieDriftBottleList == nil {
			return
		}

		for _, zombieDriftBottle := range zombieDriftBottleList {
			//随机一个僵尸
			var zombie models.UserShort
			randomSql := "SELECT * FROM user WHERE is_zombie=1 and u_id >= ((SELECT MAX(u_id) FROM user where is_zombie=1)-(SELECT MIN(u_id) FROM user where is_zombie=1)) * RAND() + (SELECT MIN(u_id) FROM user where is_zombie=1)  LIMIT 1"
			base.DBEngine.SQL(randomSql).Get(&zombie)

			var driftBottle models.DriftBottle
			driftBottle.BottleType = zombieDriftBottle.BottleType
			driftBottle.BottleName = zombieDriftBottle.BottleName
			driftBottle.SenderUid = zombie.UId
			driftBottle.Content = zombieDriftBottle.Content
			driftBottle.Picture = zombieDriftBottle.Picture
			driftBottle.Position = zombieDriftBottle.Position
			driftBottle.Weather = zombieDriftBottle.Weather
			if zombieDriftBottle.Province != "" {
				driftBottle.Province = zombieDriftBottle.Province
			} else {
				driftBottle.Province = zombie.Province
			}
			if zombieDriftBottle.City != "" {
				driftBottle.City = zombieDriftBottle.City
			} else {
				driftBottle.City = zombie.City
			}
			if zombieDriftBottle.Area != "" {
				driftBottle.Area = zombieDriftBottle.Area
			} else {
				driftBottle.Area = zombie.Area
			}
			if zombieDriftBottle.Longitude != 0 {
				driftBottle.Longitude = zombieDriftBottle.Longitude
			} else {
				driftBottle.Longitude = zombie.Longitude
			}
			if zombieDriftBottle.Latitude != 0 {
				driftBottle.Latitude = zombieDriftBottle.Latitude
			} else {
				driftBottle.Latitude = zombie.Latitude
			}
			driftBottle.Status = 1

			if models.ExpiryTime != 0 {
				driftBottle.ExpiryTime = util.UnixOfBeijingTime() + models.ExpiryTime * 60 * 60
			}
			base.DBEngine.Table("drift_bottle").InsertOne(&driftBottle)

			zombieDriftBottle.Status = 1
			base.DBEngine.Table("zombie_drift_bottle").Where("bottle_id=?", zombieDriftBottle.BottleId).Cols("status").Update(&zombieDriftBottle)
		}
	}
}