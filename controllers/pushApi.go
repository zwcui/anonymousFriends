/*
@Time : 2018/12/21 下午4:52 
@Author : zwcui
@Software: GoLand
*/
package controllers

import "anonymousFriends/models"

//推送模块
type PushController struct {
	apiController
}

func (this *PushController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
}

// @Title 普通推送
// @Description 普通推送
// @Param	uId				formData		int64	  		true		"uId"
// @Param	content			formData		string  		true		"内容"
// @Success 200 {string} success
// @router /pushCommonMessage [post]
func (this *CommentController) PushCommonMessage() {
	uId := this.MustInt64("uId")
	content := this.MustString("content")

	var message models.Message
	message.Content = content
	message.SenderUid = 0
	message.ReceiverUid = uId
	message.Type = 1
	PushCommonMessageToUser(uId, &message, "", 0, "")

	this.ReturnData = "success"
}

// @Title socket推送
// @Description socket推送
// @Param	uId				formData		int64	  		true		"uId"
// @Param	content			formData		string  		true		"内容"
// @Success 200 {string} success
// @router /pushSocketMessage [post]
func (this *CommentController) PushSocketMessage() {
	uId := this.MustInt64("uId")
	content := this.MustString("content")

	var message models.Message
	message.Content = content
	message.SenderUid = 0
	message.ReceiverUid = uId
	message.Type = 1
	PushSocketMessageToUser(uId, &message, "", 0, "", 3)

	this.ReturnData = "success"
}











