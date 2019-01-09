package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["anonymousFriends/controllers:AdminNoticeController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:AdminNoticeController"],
		beego.ControllerComments{
			Method: "GetAdminNoticeList",
			Router: `/getAdminNoticeList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:AdminNoticeController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:AdminNoticeController"],
		beego.ControllerComments{
			Method: "UpdateAdminNoticeStatus",
			Router: `/updateAdminNoticeStatus`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:CommentController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:CommentController"],
		beego.ControllerComments{
			Method: "AddComment",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"],
		beego.ControllerComments{
			Method: "ThrowDriftBottle",
			Router: `/throwDriftBottle`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"],
		beego.ControllerComments{
			Method: "PickUpDriftBottle",
			Router: `/pickUpDriftBottle`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"],
		beego.ControllerComments{
			Method: "HandleDriftBottle",
			Router: `/handleDriftBottle`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:DriftBottleController"],
		beego.ControllerComments{
			Method: "GetMyDriftBottleList",
			Router: `/getMyDriftBottleList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"],
		beego.ControllerComments{
			Method: "MakeFriends",
			Router: `/makeFriends`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"],
		beego.ControllerComments{
			Method: "HandleMakeFriendsRequest",
			Router: `/handleMakeFriendsRequest`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"],
		beego.ControllerComments{
			Method: "GetFriendList",
			Router: `/getFriendList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"],
		beego.ControllerComments{
			Method: "HandleFriend",
			Router: `/handleFriend`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:FriendController"],
		beego.ControllerComments{
			Method: "GetFriendRequestList",
			Router: `/getFriendRequestList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"],
		beego.ControllerComments{
			Method: "AddGroup",
			Router: `/addGroup`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetUnreadUserChatMessageList",
			Router: `/getUnreadUserChatMessageList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:PublicController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:PublicController"],
		beego.ControllerComments{
			Method: "TestHttpRequest",
			Router: `/testHttpRequest`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:PublicController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:PublicController"],
		beego.ControllerComments{
			Method: "SendSocketMessage",
			Router: `/sendSocketMessage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:PushController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:PushController"],
		beego.ControllerComments{
			Method: "PushCommonMessage",
			Router: `/pushCommonMessage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:PushController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:PushController"],
		beego.ControllerComments{
			Method: "PushSocketMessage",
			Router: `/pushSocketMessage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:ReportController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:ReportController"],
		beego.ControllerComments{
			Method: "PostReport",
			Router: `/postReport`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"],
		beego.ControllerComments{
			Method: "SendSharePositionRequest",
			Router: `/sendSharePositionRequest`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"],
		beego.ControllerComments{
			Method: "GetSharePositionRequest",
			Router: `/getSharePositionRequest`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"],
		beego.ControllerComments{
			Method: "HandleSharePositionRequest",
			Router: `/handleSharePositionRequest`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SharePositionController"],
		beego.ControllerComments{
			Method: "GetSharePositionGroupUserList",
			Router: `/getSharePositionGroupUserList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"],
		beego.ControllerComments{
			Method: "PostSocialDynamic",
			Router: `/postSocialDynamic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"],
		beego.ControllerComments{
			Method: "DeleteSocialDynamic",
			Router: `/deleteSocialDynamic`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"],
		beego.ControllerComments{
			Method: "GetSocialDynamicList",
			Router: `/getSocialDynamicList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:SocialDynamicsController"],
		beego.ControllerComments{
			Method: "LikeSocialDynamic",
			Router: `/likeSocialDynamic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:TagController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetTagList",
			Router: `/getTagList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:TagController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:TagController"],
		beego.ControllerComments{
			Method: "AddTag",
			Router: `/addTag`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"],
		beego.ControllerComments{
			Method: "TestAmapRegeoApi",
			Router: `/testAmapRegeoApi`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"],
		beego.ControllerComments{
			Method: "TestAmapWeatherApi",
			Router: `/testAmapWeatherApi`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:TestController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "SignUp",
			Router: `/signUp`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "SignIn",
			Router: `/signIn`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/getUserInfo`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserAccountInfo",
			Router: `/getUserAccountInfo`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUserInfo",
			Router: `/updateUserInfo`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUserPassword",
			Router: `/updateUserPassword`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUserPosition",
			Router: `/updateUserPosition`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserListByPosition",
			Router: `/getUserListByPosition`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "SignOut",
			Router: `/signOut`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
