/*
@Time : 2018/12/27 下午1:21 
@Author : zwcui
@Software: GoLand
*/
package models

//地址
type Location struct {
	AreaId           		int64  				`description:"地区Id" json:"areaId" xorm:"pk"`
	AreaCode           		string  			`description:"地区编码" json:"areaCode"`
	AreaName           		string  			`description:"地区名" json:"areaName"`
	Level           		int		  			`description:"地区级别（1:省份province,2:市city,3:区县district,4:街道street）" json:"level"`
	CityCode           		string  			`description:"城市编码" json:"cityCode"`
	Center           		string  			`description:"城市中心点（即：经纬度坐标）" json:"center"`
	ParentId           		int64  				`description:"地区父节点" json:"parentId"`
}


