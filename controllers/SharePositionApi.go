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
	this.NeedBaseAuthList = []RequestPathAndMethod{}
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
		sharePositionGroup.Remark = "用户"+strconv.FormatInt(uId, 10)+"发起位置共享"
		base.DBEngine.Table("share_position_group").InsertOne(&sharePositionGroup)
	} else {
		this.ReturnData = util.GenerateAlertMessage(models.SharePositionError100)
		return
	}

	var sharePositionMember models.SharePositionMember
	sharePositionMember.SharePositionGroupId = sharePositionGroup.GroupId
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

// @Title 查看位置共享请求
// @Description 查看位置共享请求
// @Param	uId						formData		int64  		true		"用户id"
// @Param	groupId					formData		int64  		true		"组id"
// @Success 200 {object} models.SharePositionInfoContainer
// @router /getSharePositionRequest [get]
func (this *SharePositionController) GetSharePositionRequest() {
	uId := this.MustInt64("uId")
	groupId := this.MustInt64("groupId")

	var sharePositionInfo models.SharePositionInfo
	base.DBEngine.Table("share_position_group").Select("share_position_group.*, user.nick_name as sender_nick_name, user.avatar as sender_avatar").Join("LEFT OUTER", "user", "user.u_id=share_position_group.originator").Where("share_position_group.group_id=?", groupId).And("(share_position_group.status=0 or share_position_group.status=1)").And("not exists(select 1 from share_position_member where share_position_group.id=share_position_member.share_position_groupId and share_position_member.u_id=?)", uId).Get(&sharePositionInfo)

	this.ReturnData = models.SharePositionInfoContainer{sharePositionInfo}
}