/*
@Time : 2018/12/21 下午4:52 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/util"
)

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
func (this *PushController) PushCommonMessage() {
	uId := this.MustInt64("uId")
	content := this.MustString("content")

	var message models.Message
	message.Content = content
	message.SenderUid = 0
	message.ReceiverUid = uId
	message.Type = 1
	ok, err := PushCommonMessageToUser(uId, &message, "", 0, "push")
	if !ok {
		util.Logger.Info("普通推送 失败 err:")
		if err != nil {
			util.Logger.Info(err.Error())
		}
	}

	this.ReturnData = "success"
}

// @Title socket推送
// @Description socket推送
// @Param	uId				formData		int64	  		true		"uId"
// @Param	content			formData		string  		true		"内容"
// @Success 200 {string} success
// @router /pushSocketMessage [post]
func (this *PushController) PushSocketMessage() {
	uId := this.MustInt64("uId")
	content := this.MustString("content")

	var message models.Message
	message.Content = content
	message.SenderUid = 0
	message.ReceiverUid = uId
	message.Type = 1
	ok, err := PushSocketMessageToUser(uId, &message, "", 0, "", 3)
	if !ok {
		util.Logger.Info("普通推送 失败 err:")
		if err != nil {
			util.Logger.Info(err.Error())
		}
	}

	this.ReturnData = "success"
}











