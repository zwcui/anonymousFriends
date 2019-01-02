/*
@Time : 2018/12/29 下午2:01 
@Author : zwcui
@Software: GoLand
*/
package models

//共享地理位置组
type SharePositionGroup struct {
	Id	       					int64  				`description:"记录id" json:"id" xorm:"pk autoincr"`
	GroupId    					int64  				`description:"聊天组id" json:"groupId"`
	Originator    				int64  				`description:"发起人" json:"originator"`
	Status    					int  				`description:"状态，0等待接收，1进行中，2已关闭" json:"status"`
	Remark    					string 				`description:"备注" json:"remark"`
	Created           			int64  				`description:"创建时间" json:"created" xorm:"created"`
	Updated           			int64  				`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  				`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//共享地理位置成员
type SharePositionMember struct {
	Id	       					int64  				`description:"记录id" json:"id" xorm:"pk autoincr"`
	SharePositionGroupId		int64  				`description:"共享地理位置组id" json:"sharePositionGroupId"`
	UId    						int64  				`description:"用户id" json:"uId"`
	Status    					int  				`description:"状态，1接受 2拒绝" json:"status"`
	Created           			int64  				`description:"创建时间" json:"created" xorm:"created"`
	Updated           			int64  				`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  				`description:"删除时间" json:"deleted" xorm:"deleted"`
}


//----------------------结构体如下------------------------------

type SharePositionInfo struct {
	SharePositionGroup								`description:"位置共享" xorm:"extends"`
	SenderNickName				string				`description:"申请人昵称" json:"senderNickName"`
	SenderAvatar				string				`description:"申请人头像" json:"senderAvatar"`
}

type SharePositionInfoContainer struct {
	SharePositionInfo			SharePositionInfo	`description:"登录用户信息" json:"sharePositionInfo"`
}

