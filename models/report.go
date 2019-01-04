/*
@Time : 2019/1/4 下午4:37 
@Author : zwcui
@Software: GoLand
*/
package models

type Report struct {
	Id       					int64			`description:"id" json:"id" xorm:"pk autoincr"`
	SenderUid					int64 			`description:"发起人" json:"senderUid"`
	Type       					int 			`description:"类型，1:对app的建议 2:举报朋友圈 3:举报漂流瓶 4:举报用户 " json:"type"`
	TypeId     					int64 			`description:"类型对应的id" json:"typeId" `
	Content   					string 			`description:"举报内容" json:"content" `
	Created           			int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

