/*
@Time : 2018/12/24 下午6:46 
@Author : zwcui
@Software: GoLand
*/
package models

//评论
type Comment struct {
	Id       			int64			`description:"id" json:"id" xorm:"pk autoincr"`
	Type       			int 			`description:"类型，1是朋友圈评论" json:"type"`
	TypeId	 			int64 			`description:"类型id" json:"typeId"`
	ReplyCommentId		int64 			`description:"回复评论id" json:"replyCommentId"`
	Content    			string 			`description:"评论内容" json:"content"`
	SenderUid  			int64 			`description:"评论发送人" json:"senderUid"`
	ReceiverUid			int64 			`description:"评论接收人" json:"receiverUid"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}
