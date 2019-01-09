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
func GetRandomLocation(provinceId, cityId, areaId int64) (provinceName, cityName, areaName string, longitude, latitude float64) {
	//随机省
	var province models.Location
	if provinceId == 0 {
		randomProvinceSql := "select * from location WHERE level=1 ORDER BY RAND() limit 1"
		base.DBEngine.SQL(randomProvinceSql).Get(&province)
	} else {
		base.DBEngine.Table("location").Where("area_id=?", provinceId).Get(&province)
	}

	//随机市
	var city models.Location
	if cityId == 0 {
		randomCitySql := "select * from location WHERE level=2 and parent_id='"+strconv.FormatInt(province.AreaId, 10)+"' ORDER BY RAND() limit 1"
		base.DBEngine.SQL(randomCitySql).Get(&city)
	} else {
		base.DBEngine.Table("location").Where("area_id=?", cityId).Get(&city)
	}


	//随机区
	var area models.Location
	if areaId == 0 {
		randomAreaSql := "select * from location WHERE level=3 and parent_id='" + strconv.FormatInt(city.AreaId, 10) + "' ORDER BY RAND() limit 1"
		base.DBEngine.SQL(randomAreaSql).Get(&area)
	} else {
		base.DBEngine.Table("location").Where("area_id=?", areaId).Get(&area)
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
	longitude, latitude = CalcZombiePositionByRangeMeter(centerLongitude, centerLatitude, 0, 0, 0, 0, 500)

	return province.AreaName, city.AreaName, area.AreaName, longitude, latitude
}

//根据所在城市区域获得省地址
func GetProvinceLocation(province string) (location models.Location) {
	sql := "select * from location province where province.area_name='"+province+"'"
	base.DBEngine.SQL(sql).Get(&location)
	return location
}

//根据所在城市区域获得市地址
func GetCityLocation(city string) (location models.Location) {
	sql := "select * from location city where city.area_name='"+city+"'"
	base.DBEngine.SQL(sql).Get(&location)
	return location
}

//根据所在城市区域获得区地址
func GetAreaLocation(city, area string) (location models.Location) {
	sql := "select * from location area where area.area_name='"+area+"' and exists(select 1 from location city where city.area_id=area.parent_id and city.area_name='"+city+"')"
	base.DBEngine.SQL(sql).Get(&location)
	return location
}