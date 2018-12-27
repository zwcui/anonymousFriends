/*
@Time : 2018/12/27 下午3:15 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
)

//用户模块
type TagController struct {
	apiController
}

func (this *TagController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
}

// @Title 获取所有标签
// @Description 获取所有标签
// @Success 200 {object} models.TagList
// @router /getTagList [get]
func (this *TagController) GetTagList() {
	var tagList []models.Tag
	base.DBEngine.Table("tag").Where("status=1").Asc("tag_order").Desc("created").Find(&tagList)
	if tagList == nil {
		tagList = make([]models.Tag, 0)
	}

	this.ReturnData = models.TagList{tagList}
}

// @Title 新增标签
// @Description 新增标签
// @Param	tagName				formData		string  		true		"标签名称"
// @Param	tagImage			formData		string  		false		"标签图片"
// @Param	tagOrder			formData		string  		false		"标签排序"
// @Param	parentTagId			formData		int64	  		false		"父标签id"
// @Param	status				formData		int		  		false		"状态，1正常，2隐藏"
// @Success 200 {string} success
// @router /addTag [post]
func (this *TagController) AddTag() {
	tagName := this.MustString("tagName")
	tagImage := this.GetString("tagImage", "")
	tagOrder := this.GetString("tagOrder", "")
	parentTagId, _ := this.GetInt64("parentTagId", 0)
	status, _ := this.GetInt("status", 1)

	var tag models.Tag
	tag.TagName = tagName
	tag.TagImage = tagImage
	tag.TagOrder = tagOrder
	tag.ParentTagId = parentTagId
	tag.Status = status
	base.DBEngine.Table("tag").InsertOne(&tag)

	this.ReturnData = "success"
}