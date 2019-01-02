package controllers

import (
	"anonymousFriends/base"
	"anonymousFriends/models"
	"anonymousFriends/util"
	"gopkg.in/mgo.v2/bson"
)

//消息模块
type MessageController struct {
	apiController
}

//当前api请求之前调用，用于配置哪些接口需要进行head身份验证
func (this *MessageController) Prepare(){
	//this.NeedBaseAuthList = []RequestPathAndMethod{{".+", "post"}, {".+", "patch"}, {".+", "delete"}}
	this.NeedBaseAuthList = []RequestPathAndMethod{{"/addGroup$", "post"}}
	this.bathAuth()
}

// @Title 测试新增MongoDB消息
// @Description 测试新增MongoDB消息
// @Param	content		formData		string  		true		"内容"
// @Success 200 {object} models.UserChatSocketMessage
// @Failure 403 create message failed
// @router / [post]
func (this *MessageController) Post() {
	content := this.MustString("content")

	//var user models.User
	//user.NickName = nickName
	//base.DBEngine.Table("user").InsertOne(&user)

	session, mongoDB := base.MongoDB()
	defer session.Close()
	c := mongoDB.C("message")
	err := c.Insert(&models.UserChatSocketMessage{})
	if err != nil {
		util.Logger.Info("Insert err:"+err.Error())
	}

	result := models.UserChatSocketMessage{}
	err = c.Find(bson.M{"content":content}).One(&result)
	if err != nil {
		util.Logger.Info("One err:"+err.Error())
	}
	//util.Logger.Info("One result:"+strconv.FormatInt(result.MId, 10))
	//util.Logger.Info("One result:"+strconv.FormatInt(result.GroupId, 10))
	//util.Logger.Info("One result:"+strconv.FormatInt(result.SenderUid, 10))
	//util.Logger.Info("One result:"+strconv.Itoa(result.Type))
	//util.Logger.Info("One result:"+result.Content)
	//util.Logger.Info("One result:"+strconv.FormatInt(result.Created, 10))


	this.ReturnData = result
}


// @Title 加入聊天组，包括一对一和群聊
// @Description 加入聊天组，包括一对一和群聊
// @Param	groupId				formData		int64	  		false		"聊天组id"
// @Param	groupType			formData		int		  		true		"组类型 1:一对一 2:一对多"
// @Param	groupName			formData		string	  		false		"聊天话题，群聊名称"
// @Param	uId1				formData		int64  			true		"聊天人1"
// @Param	uId2				formData		int64  			false		"聊天人2，一对一时必传"
// @Success 200 {object} models.GroupInfo
// @router /addGroup [post]
func (this *MessageController) AddGroup() {
	groupId, _ := this.GetInt64("groupId", 0)
	groupType := this.MustInt("groupType")
	groupName := this.GetString("groupName")
	uId1 := this.MustInt64("uId1")
	uId2, _ := this.GetInt64("uId2", 0)

	addGroupFlag := false
	var group models.Group
	if groupId == 0 {
		addGroupFlag = true
		group.GroupType = groupType
		group.GroupName = groupName
		base.DBEngine.Table("group").InsertOne(&group)
	} else {
		base.DBEngine.Table("group").Where("group_id=?", groupId).Get(&group)
		if group.GroupName != groupName {
			group.GroupName = groupName
			base.DBEngine.Table("group").Where("group_id=?", groupId).Cols("group_name").Update(&group)
		}
	}

	if addGroupFlag {
		hasMember1, _ := base.DBEngine.Table("member").Where("group_id=? and u_id=?", group.GroupId, uId1).Get(new(models.Member))
		if !hasMember1 {
			var member1 models.Member
			member1.GroupId = group.GroupId
			member1.UId = uId1
			base.DBEngine.Table("member").InsertOne(&member1)
		}

		hasMember2, _ := base.DBEngine.Table("member").Where("group_id=? and u_id=?", group.GroupId, uId2).Get(new(models.Member))
		if !hasMember2 {
			var member2 models.Member
			member2.GroupId = group.GroupId
			member2.UId = uId2
			base.DBEngine.Table("member").InsertOne(&member2)
		}
	} else {
		hasMember1, _ := base.DBEngine.Table("member").Where("group_id=? and u_id=?", group.GroupId, uId1).Get(new(models.Member))
		if !hasMember1 {
			var member1 models.Member
			member1.GroupId = group.GroupId
			member1.UId = uId1
			base.DBEngine.Table("member").InsertOne(&member1)
		}
	}

	this.ReturnData = models.GroupInfo{group}
}





