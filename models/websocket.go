package models

import (
	"golang.org/x/net/websocket"
)

//socket统一结构
type SocketMessage struct {
	MessageType				int					`description:"消息类型，-1为后台建立连接，0为前台建立连接，1为普通聊天，2为被挤下线，3普通推送，4刷新标识，5客户端自定义" json:"messageType" `
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
	Province        		string				`description:"省" json:"province"`
	City	        		string				`description:"市" json:"city"`
	Area	        		string				`description:"区" json:"area"`
	Longitude				float64				`description:"经度" json:"longitude"`
	Latitude				float64				`description:"纬度" json:"latitude"`
	ChangeFlag				int					`description:"是否需要变更，1是0否" json:"changeFlag"`
}

//聊天结构体
type UserChatSocketMessage struct {
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
	ImageWidth     			string 				`description:"图片宽度,客户端根据这个显示图片宽度" json:"imageWidth"`
	ImageHeight     		string 				`description:"图片高度,客户端根据这个显示图片高度" json:"imageHeight"`
}

//刷新结构体
type RefreshSocketMessage struct {
	Position				int    				`description:"刷新位置 1首页 " json:"position"`
}


//--------------------

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



