/*
@Time : 2018/12/25 上午11:02 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
	"anonymousFriends/util"
)

//评论模块
type CommentController struct {
	apiController
}

func (this *CommentController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
}

// @Title 新增评论
// @Description 新增评论
// @Param	type			formData		int		  		true		"类型，1是朋友圈评论"
// @Param	typeId			formData		int64	  		true		"类型id"
// @Param	replyCommentId	formData		int64	  		false		"回复评论id"
// @Param	content			formData		string  		true		"评论内容"
// @Param	senderUid		formData		int64	  		true		"评论发送人"
// @Param   receiverUid     formData        int64	  		true        "评论接收人"
// @Success 200 {string} success
// @router / [post]
func (this *CommentController) AddComment() {
	commentType := this.MustInt("type")
	typeId := this.MustInt64("typeId")
	replyCommentId, _ := this.GetInt64("replyCommentId", 0)
	content := this.MustString("content")
	senderUid := this.MustInt64("senderUid")
	receiverUid := this.MustInt64("receiverUid")

	var socialDynamics models.SocialDynamics
	hasSocialDynamics, _ := base.DBEngine.Table("social_dynamics").Where("id=?", typeId).Get(&socialDynamics)
	if !hasSocialDynamics {
		this.ReturnData = util.GenerateAlertMessage(models.SocialDynamicsError100)
		return
	}

	if commentType == 1 {
		CommentOnSocialDynamics(&socialDynamics, replyCommentId, content, senderUid, receiverUid)
	}

	this.ReturnData = "success"
}




//评论朋友圈
func CommentOnSocialDynamics(socialDynamics *models.SocialDynamics, replyCommentId int64, content string, senderUid int64, receiverUid int64){
	var comment models.Comment
	comment.Type = 1
	comment.TypeId = socialDynamics.Id
	comment.ReplyCommentId = replyCommentId
	comment.SenderUid = senderUid
	comment.ReceiverUid = receiverUid
	comment.Content = content
	base.DBEngine.Table("comment").InsertOne(&comment)

	socialDynamics.CommentNum += 1
	base.DBEngine.Table("social_dynamics").Where("id=?", socialDynamics.Id).Cols("comment_num").Update(&socialDynamics)

	var message models.Message
	message.Content = models.CommentOnSocialDynamics
	message.SenderUid = senderUid
	message.ReceiverUid = receiverUid
	message.Type = 2
	PushSocketMessageToUser(receiverUid, &message, "", 0, "", 3)
}







