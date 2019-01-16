/*
@Time : 2019/1/16 下午5:10 
@Author : zwcui
@Software: GoLand
*/
package models


//每日签到
type DailySignIn struct {
	Id       					int64			`description:"id" json:"id" xorm:"pk autoincr"`
	UId				       		int64			`description:"签到人" json:"uId"`
	SignInDate		       		string			`description:"签到日期" json:"signInDate"`
	Status		       			int				`description:"签到状态，1为正常" json:"status"`
	Created           			int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}
