package models

//存入MongoDB的聊天记录
//type MongoDBMessage struct {
//	GroupId				int64			`description:"聊天组id" json:"groupId" `
//	SenderUid			int64			`description:"发送人id" json:"senderUid" `
//	ReceiverUid			int64			`description:"接受人id" json:"receiverUid" `
//	ContentType			int				`description:"聊天类型，1为文本，2为语音" json:"contentType" `
//	Content				string			`description:"聊天内容" json:"content" `
//	Created           	int64  			`description:"消息时间" json:"created"`
//	ImageWidth     		string 				`description:"图片宽度,客户端根据这个显示图片宽度" json:"imageWidth"`
//	ImageHeight     	string 				`description:"图片高度,客户端根据这个显示图片高度" json:"imageHeight"`
//}

//聊天组
type Group struct {
	GroupId				int64			`description:"聊天组id" json:"groupId" xorm:"pk autoincr"`
	GroupType			int				`description:"聊天组类型 1:一对一 2:一对多 " json:"groupType" `
	GroupName  			string			`description:"聊天组名称" json:"groupName"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//聊天组成员记录
type Member struct {
	MemberId			int64			`description:"memberId" json:"memberId" xorm:"pk autoincr"`
	GroupId				int64			`description:"聊天组id" json:"groupId"`
	UId       			int64			`description:"uId" json:"uId"`
	UnReadNum   		int64  			`description:"消息未读数量" json:"unReadNum" `
	MessageDeleteUpdated     int64  `description:"消息删除时间，用于查询已经删除的消息" json:"messageDeleteUpdated"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//-----------------结构体如下----------------------
type GroupInfo struct {
	Group				Group			`description:"聊天组" json:"group" `
}