/*
@Time : 2018/12/24 下午6:10 
@Author : zwcui
@Software: GoLand
*/
package task

import (
	"anonymousFriends/base"
	"anonymousFriends/models"
	"anonymousFriends/util"
)

const AutoZombieSocialDynamicsNUmber = 5

//定时任务增加僵尸账户的朋友圈
func addZombieSocialDynamics(){
	util.Logger.Info("定时任务：增加僵尸账户的朋友圈")
	var zombieSocialDynamicList []models.ZombieSocialDynamics
	base.DBEngine.Table("zombie_social_dynamics").Where("status=0").Desc("created").Limit(AutoZombieSocialDynamicsNUmber, 0).Find(&zombieSocialDynamicList)
	if zombieSocialDynamicList != nil {
		zombieSocialDynamicList = make([]models.ZombieSocialDynamics, 0)
	}

	if len(zombieSocialDynamicList) > 0 {
		var zombieList []models.User
		base.DBEngine.Table("user").Where("is_zombie=1").Asc("updated").Limit(len(zombieSocialDynamicList), 0).Find(&zombieList)

		for index, zombie := range zombieList {
			var zombieSocialDynamics models.ZombieSocialDynamics
			if index >= len(zombieSocialDynamicList) {
				break
			}
			zombieSocialDynamics = zombieSocialDynamicList[index]

			var socialDynamics models.SocialDynamics
			socialDynamics.UId = zombie.UId
			socialDynamics.Content = zombieSocialDynamics.Content
			socialDynamics.Picture = zombieSocialDynamics.Picture
			socialDynamics.LikeNum = util.GenerateRangeNum(0, 20)
			if zombieSocialDynamics.Province != "" {
				socialDynamics.Province = zombieSocialDynamics.Province
			} else {
				socialDynamics.Province = zombie.Province
			}
			if zombieSocialDynamics.City != "" {
				socialDynamics.City = zombieSocialDynamics.City
			} else {
				socialDynamics.City = zombie.City
			}
			if zombieSocialDynamics.Area != "" {
				socialDynamics.Area = zombieSocialDynamics.Area
			} else {
				socialDynamics.Area = zombie.Area
			}
			if zombieSocialDynamics.Longitude != 0 {
				socialDynamics.Longitude = zombieSocialDynamics.Longitude
			} else {
				socialDynamics.Longitude = zombie.Longitude
			}
			if zombieSocialDynamics.Latitude != 0 {
				socialDynamics.Latitude = zombieSocialDynamics.Latitude
			} else {
				socialDynamics.Latitude = zombie.Latitude
			}
			base.DBEngine.Table("social_dynamics").InsertOne(&socialDynamics)

			//用过的备注
			zombieSocialDynamics.Status = 1
			base.DBEngine.Table("zombie_social_dynamics").Where("id=?", zombieSocialDynamics.Id).Cols("status").Update(&zombieSocialDynamics)
		}
	}
}
