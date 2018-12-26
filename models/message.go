package models

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

//系统消息、推送
type Message struct {
	MId         int64  `description:"消息id" json:"mId" xorm:"pk autoincr"`
	SenderUid   int64  `description:"发送者uid 0则是系统发送的" json:"senderUid" xorm:"index(msgSenderUid)"`
	ReceiverUid int64  `description:"接收者uid 为0则是推送给所有人" json:"receiverUid" xorm:"index(msgReceiverUid)"`
	ActionUrl   string `description:"跳转url" json:"actionUrl" xorm:"text"`
	Content     string `description:"内容" json:"content"`
	Detail      string `description:"详情" json:"detail"`
	Type        int    `description:"消息类型 1:好友请求消息 2:评论回复消息 3:漂流瓶回复消息 " json:"type" xorm:"index(msgTypeUid)"`
	Created     int64  `description:"创建时间" json:"created" xorm:"created"`
	DeletedAt   int64  `description:"删除时间" json:"-" xorm:"deleted"`
}

//-----------------结构体如下----------------------
type GroupInfo struct {
	Group				Group			`description:"聊天组" json:"group" `
}

type PushSocketMessage struct {
	Message				`description:"消息" xorm:"extends"`
	Sound		string	`description:"声音" json:"sound"`
}

//-------------------推送消息文案----------------------
const (
	SendFriendRequest 		= "您收到一个好友请求，尽快处理哦~"
	AcceptFriendRequest 	= "对方已接收您的好友请求，快开始聊天吧~"
	RejectFriendRequest 	= "对方拒绝了您的好友请求"

	CommentOnSocialDynamics = "您收到一个朋友圈评论~"
	CommentOnDriftBottle	= "您收到一个漂流瓶回复~"
)


