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

//特殊地址
type SpecialLocation struct {
	Id       				int64				`description:"id" json:"id" xorm:"pk autoincr"`
	LocationName			string				`description:"地址名称" json:"type"`
	Type					int					`description:"类型，1学校，2商场，3酒店" json:"type"`
	Province        		string				`description:"省" json:"province"`
	City	        		string				`description:"市" json:"city"`
	Area	        		string				`description:"区" json:"area"`
	Longitude				float64				`description:"经度" json:"longitude"`
	Latitude				float64				`description:"纬度" json:"latitude"`
	Status					int					`description:"状态，1正常，0不使用" json:"status"`
	RangeMemter				float64				`description:"范围，单位米" json:"rangeMemter"`
	Created           		int64  				`description:"注册时间" json:"created" xorm:"created"`
	Updated           		int64  				`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         		int64  				`description:"删除时间" json:"deleted" xorm:"deleted"`
}
