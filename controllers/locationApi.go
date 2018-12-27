/*
@Time : 2018/12/27 下午1:26 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"strconv"
	"strings"
	"anonymousFriends/util"
)

//获取随机地址
func GetRandomLocation() (provinceName, cityName, areaName string, longitude, latitude float64) {
	//随机省
	var province models.Location
	randomProvinceSql := "select * from location WHERE level=1 ORDER BY RAND() limit 1"
	base.DBEngine.SQL(randomProvinceSql).Get(&province)

	//随机市
	var city models.Location
	randomCitySql := "select * from location WHERE level=2 and parent_id='"+strconv.FormatInt(province.AreaId, 10)+"' ORDER BY RAND() limit 1"
	base.DBEngine.SQL(randomCitySql).Get(&city)

	//随机区
	var area models.Location
	randomAreaSql := "select * from location WHERE level=3 and parent_id='"+strconv.FormatInt(city.AreaId, 10)+"' ORDER BY RAND() limit 1"
	_, err := base.DBEngine.SQL(randomAreaSql).Get(&area)
	if err != nil {
		util.Logger.Info(err.Error())
	}

	//所在区的移动坐标
	centerLongitudeStr := "0"
	centerLatitudeStr := "0"
	if area.Center != "" {
		centerLongitudeStr = strings.Split(area.Center, ",")[0]
		centerLatitudeStr = strings.Split(area.Center, ",")[1]
	} else if city.Center != "" {
		centerLongitudeStr = strings.Split(city.Center, ",")[0]
		centerLatitudeStr = strings.Split(city.Center, ",")[1]
	} else if province.Center != "" {
		centerLongitudeStr = strings.Split(province.Center, ",")[0]
		centerLatitudeStr = strings.Split(province.Center, ",")[1]
	}
	centerLongitude, err := strconv.ParseFloat(centerLongitudeStr, 64)
	if err != nil {
		util.Logger.Info(area.AreaName + "  strconv.ParseFloat(centerLongitudeStr, 64) err:"+err.Error())
	}
	centerLatitude, err := strconv.ParseFloat(centerLatitudeStr, 64)
	if err != nil {
		util.Logger.Info(area.AreaName + "  strconv.ParseFloat(centerLatitudeStr, 64) err:"+err.Error())
	}
	longitude, latitude = CalcZombiePositionByRangeMeter(centerLongitude, centerLatitude, 500)

	return province.AreaName, city.AreaName, area.AreaName, longitude, latitude
}
