/*
@Time : 2018/12/27 下午4:21 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"strconv"
	"anonymousFriends/util"
	"anonymousFriends/models"
	"anonymousFriends/base"
)

//管理员通知模块
type AdminNoticeController struct {
	apiController
}

func (this *AdminNoticeController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
	}
	this.bathAuth()
}


// @Title 获取所有管理员通知
// @Description 获取所有管理员通知
// @Param	pageNum				query 			int				true		"page num start from 1"
// @Param	pageTime			query 			int64			true		"page time should be empty when pagenum == 1"
// @Param	pageSize			query 			int				false		"page size default is 15"
// @Success 200 {object} models.AdminNoticeListContainer
// @router /getAdminNoticeList [get]
func (this *AdminNoticeController) GetAdminNoticeList() {
	pageNum := this.MustInt("pageNum")
	pageTime, _ := this.GetInt64("pageTime", util.UnixOfBeijingTime())
	pageSize := this.GetPageSize("pageSize")

	totalSql := "select count(1) from admin_notice where deleted_at is null "
	dataSql := "select admin_notice.* from admin_notice where admin_notice.deleted_at is null "

	dataSql += " order by admin_notice.created desc limit "+strconv.Itoa(pageSize*(pageNum-1))+" , "+strconv.Itoa(pageSize)

	total, totalErr := base.DBEngine.SQL(totalSql).Count(new(models.AdminNotice))
	if totalErr != nil {
		util.Logger.Info("----totalErr---"+totalErr.Error())
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}

	var adminNoticeList []models.AdminNotice
	if total > 0 {
		err := base.DBEngine.SQL(dataSql).Find(&adminNoticeList)
		if err != nil {
			util.Logger.Info("----err---"+err.Error())
			this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
			return
		}
	}

	if adminNoticeList == nil {
		adminNoticeList = make([]models.AdminNotice, 0)
	}

	this.ReturnData = models.AdminNoticeListContainer{models.BaseListContainer{total, pageNum, pageTime}, adminNoticeList}
}

// @Title 更新管理员通知状态
// @Description 更新管理员通知状态
// @Param	id					formData 		int64			true		"id"
// @Param	status				formData		int				true		"状态，0未处理，1已处理，2已忽略"
// @Success 200 {string} success
// @router /updateAdminNoticeStatus [patch]
func (this *AdminNoticeController) UpdateAdminNoticeStatus() {
	id := this.MustInt64("id")
	status := this.MustInt("status")

	var adminNotice models.AdminNotice
	base.DBEngine.Table("admin_notice").Where("id=?", id).Get(&adminNotice)

	adminNotice.Status = status
	base.DBEngine.Table("admin_notice").Where("id=?", id).Cols("status").Update(&adminNotice)

	this.ReturnData = "success"
}
