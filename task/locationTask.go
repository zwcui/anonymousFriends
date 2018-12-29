/*
@Time : 2018/12/28 下午3:29 
@Author : zwcui
@Software: GoLand
*/
package task

import (
	"anonymousFriends/base"
	"anonymousFriends/models"
	"anonymousFriends/controllers"
	"anonymousFriends/util"
	"github.com/satori/go.uuid"
)

//特殊地址加僵尸账户
func createZombieAtSpecialLocation(){
	util.Logger.Info("定时任务：特殊位置创建僵尸 start")
	var specialLocationList []models.SpecialLocation
	base.DBEngine.Table("special_location").Where("status=1").Find(&specialLocationList)
	if specialLocationList == nil {
		specialLocationList = make([]models.SpecialLocation, 0)
	}
	for _, specialLocation := range specialLocationList {
		longitudeMax := specialLocation.Longitude + specialLocation.RangeMemter / 1000000.0
		longitudeMin := specialLocation.Longitude - specialLocation.RangeMemter / 1000000.0
		latitudeMax := specialLocation.Latitude + specialLocation.RangeMemter / 1000000.0
		latitudeMin := specialLocation.Latitude - specialLocation.RangeMemter / 1000000.0


		userNum, _ := base.DBEngine.Table("user").Where("longitude >= ? and longitude <= ? and latitude >= ? and latitude <= ?", longitudeMin, longitudeMax, latitudeMin, latitudeMax).Count(new(models.User))
		if int(userNum) < specialLocation.UserNumber {
			for i:=0; i< specialLocation.UserNumber-int(userNum);i++ {
				var zombie models.User
				//zombie.NickName = controllers.GetDefaultNickName()
				randomUUId, _ := uuid.NewV4()
				zombie.NickName = "苏州大学(独墅湖校区南区) 学生"+randomUUId.String()
				zombie.Gender, zombie.Avatar = controllers.GetRandomGenderAndAvatar()
				hashedPassword, salt, _ := util.EncryptPassword("iamzombie")
				zombie.Password = hashedPassword
				zombie.Salt = salt
				zombie.Birthday = controllers.GetRandomBirthday()
				zombie.Status = 1
				zombie.IsZombie = 1
				zombie.Province = specialLocation.Province
				zombie.City = specialLocation.City
				zombie.Area = specialLocation.Area
				zombie.ZombieLongitudeMax = longitudeMax
				zombie.ZombieLongitudeMin = longitudeMin
				zombie.ZombieLatitudeMax = latitudeMax
				zombie.ZombieLatitudeMin = latitudeMin
				zombie.Longitude, zombie.Latitude = controllers.CalcZombiePositionByRangeMeter(specialLocation.Longitude, specialLocation.Latitude, zombie.ZombieLongitudeMax, zombie.ZombieLongitudeMin, zombie.ZombieLatitudeMax, zombie.ZombieLatitudeMin, int(specialLocation.RangeMemter))
				base.DBEngine.Table("user").InsertOne(&zombie)
			}
		}
	}
	util.Logger.Info("定时任务：特殊位置创建僵尸 finish")
}


