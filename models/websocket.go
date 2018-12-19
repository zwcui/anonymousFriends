package models

import (
	"golang.org/x/net/websocket"
)

//socket统一结构
type SocketMessage struct {
	MessageType				int					`description:"消息类型，-1为后台建立连接，0为前台建立连接，1为普通聊天，2为被挤下线，3刷新标识，4客户端自定义" json:"messageType" `
	MessageSendTime			int64				`description:"消息发送时间" json:"messageSendTime" `
	MessageSenderUid		int64				`description:"消息发送uid" json:"messageSenderUid" `
	MessageReceiverUid		int64				`description:"消息接受uid" json:"messageReceiverUid" `
	MessageExpireTime		int64				`description:"心跳有效时间" json:"messageExpireTime" `
	MessageContent			string				`description:"消息内容，jsonString" json:"messageContent" `
	MessageSign				string				`description:"消息签名" json:"messageSign" `
	MessageToken			string				`description:"用户token" json:"messageToken" `
}

//心跳结构体，传输地理位置，返回周围人的位置信息
type HeartBeatSocketMessage struct {

}

//聊天结构体
type UserSocketMessage struct {
	FromNickName  			string 	 			`description:"fromNickName" json:"fromNickName" `
	FromUid       			int64 	 			`description:"fromUid" json:"fromUid" `
	ToNickName         		string 	 			`description:"toNickName" json:"toNickName" `
	ToUid         			int64 	 			`description:"toUid" json:"toUid" `
	GroupId           		int64	  			`description:"groupId" json:"groupId" `
	GroupType        		int    				`description:"组类型 0:一对一 1:一对多 2:系统消息（不存库，仅识别使用） 3:客服一对一 " json:"groupType"`
	From					int	   				`description:"客服显示当前咨询者进入IM的入口，1.首页 2.首页列表（张三1231）3.次首页 4.次首页列表（张三1231）5.列表页（张三1231）6.详情页（张三1231）0.其他" json:"from" xorm:"notnull default 0"`
	Param					string 				`description:"咨询者进入IM的入口相关信息(json string)" json:"param" valid:"MaxSize(300)" `
	Content           		string  			`description:"content" json:"content" `
	ActionUrl         		string  			`description:"actionUrl" json:"actionUrl" `
	Type	           		int		  			`description:"消息内容类型 0:文本 1:图片 2:语音 3:视频" json:"type" `
	ImageWidth     			string 				`description:"图片宽度,客户端根据这个显示图片宽度" json:"imageWidth"`
	ImageHeight     		string 				`description:"图片高度,客户端根据这个显示图片高度" json:"imageHeight"`
}

//刷新结构体
type RefreshSocketMessage struct {
	Position				int    				`description:"刷新位置 1首页 " json:"position"`
}

//客户端定义直连
type BillSocketMessage struct {
	BillId 					int64				`description:"billId" json:"billId" `
	CallBack        		int    				`description:"是否为回调,0:否,1:是 （客服拨打都是回拨）" json:"callBack" `
	Badge        			int    				`description:"取消通知，1是取消" json:"badge" `
	Type 					int					`description:"类型" json:"type" `
	Content 				interface{}			`description:"内容" json:"content" `
}

//socket签名key
const SOCKET_MESSAGE_SIGN_KEY string = "anonymousfriends123socketmessage"

//直联消息缓存key
const SOCKET_UNSENT_MESSAGE = "socket_unsent_message"

//连接存储
type SocketConnection struct {
	Conn				*websocket.Conn			`description:"socket连接" json:"conn"`
	ConnType				int					`description:"socket连接类型，1前台，2后台" json:"connType"`
	ExpireTime				int64				`description:"socket连接有效截止时间" json:"expireTime"`
	Token					string				`description:"用户token" json:"token"`
}



