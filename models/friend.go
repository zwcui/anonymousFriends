/*
@Time : 2018/12/21 下午4:26 
@Author : zwcui
@Software: GoLand
*/
package models

//好友请求
type FriendRequest struct {
	Id       					int64			`description:"id" json:"id" xorm:"pk autoincr"`
	SenderUid		       		int64			`description:"发送人" json:"senderUid"`
	ReceiverUid		       		int64			`description:"接收人" json:"receiverUid"`
	Status			       		int 			`description:"状态，0未处理请求，1已接收请求，2已拒绝请求" json:"status"`
	RejectReason	       		int 			`description:"拒绝原因" json:"rejectReason"`
	Created           			int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//好友
type Friend struct {
	Id       					int64			`description:"id" json:"id" xorm:"pk autoincr"`
	OwnerUid		       		int64			`description:"主人" json:"ownerUid"`
	FriendUid		       		int64			`description:"朋友" json:"friendUid"`
	Status			       		int 			`description:"状态，1正常，2删除好友，3加入黑名单" json:"status"`
	Created           			int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}






//----------------结构体-----------------



type FriendListContainer struct {
	BaseListContainer
	FriendList 					[]UserShort 	`description:"好友列表" json:"friendList"`
}





