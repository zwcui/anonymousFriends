/*
@Time : 2019/1/4 下午5:08 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"strconv"
)

//建议与举报模块
type ReportController struct {
	apiController
}

func (this *ReportController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
		{"/postReport", "post"},
	}
	this.bathAuth()
}

// @Title 提出建议或举报
// @Description 提出建议或举报
// @Param	uId					formData		int64  			true		"发起人"
// @Param	type				formData		int	  			true		"类型，1:对app的建议 2:举报朋友圈 3:举报漂流瓶 4:举报用户"
// @Param	typeId				formData		int64	  		false		"类型对应的id"
// @Param	content				formData		string  		false		"内容，如涉嫌广告、非法、色情等"
// @Success 200 {string} success
// @router /postReport [post]
func (this *ReportController) PostReport() {
	uId := this.MustInt64("uId")
	reportType := this.MustInt("type")
	typeId, _ := this.GetInt64("typeId", 0)
	content := this.GetString("content", "")

	var report models.Report
	report.SenderUid = uId
	report.Type = reportType
	report.TypeId = typeId
	report.Content = content
	base.DBEngine.Table("report").InsertOne(&report)

	//推送通知管理员
	var adminList []models.UserShort
	base.DBEngine.Table("user").Where("type=1").Find(&adminList)
	var message models.Message
	if reportType == 1 {
		message.Content = "用户"+strconv.FormatInt(uId, 10)+"提出建议："+content
	} else if reportType == 2 {
		message.Content = "用户"+strconv.FormatInt(uId, 10)+"举报朋友圈(id:"+strconv.FormatInt(typeId, 10)+")："+content
	} else if reportType == 3 {
		message.Content = "用户"+strconv.FormatInt(uId, 10)+"举报漂流瓶(id:"+strconv.FormatInt(typeId, 10)+")："+content
	} else if reportType == 4 {
		message.Content = "用户"+strconv.FormatInt(uId, 10)+"举报用户(id:"+strconv.FormatInt(typeId, 10)+")："+content
	}

	var adminNotice models.AdminNotice
	adminNotice.Type = 3
	adminNotice.TypeId = report.Id
	adminNotice.Content = message.Content
	adminNotice.Status = 0
	base.DBEngine.Table("admin_notice").InsertOne(&adminNotice)

	for _, admin := range adminList {
		message.ReceiverUid = admin.UId
		PushCommonMessageToUser(admin.UId, &message, "", 0, "report")
	}

	this.ReturnData = "success"
}