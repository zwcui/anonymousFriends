package models

import "net/url"

//聊天组
type Group struct {
	GroupId				int64			`description:"聊天组id" json:"groupId" xorm:"pk autoincr"`
	GroupType			int				`description:"聊天组类型 1:一对一 2:一对多 " json:"groupType" `
	GroupName  			string			`description:"聊天组名称" json:"groupName"`
	SenderUid          	int64  			`description:"注册时间" json:"senderUid"`
	ReceiverUid         int64  			`description:"注册时间" json:"receiverUid"`
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
	MId         		int64  			`description:"消息id" json:"mId" xorm:"pk autoincr"`
	SenderUid   		int64  			`description:"发送者uid 0则是系统发送的" json:"senderUid" xorm:"index(msgSenderUid)"`
	ReceiverUid 		int64  			`description:"接收者uid 为0则是推送给所有人" json:"receiverUid" xorm:"index(msgReceiverUid)"`
	ActionUrl   		string 			`description:"跳转url" json:"actionUrl" xorm:"text"`
	Content     		string 			`description:"内容" json:"content"`
	Detail      		string 			`description:"详情" json:"detail"`
	Type        		int    			`description:"消息类型 1:好友请求消息 2:评论回复消息 3:漂流瓶回复消息 4:共享位置请求消息 " json:"type" xorm:"index(msgTypeUid)"`
	Created     		int64  			`description:"创建时间" json:"created" xorm:"created"`
	DeletedAt   		int64  			`description:"删除时间" json:"-" xorm:"deleted"`
}

//聊天结构体，未读存入mysqlk
type UserUnsentChatMessage struct {
	Id       				int64				`description:"id" json:"id" xorm:"pk autoincr"`
	FromNickName  			string 	 			`description:"fromNickName" json:"fromNickName" `
	FromUid       			int64 	 			`description:"fromUid" json:"fromUid" `
	FromAvatar     			string 	 			`description:"fromAvatar" json:"fromAvatar" `
	ToNickName         		string 	 			`description:"toNickName" json:"toNickName" `
	ToUid         			int64 	 			`description:"toUid" json:"toUid" `
	ToAvatar     			string 	 			`description:"toAvatar" json:"toAvatar" `
	GroupId           		int64	  			`description:"groupId" json:"groupId" `
	GroupType        		int    				`description:"组类型 1:一对一 2:一对多 " json:"groupType"`
	Content           		string  			`description:"content" json:"content" `
	ContentType	           	int		  			`description:"消息内容类型 0:文本 1:图片 2:语音 3:视频 4:位置，经纬度，英文逗号,隔开" json:"contentType" `
	ImageWidth     			int 				`description:"图片宽度,客户端根据这个显示图片宽度" json:"imageWidth"`
	ImageHeight     		int 				`description:"图片高度,客户端根据这个显示图片高度" json:"imageHeight"`
	IsSent		     		int 				`description:"是否已发，1是0否" json:"isSent"`
	Created           		int64  				`description:"注册时间" json:"created" xorm:"created"`
	DeletedAt         		int64  				`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//-----------------结构体如下----------------------
type GroupInfo struct {
	Group				Group			`description:"聊天组" json:"group" `
}

type PushSocketMessage struct {
	Message				`description:"消息" xorm:"extends"`
	Sound		string	`description:"声音" json:"sound"`
}

type UserUnsentChatMessageListContainer struct {
	TotalCount 			int64 						`description:"总数" json:"totalCount"`
	MessageList			[]UserUnsentChatMessage		`description:"未读消息列表" json:"messageList" `
}

//-------------------推送消息文案----------------------
const (
	SendFriendRequest 		= "您收到一个好友请求，尽快处理哦~"
	AcceptFriendRequest 	= "对方已接收您的好友请求，快开始聊天吧~"
	RejectFriendRequest 	= "对方拒绝了您的好友请求"

	CommentOnSocialDynamics = "您收到一个朋友圈评论~"
	CommentOnDriftBottle	= "您收到一个漂流瓶回复~"

	SharePositionRequest	= "正在向您发起位置共享，请点击查看~"
)

//--------------------消息常量------------------------
//每次接口获取未读消息的条数
const UNSENT_MESSAGE_PAGE_NUM = 50

const (
	DEFAULT_SCHEME = "anonymousfriends"
	DEFAULT_HOST   = "app.anonymousfriends.cn"
)


//--------------------跳转jumpkey--------------------------
const (
	HOMEPAGE_JUMP_KEY                       = "homepage"                       //主页跳转key
)

//--------------------消息方法---------------------------
//生成通知跳转链接
//key为需要跳转的页面的key
//pramas为页面需要的参数
func JumpUrlWithKeyAndPramas(key string, pramas map[string]string) (urlStr string) {
	url := url.URL{}

	url.Scheme = DEFAULT_SCHEME
	url.Host = DEFAULT_HOST

	// Path
	url.Path = key

	if pramas != nil {
		// Query Parameters
		q := url.Query()
		for key, value := range pramas {
			q.Set(key, value)
		}
		url.RawQuery = q.Encode()
	}

	return url.String()
}