/*
@Time : 2018/12/21 下午2:00
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/base"
	"anonymousFriends/models"
	"strconv"
	"anonymousFriends/util"
)

//朋友圈模块
type SocialDynamicsController struct {
	apiController
}

func (this *SocialDynamicsController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{{"/postSocialDynamic", "post"}, {"/deleteSocialDynamic", "delete"}, {"/getSocialDynamicList", "get"}, {"/likeSocialDynamic", "post"}}
	this.bathAuth()
}

// @Title 发布朋友圈
// @Description 发布朋友圈
// @Param	uId					formData		int64	  		true		"uId"
// @Param	content				formData		string  		false		"文字内容"
// @Param	picture				formData		string  		false		"图片内容，多个英文逗号隔开"
// @Param	position			formData		string  		false		"位置名称，如全家创意产业园店"
// @Param   weather	        	formData        string  		false       "天气"
// @Param   province        	formData        string  		false       "省"
// @Param   city    			formData        string  		false       "市"
// @Param	area				formData		string	  		false		"区"
// @Param	longitude			formData		float64  		false		"经度"
// @Param   latitude         	formData        float64    		false       "纬度"
// @Success 200 {string} success
// @router /postSocialDynamic [post]
func (this *SocialDynamicsController) PostSocialDynamic() {
	uId := this.MustInt64("uId")
	content := this.GetString("content")
	picture := this.GetString("picture", "")
	position := this.GetString("position", "")
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitude, _ := this.GetFloat("longitude", 0)
	latitude, _ := this.GetFloat("latitude", 0)

	var socialDynamics models.SocialDynamics
	socialDynamics.UId = uId
	socialDynamics.Content = content
	socialDynamics.Picture = picture
	socialDynamics.LikeNum = 0
	socialDynamics.Position = position
	socialDynamics.Province = province
	socialDynamics.City = city
	socialDynamics.Area = area
	socialDynamics.Longitude = longitude
	socialDynamics.Latitude = latitude
	base.DBEngine.Table("social_dynamics").InsertOne(&socialDynamics)

	this.ReturnData = "success"
}

// @Title 删除朋友圈
// @Description 删除朋友圈
// @Param	id					query			int64	  		true		"id"
// @Success 200 {string} success
// @router /deleteSocialDynamic [delete]
func (this *SocialDynamicsController) DeleteSocialDynamic() {
	id := this.MustInt64("id")

	var socialDynamics models.SocialDynamics
	base.DBEngine.Table("social_dynamics").Where("id=?", id).Delete(&socialDynamics)

	this.ReturnData = "success"
}

// @Title 查看朋友圈
// @Description 查看朋友圈
// @Param	currentUid			query			int64	  		true		"当前查看人"
// @Param	uId					query			int64	  		false		"uId"
// @Param	position			query			string  		false		"位置名称，如全家创意产业园店"
// @Param   province        	query 	        string  		false       "省"
// @Param   city    			query 	        string  		false       "市"
// @Param	area				query			string	  		false		"区"
// @Param	longitude			query			float64  		false		"经度"
// @Param   latitude         	query	        float64    		false       "纬度"
// @Param	pageNum				query 			int				true		"page num start from 1"
// @Param	pageTime			query 			int64			true		"page time should be empty when pagenum == 1"
// @Param	pageSize			query 			int				false		"page size default is 15"
// @Success 200 {object} models.SocialDynamicListContainer
// @router /getSocialDynamicList [get]
func (this *SocialDynamicsController) GetSocialDynamicList() {
	currentUid := this.MustInt64("currentUid")
	uId, _ := this.GetInt64("uId", 0)
	position := this.GetString("position", "")
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitude, _ := this.GetFloat("longitude", 0)
	latitude, _ := this.GetFloat("latitude", 0)
	pageNum := this.MustInt("pageNum")
	pageTime, _ := this.GetInt64("pageTime", util.UnixOfBeijingTime())
	pageSize := this.GetPageSize("pageSize")

	totalSql := "select count(1) from social_dynamics where deleted_at is null "
	dataSql := "select social_dynamics.*, case when exists(select 1 from `like` where `like`.id=social_dynamics.id and `like`.u_id='"+strconv.FormatInt(currentUid, 10)+"' and `like`.type=1) then 1 else 0 end as is_like from social_dynamics where deleted_at is null "
	if uId != 0 {
		totalSql += " and social_dynamics.u_id='"+strconv.FormatInt(uId, 10)+"' "
		dataSql += " and social_dynamics.u_id='"+strconv.FormatInt(uId, 10)+"' "
	}
	if position != "" {
		totalSql += " and social_dynamics.position='"+position+"' "
		dataSql += " and social_dynamics.position='"+position+"' "
	}
	if province != "" {
		totalSql += " and social_dynamics.province='"+province+"' "
		dataSql += " and social_dynamics.province='"+province+"' "
	}
	if city != "" {
		totalSql += " and social_dynamics.city='"+city+"' "
		dataSql += " and social_dynamics.city='"+city+"' "
	}
	if area != "" {
		totalSql += " and social_dynamics.area='"+area+"' "
		dataSql += " and social_dynamics.area='"+area+"' "
	}
	if longitude != 0 {
		totalSql += " and social_dynamics.longitude='"+strconv.FormatFloat(longitude, 'f', 6, 64)+"' "
		dataSql += " and social_dynamics.longitude='"+strconv.FormatFloat(longitude, 'f', 6, 64)+"' "
	}
	if latitude != 0 {
		totalSql += " and social_dynamics.latitude='"+strconv.FormatFloat(latitude, 'f', 6, 64)+"' "
		dataSql += " and social_dynamics.latitude='"+strconv.FormatFloat(latitude, 'f', 6, 64)+"' "
	}

	dataSql += " order by social_dynamics.created desc limit "+strconv.Itoa(pageSize*(pageNum-1))+" , "+strconv.Itoa(pageSize)

	total, totalErr := base.DBEngine.SQL(totalSql).Count(new(models.SocialDynamics))
	if totalErr != nil {
		util.Logger.Info("----totalErr---"+totalErr.Error())
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}

	var socialDynamicList []models.SocialDynamicInfo
	if total > 0 {
		err := base.DBEngine.SQL(dataSql).Find(&socialDynamicList)
		if err != nil {
			util.Logger.Info("----err---"+err.Error())
			this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
			return
		}
	}

	if socialDynamicList == nil {
		socialDynamicList = make([]models.SocialDynamicInfo, 0)
	}

	var resultSocialDynamicList []models.SocialDynamicInfo
	for _, socialDynamic := range socialDynamicList {
		var result models.SocialDynamicInfo
		result.SocialDynamics = socialDynamic.SocialDynamics
		result.IsLike = socialDynamic.IsLike
		var commentList []models.Comment
		base.DBEngine.Table("comment").Where("type=1 and type_id=?", result.Id).Find(&commentList)
		result.CommentList = commentList
		resultSocialDynamicList = append(resultSocialDynamicList, result)
	}

	this.ReturnData = models.SocialDynamicListContainer{models.BaseListContainer{total, pageNum, pageTime}, resultSocialDynamicList}
}

// @Title 朋友圈点赞
// @Description 朋友圈点赞
// @Param	currentUid			formData		int64	  		true		"当前查看人"
// @Param	id					formData		int64	  		true		"朋友圈id"
// @Param	type				formData		int		  		false		"类型，1点赞 2取消点赞"
// @Success 200 {string} success
// @router /likeSocialDynamic [post]
func (this *SocialDynamicsController) LikeSocialDynamic() {
	currentUid := this.MustInt64("currentUid")
	id := this.MustInt64("id")
	likeType, _ := this.GetInt("type", 1)

	if likeType == 1 {
		var like models.Like
		like.Id = id
		like.UId = currentUid
		like.Type = 1
		base.DBEngine.Table("like").InsertOne(&like)
	} else {
		base.DBEngine.Table("like").Where("id=?", id).And("u_id=?", currentUid).And("type=1").Delete(new(models.Like))
	}

	this.ReturnData = "success"
}