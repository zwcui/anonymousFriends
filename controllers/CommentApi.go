/*
@Time : 2018/12/25 上午11:02 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/models"
	"anonymousFriends/base"
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
// @Param	replyId			formData		int64	  		false		"回复id"
// @Param	content			formData		string  		true		"评论内容"
// @Param	senderUid		formData		int64	  		true		"评论发送人"
// @Param   receiverUid     formData        int64	  		true        "评论接收人"
// @Success 200 {string} success
// @router / [post]
func (this *CommentController) AddComment() {
	commentType := this.MustInt("type")
	typeId := this.MustInt64("typeId")
	replyId, _ := this.GetInt64("replyId", 0)
	content := this.MustString("content")
	senderUid := this.MustInt64("senderUid")
	receiverUid := this.MustInt64("receiverUid")

	var comment models.Comment
	comment.Type = commentType
	comment.TypeId = typeId
	comment.ReplyId = replyId
	comment.SenderUid = senderUid
	comment.ReceiverUid = receiverUid
	comment.Content = content
	base.DBEngine.Table("comment").InsertOne(&comment)

	this.ReturnData = "success"
}
