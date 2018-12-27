package task

import (
	"strconv"
	"anonymousFriends/util"
	"anonymousFriends/models"
	"anonymousFriends/base"
	"math/rand"
	"anonymousFriends/controllers"
)

func TestTimedTask(){
	util.Logger.Info(strconv.FormatInt(util.UnixOfBeijingTime(), 10)+"-->"+util.FormatTimestamp(util.UnixOfBeijingTime()))
}

//僵尸账户定时任务主动生成总数
const ZombieLimit = 200

//僵尸账户移动
//选取未更新的僵尸账户中前50条
//按比例（1：1）选取进行更新位置
func ZombieMoveTask(){
	util.Logger.Info("定时任务：僵尸账户位置移动")
	var zombieList []models.UserShort
	base.DBEngine.Table("user").Where("is_zombie=1").And("status=1").And("longitude !=0 and latitude != 0").Asc("updated").Limit(50, 0).Find(&zombieList)

	for _, zombie := range zombieList {
		moveFlag := getRandomZombieMoveFlag()
		if moveFlag == 1 {
			//util.Logger.Info("定时任务：僵尸账户位置移动前-->"+strconv.FormatInt(zombie.UId, 10)+"："+strconv.FormatFloat(zombie.Longitude, 'f', 6, 64)+", "+strconv.FormatFloat(zombie.Latitude, 'f', 6, 64))
			zombie.Longitude, zombie.Latitude = controllers.CalcZombiePositionByRangeMeter(zombie.Longitude, zombie.Latitude, 50)
			base.DBEngine.Table("user").Where("u_id=?", zombie.UId).Cols("longitude", "latitude").Update(&zombie)
			//util.Logger.Info("定时任务：僵尸账户位置移动后-->"+strconv.FormatInt(zombie.UId, 10)+"："+strconv.FormatFloat(zombie.Longitude, 'f', 6, 64)+", "+strconv.FormatFloat(zombie.Latitude, 'f', 6, 64))
		}
	}

}

//获得随机经纬度加减
func getRandomZombieMoveFlag() int {
	sIndex := rand.Intn(len(models.ZombieMoveFlagRatio))
	return models.ZombieMoveFlagRatio[sIndex]
}

//增加僵尸账户，位置随机
func createZombie(){
	util.Logger.Info("定时任务：增加僵尸账户，位置随机")
	zombieTotal, _ := base.DBEngine.Table("user").Where("is_zombie=1").Count(new(models.User))
	if zombieTotal < ZombieLimit {
		for i:=0;i<int(ZombieLimit-zombieTotal);i++{
			var zombie models.User
			zombie.NickName = controllers.GetDefaultNickName()
			zombie.Gender, zombie.Avatar = controllers.GetRandomGenderAndAvatar()
			hashedPassword, salt, _ := util.EncryptPassword("iamzombie")
			zombie.Password = hashedPassword
			zombie.Salt = salt
			zombie.Birthday = controllers.GetRandomBirthday()
			zombie.Status = 1
			zombie.IsZombie = 1
			zombie.Province, zombie.City, zombie.Area, zombie.Longitude, zombie.Latitude = controllers.GetRandomLocation(0, 0, 0)
			base.DBEngine.Table("user").InsertOne(&zombie)
		}
	}
}

func createSuZhouZombie(){
	util.Logger.Info("定时任务：增加僵尸账户，位置随机 start")
	zombieTotal, _ := base.DBEngine.Table("user").Where("is_zombie=1").Count(new(models.User))
	if zombieTotal < ZombieLimit {
		for i:=0;i<int(ZombieLimit-zombieTotal);i++{
			var zombie models.User
			zombie.NickName = controllers.GetDefaultNickName()
			zombie.Gender, zombie.Avatar = controllers.GetRandomGenderAndAvatar()
			hashedPassword, salt, _ := util.EncryptPassword("iamzombie")
			zombie.Password = hashedPassword
			zombie.Salt = salt
			zombie.Birthday = controllers.GetRandomBirthday()
			zombie.Status = 1
			zombie.IsZombie = 1
			zombie.Province, zombie.City, zombie.Area, zombie.Longitude, zombie.Latitude = controllers.GetRandomLocation(810, 849, 0)
			base.DBEngine.Table("user").InsertOne(&zombie)
		}
	}
	util.Logger.Info("定时任务：增加僵尸账户，位置随机 finish")
}