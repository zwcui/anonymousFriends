/*
@Time : 2019/1/2 上午10:12 
@Author : zwcui
@Software: GoLand
*/
package controllers

//共享地理位置模块
type SharePositionController struct {
	apiController
}

func (this *SharePositionController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
}

// @Title 发送好友请求
// @Description 发送好友请求
// @Param	uId						formData		int64  		true		"用户id"
// @Param	groupId					formData		int64  		true		"组id"
// @Success 200 {string} success
// @router /sendSharePositionRequest [post]
func (this *SharePositionController) SendSharePositionRequest() {
	//uId := this.MustInt64("uId")
	//groupId := this.MustInt64("groupId")
	//
	//var sharePositionGroup models.SharePositionGroup
	//hasStoredSharePositionGroup, _ := base.DBEngine.Table("share_position_group").Where("group_id=?", groupId).Get(&sharePositionGroup)
	//if !hasStoredSharePositionGroup {
	//
	//
	//
	//}








}