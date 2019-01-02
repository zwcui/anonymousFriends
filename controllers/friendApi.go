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
	"strconv"
)

//好友模块
type FriendController struct {
	apiController
}

func (this *FriendController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
		{"/makeFriends", "post"},
		{"/handleMakeFriendsRequest", "patch"},
		{"/getFriendList", "get"},
		{"/getFriendRequestList", "get"},
		{"/handleFriend", "patch"}}
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

	var storedSenderRequest models.FriendRequest
	hasSenderRequest, _ := base.DBEngine.Table("friend_request").Where("sender_uid=?", senderUid).And("receiver_uid=?", receiverUid).Get(&storedSenderRequest)
	if hasSenderRequest {
		if storedSenderRequest.Status == 0 {
			this.ReturnData = util.GenerateAlertMessage(models.FriendError100)
			return
		} else if storedSenderRequest.Status == 1 {
			this.ReturnData = util.GenerateAlertMessage(models.FriendError200)
			return
		}
	} else {
		var storedReceiverRequest models.FriendRequest
		hasReceiverRequest, _ := base.DBEngine.Table("friend_request").Where("sender_uid=?", receiverUid).And("receiver_uid=?", senderUid).Get(&storedReceiverRequest)
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

	var friendRequest models.FriendRequest
	friendRequest.SenderUid = senderUid
	friendRequest.ReceiverUid = receiverUid
	friendRequest.Status = 0
	base.DBEngine.Table("friend_request").InsertOne(&friendRequest)

	//推送
	var message models.Message
	message.Content = models.SendFriendRequest
	message.SenderUid = senderUid
	message.ReceiverUid = receiverUid
	message.Type = 1
	PushSocketMessageToUser(receiverUid, &message, "", 0, "", 3)

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

	var friendRequest models.FriendRequest
	base.DBEngine.Table("friend_request").Where("id=?", id).Get(&friendRequest)
	if friendRequest.Status == 1 || friendRequest.Status == 2 {
		this.ReturnData = util.GenerateAlertMessage(models.FriendError400)
		return
	}

	if result == 1 {
		friendRequest.Status = 1
		var senderFriend models.Friend
		hasSenderFriend, _ := base.DBEngine.Table("friend").Where("owner_uid=? and friend_uid=?", friendRequest.SenderUid, friendRequest.ReceiverUid).And("status=1").Get(new(models.Friend))
		if !hasSenderFriend {
			senderFriend.OwnerUid = friendRequest.SenderUid
			senderFriend.FriendUid = friendRequest.ReceiverUid
			senderFriend.Status = 1
			base.DBEngine.Table("friend").InsertOne(&senderFriend)
		}

		var receiverFriend models.Friend
		hasReceiverFriend, _ := base.DBEngine.Table("friend").Where("owner_uid=? and friend_uid=?", friendRequest.ReceiverUid, friendRequest.SenderUid).And("status=1").Get(new(models.Friend))
		if !hasReceiverFriend {
			receiverFriend.OwnerUid = friendRequest.ReceiverUid
			receiverFriend.FriendUid = friendRequest.SenderUid
			receiverFriend.Status = 1
			base.DBEngine.Table("friend").InsertOne(&receiverFriend)
		}
	} else if result == 2 {
		friendRequest.Status = 2
	}
	base.DBEngine.Table("friend_request").Where("id=?", id).Cols("status").Update(&friendRequest)


	//推送
	if result == 1 {
		var message models.Message
		message.Content = models.AcceptFriendRequest
		message.SenderUid = friendRequest.ReceiverUid
		message.ReceiverUid = friendRequest.SenderUid
		message.Type = 1
		PushSocketMessageToUser(friendRequest.SenderUid, &message, "", 0, "", 3)
	} else if result == 2 {
		var message models.Message
		message.Content = models.RejectFriendRequest
		message.SenderUid = friendRequest.ReceiverUid
		message.ReceiverUid = friendRequest.SenderUid
		message.Type = 1
		PushSocketMessageToUser(friendRequest.SenderUid, &message, "", 0, "", 3)
	}

	this.ReturnData = "success"
}

// @Title 获取好友列表
// @Description 获取好友列表
// @Param	uId					query 			int64			true		"uId"
// @Param	sortType			query 			int				false		"排序类型，1按成为朋友倒序，2按活跃倒序"
// @Param	pageNum				query 			int				true		"page num start from 1"
// @Param	pageTime			query 			int64			true		"page time should be empty when pagenum == 1"
// @Param	pageSize			query 			int				false		"page size default is 15"
// @Success 200 {object} models.FriendListContainer
// @router /getFriendList [get]
func (this *FriendController) GetFriendList() {
	uId := this.MustInt64("uId")
	sortType, _ := this.GetInt("sortType", 1)
	pageNum := this.MustInt("pageNum")
	pageTime, _ := this.GetInt64("pageTime", util.UnixOfBeijingTime())
	pageSize := this.GetPageSize("pageSize")

	totalSql := "select count(1) from friend where deleted_at is null and status=1 and owner_uid='"+strconv.FormatInt(uId, 10)+"' "
	dataSql := "select user.* from user left join friend on friend.owner_uid='"+strconv.FormatInt(uId, 10)+"' and friend.friend_uid=user.u_id and friend.status=1 where friend.id is not null and friend.deleted_at is null "
	if sortType == 1 {
		dataSql += " order by friend.created desc "
	} else if sortType == 2 {
		dataSql += " order by FIELD(user.status, 3, 1, 2, 0) "
	}
	dataSql += " limit "+strconv.Itoa(pageSize*(pageNum-1))+" , "+strconv.Itoa(pageSize)

	total, totalErr := base.DBEngine.SQL(totalSql).Count(new(models.UserShort))
	if totalErr != nil {
		util.Logger.Info("----totalErr---"+totalErr.Error())
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}

	var userList []models.UserShort
	if total > 0 {
		err := base.DBEngine.SQL(dataSql).Find(&userList)
		if err != nil {
			util.Logger.Info("----err---"+err.Error())
			this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
			return
		}
	}

	if userList == nil {
		userList = make([]models.UserShort, 0)
	}

	this.ReturnData = models.FriendListContainer{models.BaseListContainer{total, pageNum, pageTime}, userList}
}

// @Title 处理好友关系
// @Description 处理好友关系
// @Param	currentUid				formData			int64  		true		"当前uid"
// @Param	friendUid				formData			int64  		true		"好友uid"
// @Param	result					formData			int  		true		"处理结果，1删除（单方删除） 2加入黑名单（双方删除）"
// @Success 200 {string} success
// @router /handleFriend [patch]
func (this *FriendController) HandleFriend() {
	currentUid := this.MustInt64("currentUid")
	friendUid := this.MustInt64("friendUid")
	result := this.MustInt("result")

	if result == 1 {
		var ownerFriend models.Friend
		base.DBEngine.Table("friend").Where("owner_uid=?", currentUid).And("friend_uid=?", friendUid).Get(&ownerFriend)
		ownerFriend.Status = 2
		base.DBEngine.Table("friend").Where("id=?", ownerFriend.Id).Cols("status").Update(&ownerFriend)
	} else if result == 2 {
		var ownerFriend models.Friend
		base.DBEngine.Table("friend").Where("owner_uid=?", currentUid).And("friend_uid=?", friendUid).Get(&ownerFriend)
		ownerFriend.Status = 3
		base.DBEngine.Table("friend").Where("id=?", ownerFriend.Id).Cols("status").Update(&ownerFriend)

		var friend models.Friend
		base.DBEngine.Table("friend").Where("owner_uid=?", friendUid).And("friend_uid=?", currentUid).Get(&friend)
		friend.Status = 3
		base.DBEngine.Table("friend").Where("id=?", friend.Id).Cols("status").Update(&friend)
	}

	this.ReturnData = "success"
}

// @Title 获取好友请求列表
// @Description 获取好友请求列表
// @Param	uId					query 			int64			true		"uId"
// @Param	pageNum				query 			int				true		"page num start from 1"
// @Param	pageTime			query 			int64			true		"page time should be empty when pagenum == 1"
// @Param	pageSize			query 			int				false		"page size default is 15"
// @Success 200 {object} models.FriendRequestListContainer
// @router /getFriendRequestList [get]
func (this *FriendController) GetFriendRequestList() {
	uId := this.MustInt64("uId")
	pageNum := this.MustInt("pageNum")
	pageTime, _ := this.GetInt64("pageTime", util.UnixOfBeijingTime())
	pageSize := this.GetPageSize("pageSize")

	totalSql := "select count(1) from friend_request where deleted_at is null and receiver_uid='"+strconv.FormatInt(uId, 10)+"' and exists(select 1 from user where user.u_id=friend_request.sender_uid and user.deleted_at is null) "
	dataSql := "select friend_request.*, user.nick_name as sender_nick_name from friend_request left join user on friend_request.sender_uid=user.u_id where user.u_id is not null and friend_request.deleted_at is null and friend_request.receiver_uid='"+strconv.FormatInt(uId, 10)+"' "
	dataSql += " order by friend_request.created desc "
	dataSql += " limit "+strconv.Itoa(pageSize*(pageNum-1))+" , "+strconv.Itoa(pageSize)

	total, totalErr := base.DBEngine.SQL(totalSql).Count(new(models.FriendRequest))
	if totalErr != nil {
		util.Logger.Info("----totalErr---"+totalErr.Error())
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}

	var friendRequestList []models.FriendRequestInfo
	if total > 0 {
		err := base.DBEngine.SQL(dataSql).Find(&friendRequestList)
		if err != nil {
			util.Logger.Info("----err---"+err.Error())
			this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
			return
		}
	}

	if friendRequestList == nil {
		friendRequestList = make([]models.FriendRequestInfo, 0)
	}

	this.ReturnData = models.FriendRequestListContainer{models.BaseListContainer{total, pageNum, pageTime}, friendRequestList}
}


