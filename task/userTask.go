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
			zombie.Longitude, zombie.Latitude = calcZombiePositionByTimedTask(zombie.Longitude, zombie.Latitude)
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

//计算僵尸每分钟移动位置
//50米的范围
//考虑不能去的位置，如河海
func calcZombiePositionByTimedTask(longitude float64, latitude float64) (float64, float64) {
	zombieLongitudeChange := float64(util.GenerateRangeNum(0, 50))/1000000.0 * controllers.GetRandomChange()
	zombieLatitudeChange := float64(util.GenerateRangeNum(0, 50))/1000000.0 * controllers.GetRandomChange()
	return longitude + zombieLongitudeChange, latitude + zombieLatitudeChange
}