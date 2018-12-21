/*
@Time : 2018/12/21 下午4:33 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"anonymousFriends/util"
)

//好友模块
type FriendController struct {
	apiController
}

func (this *FriendController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
}

// @Title 发送好友请求
// @Description 发送好友请求
// @Param	senderUid				formData		int64  		true		"发送人id"
// @Param	receiverUid				formData		int64  		true		"接收人id"
// @Success 200 {string} success
// @router /makeFriends [post]
func (this *FriendController) MakeFriends() {
	senderUid := this.MustInt64("senderUid")
	receiverUid := this.MustInt64("receiverUid")

	var storedSenderRequest models.Friend
	hasSenderRequest, _ := base.DBEngine.Table("friend").Where("sender_uid=?", senderUid).And("receiver_uid=?", receiverUid).Get(&storedSenderRequest)
	if hasSenderRequest {
		if storedSenderRequest.Status == 0 {
			this.ReturnData = util.GenerateAlertMessage(models.FriendError100)
			return
		} else if storedSenderRequest.Status == 1 {
			this.ReturnData = util.GenerateAlertMessage(models.FriendError200)
			return
		}
	} else {
		var storedReceiverRequest models.Friend
		hasReceiverRequest, _ := base.DBEngine.Table("friend").Where("sender_uid=?", receiverUid).And("receiver_uid=?", senderUid).Get(&storedReceiverRequest)
		if hasReceiverRequest {
			if storedReceiverRequest.Status == 0 {
				this.ReturnData = util.GenerateAlertMessage(models.FriendError300)
				return
			} else if storedReceiverRequest.Status == 1 {
				this.ReturnData = util.GenerateAlertMessage(models.FriendError200)
				return
			}
		}
	}

	var friend models.Friend
	friend.SenderUid = senderUid
	friend.ReceiverUid = receiverUid
	friend.Status = 0
	base.DBEngine.Table("friend").InsertOne(&friend)

	//推送
	PushCommonMessageToUser()

	this.ReturnData = "success"
}

// @Title 处理好友请求
// @Description 处理好友请求
// @Param	id						formData		int64  		true		"申请id"
// @Param	result					formData		int  		true		"处理结果，1接收 2拒绝"
// @Success 200 {string} success
// @router /handleMakeFriendsRequest [patch]
func (this *FriendController) HandleMakeFriendsRequest() {
	id := this.MustInt64("id")
	result := this.MustInt("result")

	var friend models.Friend
	base.DBEngine.Table("friend").Where("id=?", id).Get(&friend)
	if friend.Status == 1 || friend.Status == 2 {
		this.ReturnData = util.GenerateAlertMessage(models.FriendError400)
		return
	}

	if result == 1 {
		friend.Status = 1
	} else if result == 2 {
		friend.Status = 2
	}
	base.DBEngine.Table("friend").Where("id=?", id).Cols("status").Update(&friend)


	//推送
	if result == 1 {
		PushCommonMessageToUser()
	} else if result == 2 {
		PushCommonMessageToUser()
	}

	this.ReturnData = "success"
}