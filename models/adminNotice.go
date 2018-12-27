/*
@Time : 2018/12/27 下午4:16 
@Author : zwcui
@Software: GoLand
*/
package models

//管理员通知，存储管理员需做的事情
type AdminNotice struct {
	Id		           		int64  				`description:"Id" json:"id" xorm:"pk autoincr"`
	Type					int 				`description:"类型，1用户与僵尸账户聊天 2有用户注册" json:"type"`
	TypeId					int64 				`description:"类型id" json:"typeId"`
	Content	           		string 				`description:"内容" json:"content"`
	Status	           		int 				`description:"状态，0未处理，1已处理，2已忽略" json:"status"`
	Created           		int64  				`description:"注册时间" json:"created" xorm:"created"`
	Updated           		int64  				`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         		int64  				`description:"删除时间" json:"deleted" xorm:"deleted"`
}


//-------------------------结构体如下-------------------------------------

type AdminNoticeListContainer struct {
	BaseListContainer
	AdminNoticeList 		[]AdminNotice 		`description:"管理员通知列表" json:"adminNoticeList"`
}

