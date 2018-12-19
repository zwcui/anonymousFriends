package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
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
			Method: "UpdateUserInfo",
			Router: `/updateUserInfo`,
			AllowHTTPMethods: []string{"patch"},
			Params: nil})

}
