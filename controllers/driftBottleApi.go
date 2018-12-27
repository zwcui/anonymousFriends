/*
@Time : 2018/12/26 下午5:04 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/util"
	"anonymousFriends/base"
	"strconv"
)

//漂流瓶模块
type DriftBottleController struct {
	apiController
}

func (this *DriftBottleController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
		{"/throwDriftBottle", "post"},
		{"/pickUpDriftBottle", "get"},
		{"/handleDriftBottle", "patch"},
		{"/getMyDriftBottleList", "get"},
	}
	this.bathAuth()
}

// @Title 扔漂流瓶
// @Description 扔漂流瓶
// @Param	uId					formData		int64	  		true		"uId"
// @Param	bottleType			formData		int		  		true		"类型，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶"
// @Param	bottleName			formData		string	  		true		"类型名称，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶"
// @Param	content				formData		string  		false		"文字内容"
// @Param	picture				formData		string  		false		"图片内容，多个英文逗号隔开"
// @Param	position			formData		string  		false		"位置名称，如全家创意产业园店"
// @Param   weather	        	formData        string  		false       "天气"
// @Param   province        	formData        string  		false       "省"
// @Param   city    			formData        string  		false       "市"
// @Param	area				formData		string	  		false		"区"
// @Param	longitude			formData		float64  		false		"经度"
// @Param   latitude         	formData        float64    		false       "纬度"
// @Param   status	         	formData        int	    		true	    "状态，0未抛出暂存，1抛出"
// @Success 200 {string} success
// @router /throwDriftBottle [post]
func (this *DriftBottleController) ThrowDriftBottle() {
	uId := this.MustInt64("uId")
	bottleType := this.MustInt("bottleType")
	bottleName := this.MustInt("bottleName")
	content := this.GetString("content")
	picture := this.GetString("picture", "")
	position := this.GetString("position", "")
	weather := this.GetString("weather", "")
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitude, _ := this.GetFloat("longitude", 0)
	latitude, _ := this.GetFloat("latitude", 0)
	status := this.MustInt("status")

	var driftBottle models.DriftBottle
	driftBottle.BottleType = bottleType
	driftBottle.BottleName = bottleName
	driftBottle.SenderUid = uId
	driftBottle.Content = content
	driftBottle.Picture = picture
	driftBottle.Position = position
	driftBottle.Weather = weather
	driftBottle.Province = province
	driftBottle.City = city
	driftBottle.Area = area
	driftBottle.Longitude = longitude
	driftBottle.Latitude = latitude
	driftBottle.Status = status

	if models.ExpiryTime != 0 {
		driftBottle.ExpiryTime = util.UnixOfBeijingTime() + models.ExpiryTime * 60 * 60
	}
	base.DBEngine.Table("drift_bottle").InsertOne(&driftBottle)

	this.ReturnData = "success"
}

// @Title 拾漂流瓶或查看漂流瓶
// @Description 拾漂流瓶或查看漂流瓶
// @Param	uId					query			int64	  		true		"uId"
// @Param	bottleId			query			int64	  		false		"漂流瓶id，有则查看详情，无则拾起漂流瓶"
// @Success 200 {object} models.DriftBottleInfo
// @router /pickUpDriftBottle [get]
func (this *DriftBottleController) PickUpDriftBottle() {
	uId := this.MustInt64("uId")
	bottleId, _ := this.GetInt64("bottleId", 0)

	var driftBottle models.DriftBottle
	if bottleId == 0 {
		randomSql := "SELECT * FROM drift_bottle WHERE bottle_id >= ((SELECT MAX(bottle_id) FROM drift_bottle)-(SELECT MIN(bottle_id) FROM drift_bottle)) * RAND() + (SELECT MIN(bottle_id) FROM drift_bottle)  LIMIT 1"
		base.DBEngine.SQL(randomSql).Get(&driftBottle)

		driftBottle.ReceiverUid = uId
		driftBottle.Status = 2
		base.DBEngine.Table("drift_bottle").Where("bottle_id=?", driftBottle.BottleId).Cols("receiver_uid", "status").Update(&driftBottle)
	} else {
		base.DBEngine.Table("drift_bottle").Where("bottle_id=?", bottleId).Get(&driftBottle)
	}

	var commentList []models.CommentInfo
	err := base.DBEngine.Table("comment").Select("comment.*, sender.nick_name as sender_nick_name, sender.u_id as sender_uid, receiver.nick_name as receiver_nick_name, receiver.u_id as receiver_uid").Join("LEFT OUTER", "user sender", "sender.u_id=comment.sender_uid").Join("LEFT OUTER", "user receiver", "receiver.u_id=comment.receiver_uid").Where("comment.type=2 and comment.type_id=?", bottleId).Desc("created").Find(&commentList)
	if err != nil {
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}
	if commentList == nil {
		commentList = make([]models.CommentInfo, 0)
	}

	this.ReturnData = models.DriftBottleInfo{driftBottle, commentList}
}

// @Title 处理漂流瓶
// @Description 处理漂流瓶
// @Param	bottleId			formData		int64	  		true		"bottleId"
// @Param	uId					formData		int64	  		true		"uId"
// @Param	result				formData		int		  		true		"处理，1回复，2扔回大海"
// @Param	content				formData		string  		false		"回复内容"
// @Success 200 {string} success
// @router /handleDriftBottle [patch]
func (this *DriftBottleController) HandleDriftBottle() {
	bottleId := this.MustInt64("bottleId")
	uId := this.MustInt64("uId")
	result := this.MustInt("result")
	content := this.GetString("content", "")

	var driftBottle models.DriftBottle
	hasDriftBottle, _ := base.DBEngine.Table("drift_bottle").Where("bottle_id=?", bottleId).Get(&driftBottle)
	if !hasDriftBottle {
		this.ReturnData = util.GenerateAlertMessage(models.DriftBottleError100)
		return
	}

	if result == 1 {
		if driftBottle.ReceiverUid == 0 {
			driftBottle.ReceiverUid = uId
		}

		var comment models.Comment
		comment.Type = 2
		comment.TypeId = driftBottle.BottleId
		comment.Content = content
		comment.SenderUid = uId
		if uId == driftBottle.SenderUid {
			comment.ReceiverUid = driftBottle.ReceiverUid
		} else {
			comment.ReceiverUid = driftBottle.SenderUid
		}
		base.DBEngine.Table("comment").InsertOne(&comment)

		driftBottle.ReplyNum += 1
		driftBottle.Status = 2
		base.DBEngine.Table("drift_bottle").Where("bottle_id=?", bottleId).Cols("receiver_uid", "reply_num", "status").Update(&driftBottle)

		var message models.Message
		message.SenderUid = uId
		message.ReceiverUid = comment.ReceiverUid
		message.Content = models.CommentOnDriftBottle
		message.Type = 3
		PushCommonMessageToUser(comment.ReceiverUid, &message, "", 0, "")
	} else if result == 2 {
		driftBottle.Remark += "被uId:"+strconv.FormatInt(uId, 10)+"抛回;"
		base.DBEngine.Table("drift_bottle").Where("bottle_id=?", bottleId).Cols("remark").Update(&driftBottle)
	}

	this.ReturnData = "success"
}

// @Title 我的漂流瓶
// @Description 我的漂流瓶
// @Param	uId					query			int64	  		true		"uId"
// @Param	type				query			int64	  		true		"类型，1为我抛出的，2为我拾起的"
// @Param	pageNum				query 			int				true		"page num start from 1"
// @Param	pageTime			query 			int64			true		"page time should be empty when pagenum == 1"
// @Param	pageSize			query 			int				false		"page size default is 15"
// @Success 200 {object} models.DriftBottleListContainer
// @router /getMyDriftBottleList [get]
func (this *DriftBottleController) GetMyDriftBottleList() {
	uId := this.MustInt64("uId")
	queryType := this.MustInt("type")
	pageNum := this.MustInt("pageNum")
	pageTime, _ := this.GetInt64("pageTime", util.UnixOfBeijingTime())
	pageSize := this.GetPageSize("pageSize")

	totalSql := "select count(1) from drift_bottle where drift_bottle.deleted_at is null "
	dataSql := "select drift_bottle.* from drift_bottle where drift_bottle.deleted_at is null "
	if queryType == 1 {
		totalSql += " and drift_bottle.sender_uid='"+strconv.FormatInt(uId, 10)+"' "
		dataSql += " and drift_bottle.sender_uid='"+strconv.FormatInt(uId, 10)+"' "
	} else if queryType == 2 {
		totalSql += " and drift_bottle.receiver_uid='"+strconv.FormatInt(uId, 10)+"' "
		dataSql += " and drift_bottle.receiver_uid='"+strconv.FormatInt(uId, 10)+"' "
	}

	dataSql += " order by drift_bottle.created desc limit "+strconv.Itoa(pageSize*(pageNum-1))+" , "+strconv.Itoa(pageSize)

	total, totalErr := base.DBEngine.SQL(totalSql).Count(new(models.DriftBottle))
	if totalErr != nil {
		util.Logger.Info("----totalErr---"+totalErr.Error())
		this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
		return
	}

	var driftBottleList []models.DriftBottle
	if total > 0 {
		err := base.DBEngine.SQL(dataSql).Find(&driftBottleList)
		if err != nil {
			util.Logger.Info("----err---"+err.Error())
			this.ReturnData = util.GenerateAlertMessage(models.CommonError100)
			return
		}
	}

	if driftBottleList == nil {
		driftBottleList = make([]models.DriftBottle, 0)
	}

	this.ReturnData = models.DriftBottleListContainer{models.BaseListContainer{total, pageNum, pageTime}, driftBottleList}
}