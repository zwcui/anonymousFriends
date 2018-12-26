package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["anonymousFriends/controllers:CommentController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:CommentController"],
		beego.ControllerComments{
			Method: "AddComment",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
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
