/*
@Time : 2018/12/18 上午10:31
@Author : zwcui
@Software: GoLand
*/
package models

const (
	ErrorSpliter = "###"
)

const (
	CommonError100 		= "common100###服务器出错"

	UserError100 		= "User100###昵称重复，请重新输入"
	UserError101 		= "User101###密码加密失败"
	UserError102 		= "User102###未找到该用户，请检查昵称是否正确"
	UserError103 		= "User103###密码错误，请重新输入"
	UserError104 		= "User103###密码长度需在6-30之间"

	WebsocketError100 	= "Websocket100###您的账户已在其他地方登陆"

	FriendError100 		= "Friend100###您已发起好友申请，稍等对方接收哦~"
	FriendError200 		= "Friend200###你们已经成为好友了哦~"
	FriendError300 		= "Friend300###对方也向您发起了好友申请，请先接收哦~"
	FriendError400 		= "Friend400###您已处理该条申请哦~"

)


