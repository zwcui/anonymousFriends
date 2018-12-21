/*
@Time : 2018/12/21 下午2:10
@Author : zwcui
@Software: GoLand
*/
package models

//社交动态
type SocialDynamics struct {
	Id       			int64			`description:"id" json:"id" xorm:"pk autoincr"`
	UId       			int64			`description:"uId" json:"uId"`
	Content				string			`description:"文字内容" json:"content"`
	Picture				string			`description:"图片内容，多个英文逗号隔开" json:"picture"`
	LikeNum				int				`description:"点赞数" json:"likeNum"`
	Position        	string			`description:"位置名称，如全家创意产业园店" json:"province"`
	Province        	string			`description:"省" json:"province"`
	City	        	string			`description:"市" json:"city"`
	Area	        	string			`description:"区" json:"area"`
	Longitude			float64			`description:"经度" json:"longitude"`
	Latitude			float64			`description:"纬度" json:"latitude"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//点赞
type Like struct {
	Id       			int64			`description:"id" json:"id"`
	Type       			int 			`description:"type，1为朋友圈点赞" json:"type"`
	UId       			int64			`description:"uId" json:"uId"`
}


//--------------结构体-----------------

type SocialDynamicInfo struct {
	SocialDynamics								`description:"社交动态" xorm:"extends"`
	IsLike				int						`description:"是否点赞，1是0否" json:"isLike"`
}

type SocialDynamicListContainer struct {
	BaseListContainer
	SocialDynamicList 	[]SocialDynamicInfo 	`description:"社交动态列表" json:"socialDynamicList"`
}

