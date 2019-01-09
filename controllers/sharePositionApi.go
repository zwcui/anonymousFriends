/*
@Time : 2019/1/2 上午10:12 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"strconv"
	"anonymousFriends/util"
)

//共享地理位置模块
type SharePositionController struct {
	apiController
}

func (this *SharePositionController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
		{"/sendSharePositionRequest", "post"},
		{"/getSharePositionRequest", "get"},
		{"/handleSharePositionRequest", "patch"},
		{"/getSharePositionGroupUserList", "get"},
	}
	this.bathAuth()
}

// @Title 发送位置共享请求
// @Description 发送位置共享请求
// @Param	uId						formData		int64  		true		"用户id"
// @Param	groupId					formData		int64  		true		"组id"
// @Success 200 {string} success
// @router /sendSharePositionRequest [post]
func (this *SharePositionController) SendSharePositionRequest() {
	uId := this.MustInt64("uId")
	groupId := this.MustInt64("groupId")

	var user models.UserShort
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)

	var sharePositionGroup models.SharePositionGroup
	hasStoredSharePositionGroup, _ := base.DBEngine.Table("share_position_group").Where("group_id=?", groupId).And("(status=0 or status=1)").Get(&sharePositionGroup)
	if !hasStoredSharePositionGroup {
		sharePositionGroup.GroupId = groupId
		sharePositionGroup.Originator = uId
		sharePositionGroup.Status = 0
		sharePositionGroup.Remark = "用户"+strconv.FormatInt(uId, 10)+"发起位置共享;"
		base.DBEngine.Table("share_position_group").InsertOne(&sharePositionGroup)
	} else {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError100)
		return
	}

	var sharePositionMember models.SharePositionMember
	sharePositionMember.SharePositionGroupId = sharePositionGroup.Id
	sharePositionMember.UId = uId
	sharePositionMember.Status = 1
	base.DBEngine.Table("share_position_member").InsertOne(&sharePositionMember)

	//推送消息给组其他成员
	var message models.Message
	message.Type = 4
	message.Content = user.NickName + models.SharePositionRequest
	var memberList []models.Member
	base.DBEngine.Table("member").Where("group_id=?", groupId).Find(&memberList)
	if memberList == nil {
		memberList = make([]models.Member, 0)
	}
	for _, member := range memberList {
		message.ReceiverUid = member.UId
		PushCommonMessageToUser(member.UId, &message, "", 0 , "")
	}

	this.ReturnData = "success"
}

// @Title 查看位置共享（每次进IM调用，接受或拒绝后重新调用）
// @Description 查看位置共享（每次进IM调用，接受或拒绝后重新调用）
// @Param	uId						query			int64  		true		"用户id"
// @Param	groupId					query			int64  		true		"组id"
// @Success 200 {object} models.SharePositionInfoContainer
// @router /getSharePositionRequest [get]
func (this *SharePositionController) GetSharePositionRequest() {
	uId := this.MustInt64("uId")
	groupId := this.MustInt64("groupId")

	var sharePositionInfo models.SharePositionInfo
	var sharePositionGroup models.SharePositionGroup
	hasGroup, _ := base.DBEngine.Table("share_position_group").Where("group_id=?", groupId).And("(status=0 or status=1)").Get(&sharePositionGroup)
	if hasGroup {
		var user models.UserShort
		base.DBEngine.Table("user").Where("u_id=?", sharePositionGroup.Originator).Get(&user)
		sharePositionInfo.SharePositionGroup = sharePositionGroup
		sharePositionInfo.SenderNickName = user.NickName
		sharePositionInfo.SenderAvatar = user.Avatar

		var sharePositionMember models.SharePositionMember
		hasMember, _ := base.DBEngine.Table("share_position_member").Where("share_position_group_id=?", sharePositionGroup.GroupId).And("u_id=?", uId).Get(&sharePositionMember)
		if hasMember {
			if sharePositionMember.Status == 2 {
				sharePositionInfo.Status = 2
			}
		}
	} else {
		sharePositionInfo.SharePositionGroup = sharePositionGroup
		sharePositionInfo.SenderNickName = ""
		sharePositionInfo.SenderAvatar = ""
	}

	//base.DBEngine.Table("share_position_group").Select("share_position_group.*, user.nick_name as sender_nick_name, user.avatar as sender_avatar").Join("LEFT OUTER", "user", "user.u_id=share_position_group.originator").Where("share_position_group.group_id=?", groupId).And("(share_position_group.status=0 or share_position_group.status=1)").And("not exists(select 1 from share_position_member where share_position_group.id=share_position_member.share_position_groupId and share_position_member.u_id=?)", uId).Get(&sharePositionInfo)

	this.ReturnData = models.SharePositionInfoContainer{sharePositionInfo}
}

// @Title 处理位置共享请求
// @Description 处理位置共享请求
// @Param	id						formData		int64  		true		"id"
// @Param	uId						formData		int64  		true		"用户id"
// @Param	result					formData		int  		true		"处理结果，1接受 2拒绝 3接收后退出位置共享 4发起人自己取消位置共享"
// @Success 200 {object} models.SharePositionInfoContainer
// @router /handleSharePositionRequest [patch]
func (this *SharePositionController) HandleSharePositionRequest() {
	id := this.MustInt64("id")
	uId := this.MustInt64("uId")
	result := this.MustInt("result")

	var sharePositionGroup models.SharePositionGroup
	hasGroup, _ := base.DBEngine.Table("share_position_group").Where("id=?", id).Get(&sharePositionGroup)
	if !hasGroup {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError200)
		return
	}
	if sharePositionGroup.Status == 2 {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError300)
		return
	}

	var sharePositionMember models.SharePositionMember
	hasMember, _ := base.DBEngine.Table("share_position_member").Where("share_position_group_id=?", sharePositionGroup.Id).And("u_id=?", uId).Get(&sharePositionMember)
	if hasMember && (result == 1 || result == 2) {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError400)
		return
	}

	if hasMember {
		if result == 3 {
			sharePositionMember.Status = 3
			base.DBEngine.Table("share_position_member").Where("id=?", sharePositionMember.Id).Cols("status").Update(&sharePositionMember)

			sharePositionGroup.Remark += "用户"+strconv.FormatInt(uId, 10)+"接收后退出位置共享;"
			//如果位置共享人数小于等于1人，则关闭组
			total, _ := base.DBEngine.Table("share_position_member").Where("share_position_group_id=?", id).And("status=1").Count(new(models.SharePositionMember))
			if total <= 1 {
				sharePositionGroup.Status = 2
			}
			base.DBEngine.Table("share_position_group").Where("id=?", id).Cols("status", "remark").Update(&sharePositionGroup)
		} else if result == 4 {
			sharePositionMember.Status = 3
			base.DBEngine.Table("share_position_member").Where("id=?", sharePositionMember.Id).Cols("status").Update(&sharePositionMember)
			sharePositionGroup.Status = 2
			sharePositionGroup.Remark += "用户"+strconv.FormatInt(uId, 10)+"发起人自己取消位置共享;"
			base.DBEngine.Table("share_position_group").Where("id=?", id).Cols("status", "remark").Update(&sharePositionGroup)
		}
	} else {
		sharePositionMember.SharePositionGroupId = sharePositionGroup.Id
		sharePositionMember.UId = uId
		if result == 1 {
			sharePositionMember.Status = 1
			if sharePositionGroup.Status != 1 {
				sharePositionGroup.Status = 1
			}
			sharePositionGroup.Remark += "用户"+strconv.FormatInt(uId, 10)+"接收位置共享;"
			base.DBEngine.Table("share_position_group").Where("id=?", id).Cols("status", "remark").Update(&sharePositionGroup)
		} else if result == 2 {
			sharePositionMember.Status = 2
			//如果位置共享人数小于等于1人，则关闭组
			total, _ := base.DBEngine.Table("share_position_member").Where("share_position_group_id=?", id).And("status=1").Count(new(models.SharePositionMember))
			if total <= 1 {
				sharePositionGroup.Status = 2
			}
			sharePositionGroup.Remark += "用户"+strconv.FormatInt(uId, 10)+"拒绝位置共享;"
			base.DBEngine.Table("share_position_group").Where("id=?", id).Cols("status", "remark").Update(&sharePositionGroup)
		}
		base.DBEngine.Table("share_position_member").InsertOne(&sharePositionMember)
	}

	this.ReturnData = "success"
}

// @Title 获取位置共享组中的用户列表
// @Description 获取位置共享组中的用户列表
// @Param	id				query			int64	  		true		"id"
// @Success 200 {object} models.UserList
// @router /getSharePositionGroupUserList [get]
func (this *SharePositionController) GetSharePositionGroupUserList() {
	id := this.MustInt64("id")

	var sharePositionGroup models.SharePositionGroup
	hasGroup, _ := base.DBEngine.Table("share_position_group").Where("id=?", id).Get(&sharePositionGroup)
	if !hasGroup {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError200)
		return
	}

	if sharePositionGroup.Status == 2 {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError300)
		return
	}

	var userList []models.UserShort
	base.DBEngine.Table("user").Join("LEFT OUTER", "share_position_member", "share_position_member.u_id=user.u_id").Select("user.*").Where("share_position_member.id is not null and share_position_member.share_position_group_id=? and share_position_member.status=1").Find(&userList)

	if userList == nil {
		userList = make([]models.UserShort, 0)
	}

	this.ReturnData = models.UserList{userList}
}
